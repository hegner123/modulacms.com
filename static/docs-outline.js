(function () {
  "use strict";

  class DocsOutline extends HTMLElement {
    connectedCallback() {
      this._buildTOC();
      this._setupScrollSpy();
    }

    disconnectedCallback() {
      if (this._observer) this._observer.disconnect();
    }

    _buildTOC() {
      var article = document.getElementById("docs-content");
      if (!article) return;

      var headings = article.querySelectorAll("h2[id], h3[id]");
      var ol = this.querySelector("#docs-toc");
      if (!ol || headings.length === 0) return;

      this._headings = [];
      var fragment = document.createDocumentFragment();

      for (var i = 0; i < headings.length; i++) {
        var h = headings[i];
        var level = h.tagName === "H2" ? 2 : 3;
        var li = document.createElement("li");
        var a = document.createElement("a");
        a.href = "#" + h.id;
        a.textContent = h.textContent;
        a.className = "docs-toc-link";
        a.dataset.level = String(level);
        li.appendChild(a);
        fragment.appendChild(li);
        this._headings.push({ id: h.id, el: h, link: a, visible: false });
      }

      ol.appendChild(fragment);
    }

    _setupScrollSpy() {
      if (!this._headings || this._headings.length === 0) return;

      var self = this;
      this._observer = new IntersectionObserver(
        function (entries) {
          for (var i = 0; i < entries.length; i++) {
            var entry = entries[i];
            for (var j = 0; j < self._headings.length; j++) {
              if (self._headings[j].el === entry.target) {
                self._headings[j].visible = entry.isIntersecting;
                break;
              }
            }
          }
          self._updateActive();
        },
        { rootMargin: "-80px 0px -70% 0px" }
      );

      for (var i = 0; i < this._headings.length; i++) {
        this._observer.observe(this._headings[i].el);
      }
    }

    _updateActive() {
      var active = null;

      // First visible heading wins
      for (var i = 0; i < this._headings.length; i++) {
        if (this._headings[i].visible) {
          active = this._headings[i];
          break;
        }
      }

      // Fallback: last heading above viewport top
      if (!active) {
        var scrollY = window.scrollY + 100;
        for (var i = 0; i < this._headings.length; i++) {
          if (this._headings[i].el.offsetTop <= scrollY) {
            active = this._headings[i];
          }
        }
      }

      for (var i = 0; i < this._headings.length; i++) {
        var h = this._headings[i];
        if (h === active) {
          h.link.classList.add("active");
        } else {
          h.link.classList.remove("active");
        }
      }
    }
  }

  customElements.define("docs-outline", DocsOutline);
})();
