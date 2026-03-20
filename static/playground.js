// ── Content Model Playground ──
// Single custom element, event delegation, innerHTML rendering.
// No Shadow DOM — site CSS cascades in.

(function () {
  "use strict";

  // ── ULID-like ID generator (26-char Crockford base32) ──

  const CROCKFORD = "0123456789ABCDEFGHJKMNPQRSTVWXYZ";

  function generateId() {
    // Timestamp portion (10 chars) — milliseconds since epoch
    let ts = Date.now();
    const timePart = [];
    for (let i = 0; i < 10; i++) {
      timePart.unshift(CROCKFORD[ts % 32]);
      ts = Math.floor(ts / 32);
    }
    // Random portion (16 chars)
    const randPart = [];
    for (let i = 0; i < 16; i++) {
      randPart.push(CROCKFORD[Math.floor(Math.random() * 32)]);
    }
    return timePart.join("") + randPart.join("");
  }

  function isoNow() {
    return new Date().toISOString();
  }

  function slugify(str) {
    return str
      .toLowerCase()
      .replace(/[^a-z0-9]+/g, "-")
      .replace(/^-|-$/g, "");
  }

  // ── Field types ──

  const FIELD_TYPES = [
    "text",
    "textarea",
    "richtext",
    "number",
    "boolean",
    "date",
    "datetime",
    "media",
    "relation",
    "select",
    "slug",
    "url",
    "email",
    "json",
  ];

  const DATATYPE_TYPES = ["page", "post", "collection", "component", "menu"];

  // ── Placeholder values for API preview ──

  const PLACEHOLDER_VALUES = {
    text: "Lorem ipsum dolor sit amet",
    textarea:
      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
    richtext:
      "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>",
    number: 42,
    boolean: true,
    date: "2026-03-15",
    datetime: "2026-03-15T09:30:00Z",
    media: "med_01JARQ5N8KXVPB3RMDF7HG2W4Y",
    relation: null,
    select: "option_1",
    slug: "example-content",
    url: "https://example.com",
    email: "hello@example.com",
    json: { key: "value" },
    color: "#38bdf8",
    password: "********",
  };

  // ── Quick start templates ──

  const TEMPLATES = {
    blog: {
      name: "blog_post",
      label: "Blog Post",
      type: "post",
      fields: [
        { name: "title", label: "Title", type: "text" },
        { name: "slug", label: "Slug", type: "slug" },
        { name: "body", label: "Body", type: "richtext" },
        { name: "featured_image", label: "Featured Image", type: "media" },
        { name: "published", label: "Published", type: "boolean" },
        { name: "publish_date", label: "Publish Date", type: "datetime" },
      ],
    },
    ecommerce: {
      name: "product",
      label: "Product",
      type: "collection",
      fields: [
        { name: "name", label: "Name", type: "text" },
        { name: "price", label: "Price", type: "number", validation: JSON.stringify({ min: 0 }) },
        { name: "description", label: "Description", type: "richtext" },
        { name: "images", label: "Images", type: "media" },
        { name: "sku", label: "SKU", type: "text" },
        { name: "in_stock", label: "In Stock", type: "boolean" },
      ],
    },
    landing: {
      name: "hero_section",
      label: "Hero Section",
      type: "component",
      fields: [
        { name: "heading", label: "Heading", type: "text" },
        { name: "subheading", label: "Subheading", type: "textarea" },
        { name: "cta_text", label: "CTA Text", type: "text" },
        { name: "cta_url", label: "CTA URL", type: "url" },
        { name: "background_image", label: "Background Image", type: "media" },
      ],
    },
  };

  // ── Escape HTML ──

  function esc(str) {
    if (str === null || str === undefined) return "";
    return String(str)
      .replace(/&/g, "&amp;")
      .replace(/</g, "&lt;")
      .replace(/>/g, "&gt;")
      .replace(/"/g, "&quot;");
  }

  // ── Custom Element ──

  class CmsPlayground extends HTMLElement {
    constructor() {
      super();
      this.state = {
        datatypes: [],
        activeDatatypeId: null,
        editingFieldId: null,
      };
    }

    connectedCallback() {
      this.innerHTML = this.renderShell();
      this.setupDelegation();
      this.renderAll();
    }

    // ── State ──

    setState(patch) {
      Object.assign(this.state, patch);
      this.renderAll();
    }

    getActiveDatatype() {
      if (!this.state.activeDatatypeId) return null;
      return this.state.datatypes.find(
        (d) => d.datatype_id === this.state.activeDatatypeId
      );
    }

    // ── Shell HTML ──

    renderShell() {
      return `
        <div class="pg-container">
          <div class="pg-editor">
            <div class="pg-quick-start" id="pg-quick-start">
              <span class="pg-quick-start-label">Quick Start</span>
              <button class="pg-quick-btn" data-action="template" data-template="blog">Blog</button>
              <button class="pg-quick-btn" data-action="template" data-template="ecommerce">E-commerce</button>
              <button class="pg-quick-btn" data-action="template" data-template="landing">Landing Page</button>
            </div>

            <div class="pg-panel">
              <h2 class="pg-panel-title">Create Datatype</h2>
              <div class="pg-form-row">
                <div class="pg-form-group">
                  <label for="dt-name">Name</label>
                  <input class="pg-input" id="dt-name" placeholder="blog_post" data-auto-slug="dt-label"/>
                </div>
                <div class="pg-form-group">
                  <label for="dt-label">Label</label>
                  <input class="pg-input" id="dt-label" placeholder="Blog Post"/>
                </div>
                <div class="pg-form-group">
                  <label for="dt-type">Type</label>
                  <select class="pg-select" id="dt-type">
                    ${DATATYPE_TYPES.map((t) => `<option value="${t}">${t}</option>`).join("")}
                  </select>
                </div>
                <button class="pg-btn pg-btn-primary" data-action="add-datatype">Add</button>
              </div>
            </div>

            <div id="pg-datatype-list"></div>
            <div id="pg-field-panel"></div>
          </div>

          <div class="pg-preview">
            <div class="pg-panel">
              <div class="pg-tabs">
                <button class="pg-tab active" data-action="tab" data-tab="schema">Schema Table</button>
                <button class="pg-tab" data-action="tab" data-tab="api">API Preview</button>
              </div>
              <div id="pg-tab-schema" class="pg-tab-content active"></div>
              <div id="pg-tab-api" class="pg-tab-content"></div>
            </div>
          </div>
        </div>`;
    }

    // ── Event Delegation ──

    setupDelegation() {
      this.addEventListener("click", (e) => {
        const btn = e.target.closest("[data-action]");
        if (!btn) return;
        const action = btn.dataset.action;

        switch (action) {
          case "template":
            this.applyTemplate(btn.dataset.template);
            break;
          case "add-datatype":
            this.addDatatype();
            break;
          case "select-datatype":
            this.selectDatatype(btn.dataset.id);
            break;
          case "delete-datatype":
            e.stopPropagation();
            this.deleteDatatype(btn.dataset.id);
            break;
          case "add-field":
            this.addField();
            break;
          case "move-field":
            this.moveField(btn.dataset.id, parseInt(btn.dataset.dir, 10));
            break;
          case "delete-field":
            this.deleteField(btn.dataset.id);
            break;
          case "tab":
            this.switchTab(btn.dataset.tab);
            break;
          case "toggle-bool":
            this.toggleFieldDefault(btn.dataset.id);
            break;
        }
      });

      // Auto-slug: typing in label fills name
      this.addEventListener("input", (e) => {
        if (e.target.id === "dt-label") {
          const nameInput = this.querySelector("#dt-name");
          if (nameInput) {
            nameInput.value = slugify(e.target.value).replace(/-/g, "_");
          }
        }
        if (e.target.id === "field-label") {
          const nameInput = this.querySelector("#field-name");
          if (nameInput) {
            nameInput.value = slugify(e.target.value).replace(/-/g, "_");
          }
        }
        // Re-render type config when field type changes
        if (e.target.id === "field-type") {
          this.renderTypeConfig();
        }
      });

      this.addEventListener("change", (e) => {
        if (e.target.id === "field-type") {
          this.renderTypeConfig();
        }
      });
    }

    // ── Render orchestrator ──

    renderAll() {
      this.renderDatatypeList();
      this.renderFieldPanel();
      this.renderSchemaPreview();
      this.renderApiPreview();
    }

    // ── Datatype CRUD ──

    addDatatype() {
      const nameEl = this.querySelector("#dt-name");
      const labelEl = this.querySelector("#dt-label");
      const typeEl = this.querySelector("#dt-type");

      const name = nameEl.value.trim();
      const label = labelEl.value.trim();
      const type = typeEl.value;

      if (!name || !label) return;

      const now = isoNow();
      const dt = {
        datatype_id: generateId(),
        parent_id: null,
        sort_order: this.state.datatypes.length,
        name: name,
        label: label,
        type: type,
        date_created: now,
        date_modified: now,
        fields: [],
      };

      nameEl.value = "";
      labelEl.value = "";

      this.state.datatypes.push(dt);
      this.state.activeDatatypeId = dt.datatype_id;
      this.state.editingFieldId = null;
      this.renderAll();
    }

    selectDatatype(id) {
      this.setState({
        activeDatatypeId: id,
        editingFieldId: null,
      });
    }

    deleteDatatype(id) {
      this.state.datatypes = this.state.datatypes.filter(
        (d) => d.datatype_id !== id
      );
      if (this.state.activeDatatypeId === id) {
        this.state.activeDatatypeId =
          this.state.datatypes.length > 0
            ? this.state.datatypes[0].datatype_id
            : null;
      }
      this.state.editingFieldId = null;
      this.renderAll();
    }

    renderDatatypeList() {
      const el = this.querySelector("#pg-datatype-list");
      if (!el) return;

      if (this.state.datatypes.length === 0) {
        el.innerHTML = "";
        return;
      }

      el.innerHTML = `
        <div class="pg-panel">
          <h2 class="pg-panel-title">Datatypes</h2>
          <div class="pg-datatype-list">
            ${this.state.datatypes
              .map(
                (dt) => `
              <div class="pg-datatype-card${dt.datatype_id === this.state.activeDatatypeId ? " active" : ""}"
                   data-action="select-datatype" data-id="${dt.datatype_id}">
                <div>
                  <div class="pg-datatype-name">${esc(dt.label)}</div>
                  <div class="pg-datatype-label">${esc(dt.name)}</div>
                </div>
                <div class="pg-datatype-meta">
                  <span class="pg-datatype-type">${esc(dt.type)}</span>
                  <span style="color: var(--text-muted); font-size: var(--text-xs);">${dt.fields.length} field${dt.fields.length !== 1 ? "s" : ""}</span>
                  <button class="pg-btn pg-btn-danger pg-btn-sm" data-action="delete-datatype" data-id="${dt.datatype_id}">&times;</button>
                </div>
              </div>`
              )
              .join("")}
          </div>
        </div>`;
    }

    // ── Field CRUD ──

    addField() {
      const dt = this.getActiveDatatype();
      if (!dt) return;

      const nameEl = this.querySelector("#field-name");
      const labelEl = this.querySelector("#field-label");
      const typeEl = this.querySelector("#field-type");

      const name = nameEl.value.trim();
      const label = labelEl.value.trim();
      const type = typeEl.value;

      if (!name || !label) return;

      const now = isoNow();
      const field = {
        field_id: generateId(),
        parent_id: dt.datatype_id,
        sort_order: dt.fields.length,
        name: name,
        label: label,
        type: type,
        data: null,
        validation: null,
        ui_config: null,
        translatable: false,
        roles: null,
        date_created: now,
        date_modified: now,
      };

      // Gather type-specific config
      this.applyTypeConfig(field);

      dt.fields.push(field);
      dt.date_modified = isoNow();

      nameEl.value = "";
      labelEl.value = "";

      this.renderAll();
    }

    applyTypeConfig(field) {
      switch (field.type) {
        case "select": {
          const optEl = this.querySelector("#field-options");
          if (optEl) {
            const options = optEl.value
              .split("\n")
              .map((s) => s.trim())
              .filter(Boolean);
            if (options.length > 0) {
              field.data = JSON.stringify(options);
            }
          }
          break;
        }
        case "relation": {
          const relEl = this.querySelector("#field-relation-target");
          if (relEl && relEl.value) {
            field.data = JSON.stringify({ target: relEl.value });
          }
          break;
        }
        case "number": {
          const minEl = this.querySelector("#field-min");
          const maxEl = this.querySelector("#field-max");
          const v = {};
          if (minEl && minEl.value !== "") v.min = Number(minEl.value);
          if (maxEl && maxEl.value !== "") v.max = Number(maxEl.value);
          if (Object.keys(v).length > 0) {
            field.validation = JSON.stringify(v);
          }
          break;
        }
        case "text":
        case "textarea":
        case "richtext": {
          const maxLenEl = this.querySelector("#field-maxlength");
          if (maxLenEl && maxLenEl.value !== "") {
            field.validation = JSON.stringify({
              maxLength: Number(maxLenEl.value),
            });
          }
          const phEl = this.querySelector("#field-placeholder");
          if (phEl && phEl.value !== "") {
            field.ui_config = JSON.stringify({ placeholder: phEl.value });
          }
          break;
        }
        case "boolean": {
          const togEl = this.querySelector("#field-default-bool");
          if (togEl) {
            field.data = JSON.stringify({
              default: togEl.classList.contains("on"),
            });
          }
          break;
        }
      }
    }

    moveField(fieldId, dir) {
      const dt = this.getActiveDatatype();
      if (!dt) return;

      const idx = dt.fields.findIndex((f) => f.field_id === fieldId);
      if (idx < 0) return;
      const newIdx = idx + dir;
      if (newIdx < 0 || newIdx >= dt.fields.length) return;

      const tmp = dt.fields[idx];
      dt.fields[idx] = dt.fields[newIdx];
      dt.fields[newIdx] = tmp;

      // Update sort_order
      dt.fields.forEach((f, i) => {
        f.sort_order = i;
      });

      dt.date_modified = isoNow();
      this.renderAll();
    }

    deleteField(fieldId) {
      const dt = this.getActiveDatatype();
      if (!dt) return;

      dt.fields = dt.fields.filter((f) => f.field_id !== fieldId);
      dt.fields.forEach((f, i) => {
        f.sort_order = i;
      });
      dt.date_modified = isoNow();
      this.renderAll();
    }

    toggleFieldDefault(fieldId) {
      const togEl = this.querySelector("#field-default-bool");
      if (togEl) {
        togEl.classList.toggle("on");
      }
    }

    renderFieldPanel() {
      const el = this.querySelector("#pg-field-panel");
      if (!el) return;

      const dt = this.getActiveDatatype();
      if (!dt) {
        el.innerHTML = "";
        return;
      }

      const fieldListHtml =
        dt.fields.length === 0
          ? `<div class="pg-empty">No fields yet. Add one below.</div>`
          : `<div class="pg-field-list">
              ${dt.fields
                .map(
                  (f, i) => `
                <div class="pg-field-row" data-type="${esc(f.type)}">
                  <span class="pg-type-badge" data-type="${esc(f.type)}">${esc(f.type)}</span>
                  <div class="pg-field-info">
                    <div class="pg-field-name">${esc(f.name)}</div>
                    <div class="pg-field-label">${esc(f.label)}</div>
                  </div>
                  <div class="pg-field-actions">
                    <button class="pg-btn pg-btn-ghost pg-btn-sm" data-action="move-field" data-id="${f.field_id}" data-dir="-1" ${i === 0 ? "disabled" : ""}>&#9650;</button>
                    <button class="pg-btn pg-btn-ghost pg-btn-sm" data-action="move-field" data-id="${f.field_id}" data-dir="1" ${i === dt.fields.length - 1 ? "disabled" : ""}>&#9660;</button>
                    <button class="pg-btn pg-btn-danger pg-btn-sm" data-action="delete-field" data-id="${f.field_id}">&times;</button>
                  </div>
                </div>`
                )
                .join("")}
            </div>`;

      el.innerHTML = `
        <div class="pg-panel">
          <h2 class="pg-panel-title">Fields — ${esc(dt.label)}</h2>
          ${fieldListHtml}
          <div class="pg-divider"></div>
          <h3 style="font-size: var(--text-sm); color: var(--text-secondary); margin: var(--space-3) 0 var(--space-2) 0; font-weight: 500;">Add Field</h3>
          <div class="pg-form-row">
            <div class="pg-form-group">
              <label for="field-name">Name</label>
              <input class="pg-input" id="field-name" placeholder="field_name"/>
            </div>
            <div class="pg-form-group">
              <label for="field-label">Label</label>
              <input class="pg-input" id="field-label" placeholder="Field Label"/>
            </div>
            <div class="pg-form-group">
              <label for="field-type">Type</label>
              <select class="pg-select" id="field-type">
                ${FIELD_TYPES.map((t) => `<option value="${t}">${t}</option>`).join("")}
              </select>
            </div>
            <button class="pg-btn pg-btn-primary" data-action="add-field">Add</button>
          </div>
          <div id="pg-type-config"></div>
        </div>`;

      this.renderTypeConfig();
    }

    renderTypeConfig() {
      const el = this.querySelector("#pg-type-config");
      if (!el) return;

      const typeEl = this.querySelector("#field-type");
      if (!typeEl) return;

      const type = typeEl.value;
      let html = "";

      switch (type) {
        case "select":
          html = `
            <div class="pg-type-config">
              <div class="pg-type-config-label">Options (one per line)</div>
              <textarea class="pg-textarea" id="field-options" placeholder="option_1&#10;option_2&#10;option_3"></textarea>
            </div>`;
          break;
        case "relation":
          html = `
            <div class="pg-type-config">
              <div class="pg-type-config-label">Target Datatype</div>
              <select class="pg-select" id="field-relation-target">
                <option value="">Select target...</option>
                ${this.state.datatypes
                  .map(
                    (d) =>
                      `<option value="${d.datatype_id}">${esc(d.label)} (${esc(d.name)})</option>`
                  )
                  .join("")}
              </select>
            </div>`;
          break;
        case "number":
          html = `
            <div class="pg-type-config">
              <div class="pg-type-config-label">Validation</div>
              <div class="pg-form-row">
                <div class="pg-form-group">
                  <label for="field-min">Min</label>
                  <input class="pg-input" id="field-min" type="number" placeholder="0"/>
                </div>
                <div class="pg-form-group">
                  <label for="field-max">Max</label>
                  <input class="pg-input" id="field-max" type="number" placeholder="100"/>
                </div>
              </div>
            </div>`;
          break;
        case "text":
        case "textarea":
          html = `
            <div class="pg-type-config">
              <div class="pg-type-config-label">Options</div>
              <div class="pg-form-row">
                <div class="pg-form-group">
                  <label for="field-maxlength">Max Length</label>
                  <input class="pg-input" id="field-maxlength" type="number" placeholder="255"/>
                </div>
                <div class="pg-form-group">
                  <label for="field-placeholder">Placeholder</label>
                  <input class="pg-input" id="field-placeholder" placeholder="Enter text..."/>
                </div>
              </div>
            </div>`;
          break;
        case "richtext":
          html = `
            <div class="pg-type-config">
              <div class="pg-type-config-label">Options</div>
              <div class="pg-form-row">
                <div class="pg-form-group">
                  <label for="field-maxlength">Max Length</label>
                  <input class="pg-input" id="field-maxlength" type="number" placeholder="10000"/>
                </div>
              </div>
            </div>`;
          break;
        case "boolean":
          html = `
            <div class="pg-type-config">
              <div class="pg-type-config-label">Default Value</div>
              <div style="display: flex; align-items: center; gap: var(--space-2);">
                <div class="pg-toggle" id="field-default-bool" data-action="toggle-bool" data-id="new"></div>
                <span style="font-size: var(--text-sm); color: var(--text-secondary);">false</span>
              </div>
            </div>`;
          break;
        case "media":
          html = `
            <div class="pg-type-config">
              <div class="pg-info-note">Stores a MediaID reference. The API resolves this to a full media object with URL, dimensions, and metadata.</div>
            </div>`;
          break;
      }

      el.innerHTML = html;
    }

    // ── Schema Preview ──

    renderSchemaPreview() {
      const el = this.querySelector("#pg-tab-schema");
      if (!el) return;

      if (this.state.datatypes.length === 0) {
        el.innerHTML = `<div class="pg-empty">Create a datatype to see the schema.</div>`;
        return;
      }

      // Datatypes table
      const dtRows = this.state.datatypes
        .map(
          (dt) => `
        <tr>
          <td>${esc(dt.datatype_id)}</td>
          <td class="${dt.parent_id ? "" : "null"}">${dt.parent_id ? esc(dt.parent_id) : "NULL"}</td>
          <td>${dt.sort_order}</td>
          <td>${esc(dt.name)}</td>
          <td>${esc(dt.label)}</td>
          <td>${esc(dt.type)}</td>
          <td>${esc(dt.date_created)}</td>
          <td>${esc(dt.date_modified)}</td>
        </tr>`
        )
        .join("");

      // Fields table
      const allFields = this.state.datatypes.flatMap((dt) => dt.fields);
      const fieldRows =
        allFields.length === 0
          ? `<tr><td colspan="12" class="null" style="text-align:center;">No fields defined</td></tr>`
          : allFields
              .map(
                (f) => `
        <tr>
          <td>${esc(f.field_id)}</td>
          <td>${esc(f.parent_id)}</td>
          <td>${f.sort_order}</td>
          <td>${esc(f.name)}</td>
          <td>${esc(f.label)}</td>
          <td>${esc(f.type)}</td>
          <td class="${f.data ? "" : "null"}" title="${f.data ? esc(f.data) : ""}">${f.data ? esc(truncateJson(f.data)) : "NULL"}</td>
          <td class="${f.validation ? "" : "null"}" title="${f.validation ? esc(f.validation) : ""}">${f.validation ? esc(truncateJson(f.validation)) : "NULL"}</td>
          <td class="${f.ui_config ? "" : "null"}" title="${f.ui_config ? esc(f.ui_config) : ""}">${f.ui_config ? esc(truncateJson(f.ui_config)) : "NULL"}</td>
          <td>${f.translatable}</td>
          <td class="${f.roles ? "" : "null"}">${f.roles ? esc(f.roles) : "NULL"}</td>
          <td>${esc(f.date_created)}</td>
        </tr>`
              )
              .join("");

      el.innerHTML = `
        <div class="pg-table-label">datatypes</div>
        <div class="pg-schema-table-wrap">
          <table class="pg-schema-table">
            <thead>
              <tr>
                <th>datatype_id</th>
                <th>parent_id</th>
                <th>sort_order</th>
                <th>name</th>
                <th>label</th>
                <th>type</th>
                <th>date_created</th>
                <th>date_modified</th>
              </tr>
            </thead>
            <tbody>${dtRows}</tbody>
          </table>
        </div>

        <div style="height: var(--space-6);"></div>

        <div class="pg-table-label">fields</div>
        <div class="pg-schema-table-wrap">
          <table class="pg-schema-table">
            <thead>
              <tr>
                <th>field_id</th>
                <th>parent_id</th>
                <th>sort_order</th>
                <th>name</th>
                <th>label</th>
                <th>type</th>
                <th>data</th>
                <th>validation</th>
                <th>ui_config</th>
                <th>translatable</th>
                <th>roles</th>
                <th>date_created</th>
              </tr>
            </thead>
            <tbody>${fieldRows}</tbody>
          </table>
        </div>`;
    }

    // ── API Preview ──

    renderApiPreview() {
      const el = this.querySelector("#pg-tab-api");
      if (!el) return;

      if (this.state.datatypes.length === 0) {
        el.innerHTML = `<div class="pg-empty">Create a datatype to see the API response.</div>`;
        return;
      }

      const response = this.buildApiResponse();
      const jsonStr = JSON.stringify(response, null, 2);
      const highlighted = syntaxHighlight(jsonStr);

      el.innerHTML = `
        <div class="pg-json-preview">
          <pre>${highlighted}</pre>
        </div>`;
    }

    buildApiResponse() {
      const result = {};
      for (const dt of this.state.datatypes) {
        const entry = {
          _meta: {
            datatype_id: dt.datatype_id,
            type: dt.type,
            date_created: dt.date_created,
            date_modified: dt.date_modified,
          },
        };
        for (const f of dt.fields) {
          entry[f.name] = PLACEHOLDER_VALUES[f.type] !== undefined
            ? PLACEHOLDER_VALUES[f.type]
            : null;
        }
        result[dt.name] = entry;
      }

      return {
        format: "clean",
        data: result,
      };
    }

    // ── Tab switching ──

    switchTab(tab) {
      const tabs = this.querySelectorAll(".pg-tab");
      const contents = this.querySelectorAll(".pg-tab-content");

      tabs.forEach((t) => {
        t.classList.toggle("active", t.dataset.tab === tab);
      });
      contents.forEach((c) => {
        c.classList.toggle(
          "active",
          c.id === "pg-tab-" + tab
        );
      });
    }

    // ── Templates ──

    applyTemplate(templateKey) {
      const tmpl = TEMPLATES[templateKey];
      if (!tmpl) return;

      const now = isoNow();
      const dtId = generateId();

      const fields = tmpl.fields.map((f, i) => ({
        field_id: generateId(),
        parent_id: dtId,
        sort_order: i,
        name: f.name,
        label: f.label,
        type: f.type,
        data: f.data || null,
        validation: f.validation || null,
        ui_config: f.ui_config || null,
        translatable: false,
        roles: null,
        date_created: now,
        date_modified: now,
      }));

      const dt = {
        datatype_id: dtId,
        parent_id: null,
        sort_order: this.state.datatypes.length,
        name: tmpl.name,
        label: tmpl.label,
        type: tmpl.type,
        date_created: now,
        date_modified: now,
        fields: fields,
      };

      this.state.datatypes.push(dt);
      this.state.activeDatatypeId = dt.datatype_id;
      this.state.editingFieldId = null;
      this.renderAll();
    }
  }

  // ── Helpers ──

  function truncateJson(str) {
    if (!str) return "";
    if (str.length <= 30) return str;
    return str.substring(0, 27) + "...";
  }

  function syntaxHighlight(json) {
    return json.replace(
      /("(\\u[a-fA-F0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false)\b|-?\d+(?:\.\d*)?(?:[eE][+-]?\d+)?|\bnull\b)/g,
      function (match) {
        let cls = "json-number";
        if (/^"/.test(match)) {
          if (/:$/.test(match)) {
            cls = "json-key";
            // Remove trailing colon for wrapping, add it back after
            return '<span class="' + cls + '">' + esc(match.slice(0, -1)) + "</span>:";
          }
          cls = "json-string";
        } else if (/true|false/.test(match)) {
          cls = "json-bool";
        } else if (/null/.test(match)) {
          cls = "json-null";
        }
        return '<span class="' + cls + '">' + esc(match) + "</span>";
      }
    );
  }

  // ── Register ──

  customElements.define("cms-playground", CmsPlayground);
})();
