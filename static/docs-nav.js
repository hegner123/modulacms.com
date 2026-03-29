(function () {
  "use strict";

  class DocsNav extends HTMLElement {
    connectedCallback() {
      this._highlightActive();
      this._injectMobileToggle();
    }

    _highlightActive() {
      var current = this.dataset.current || window.location.pathname;
      var links = this.querySelectorAll("a.docs-nav-link");
      for (var i = 0; i < links.length; i++) {
        var href = new URL(links[i].href, window.location.origin).pathname;
        if (href === current) {
          links[i].setAttribute("aria-current", "page");
        }
      }
    }

    _injectMobileToggle() {
      var self = this;

      // FAB button (visible < lg)
      var btn = document.createElement("button");
      btn.className =
        "lg:hidden fixed bottom-4 left-4 z-50 flex items-center justify-center size-12 rounded-full bg-accent text-surface-0 shadow-lg";
      btn.setAttribute("aria-label", "Open docs navigation");
      btn.innerHTML =
        '<svg class="size-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path d="M4 6h16M4 12h16M4 18h16"/></svg>';

      // Overlay
      var overlay = document.createElement("div");
      overlay.className = "docs-drawer-overlay";

      // Drawer panel with cloned nav
      var panel = document.createElement("div");
      panel.className = "docs-drawer-panel";
      var navClone = self.querySelector("nav").cloneNode(true);
      panel.appendChild(navClone);

      // Highlight active in clone
      var current = self.dataset.current || window.location.pathname;
      var clonedLinks = panel.querySelectorAll("a.docs-nav-link");
      for (var i = 0; i < clonedLinks.length; i++) {
        var href = new URL(clonedLinks[i].href, window.location.origin).pathname;
        if (href === current) {
          clonedLinks[i].setAttribute("aria-current", "page");
        }
      }

      document.body.appendChild(overlay);
      document.body.appendChild(panel);
      document.body.appendChild(btn);

      function open() {
        overlay.classList.add("open");
        panel.classList.add("open");
      }
      function close() {
        overlay.classList.remove("open");
        panel.classList.remove("open");
      }

      btn.addEventListener("click", open);
      overlay.addEventListener("click", close);
      document.addEventListener("keydown", function (e) {
        if (e.key === "Escape") close();
      });
    }
  }

  customElements.define("docs-nav", DocsNav);
})();
