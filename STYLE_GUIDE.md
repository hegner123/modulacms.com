# Style Guide

Visual design tokens and conventions for the modulacms.com marketing site.

**Approach:** Vanilla CSS. All tokens defined as CSS custom properties on `:root`. Dark theme is the default and only theme at launch.

**Aesthetic:** Minimal, technical, whitespace-heavy, code-forward. The site should feel like a well-designed man page, not a SaaS landing page.

---

## Color Palette

### Backgrounds (surface layers)

Darkest to lightest. Each layer sits visually above the previous one.

| Token | Value | Use |
|-------|-------|-----|
| `--surface-0` | `#0a0a0f` | Page background |
| `--surface-1` | `#12121a` | Section backgrounds, cards |
| `--surface-2` | `#1a1a25` | Elevated elements, code blocks |
| `--surface-3` | `#242430` | Hover states, active inputs |

### Foreground (text hierarchy)

| Token | Value | Contrast on surface-0 | Use |
|-------|-------|-----------------------|-----|
| `--text-primary` | `#e8e8ed` | 16.2:1 | Body text, headings |
| `--text-secondary` | `#a0a0b0` | 7.7:1 | Descriptions, secondary labels |
| `--text-muted` | `#8585a0` | 5.5:1 | Captions, metadata, placeholders |

All three meet WCAG AA (4.5:1 minimum) against `--surface-0`.

### Accent (blue/cyan)

| Token | Value | Use |
|-------|-------|-----|
| `--accent` | `#38bdf8` | Links, primary buttons, highlights |
| `--accent-hover` | `#7dd3fc` | Hover state for accent elements |
| `--accent-active` | `#0ea5e9` | Active/pressed state |
| `--accent-muted` | `rgba(56, 189, 248, 0.12)` | Accent backgrounds, subtle highlights |

`--accent` on `--surface-0`: 9.2:1 contrast ratio (passes AAA).

### Semantic

| Token | Value | Use |
|-------|-------|-----|
| `--color-error` | `#f87171` | Error messages, destructive actions |
| `--color-success` | `#4ade80` | Success messages, confirmations |
| `--color-warning` | `#fbbf24` | Warnings, caution notices |

### Borders

| Token | Value | Use |
|-------|-------|-----|
| `--border-default` | `#1e1e2a` | Card borders, dividers |
| `--border-subtle` | `#16161f` | Faint separators |
| `--border-accent` | `rgba(56, 189, 248, 0.3)` | Accent borders (code blocks, focus rings) |

---

## Typography

### Font stacks

```css
--font-sans: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont,
  "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;

--font-mono: ui-monospace, SFMono-Regular, "SF Mono", Menlo, Consolas,
  "Liberation Mono", monospace;
```

No external font loads. System fonts only.

### Type scale

Base size: `16px` (1rem). Approximate scale factor: `1.25` (major third), rounded to clean rem values.

| Token | Size | Line height | Use |
|-------|------|-------------|-----|
| `--text-xs` | `0.75rem` (12px) | 1.5 | Fine print, badges |
| `--text-sm` | `0.875rem` (14px) | 1.5 | Captions, metadata |
| `--text-base` | `1rem` (16px) | 1.625 | Body text |
| `--text-lg` | `1.25rem` (20px) | 1.5 | Lead paragraphs, large body |
| `--text-xl` | `1.5rem` (24px) | 1.35 | H4, subheadings |
| `--text-2xl` | `2rem` (32px) | 1.25 | H3 |
| `--text-3xl` | `2.5rem` (40px) | 1.2 | H2 |
| `--text-4xl` | `3rem` (48px) | 1.1 | H1, hero headings |

### Heading weights

| Element | Weight |
|---------|--------|
| H1 | `700` (bold) |
| H2 | `600` (semibold) |
| H3 | `600` (semibold) |
| H4-H6 | `500` (medium) |

Body text uses `400` (regular). Code uses `400`.

---

## Spacing

Base unit: `4px`. All spacing derives from this base.

| Token | Value | Use |
|-------|-------|-----|
| `--space-1` | `0.25rem` (4px) | Tight inline gaps |
| `--space-2` | `0.5rem` (8px) | Icon gaps, compact padding |
| `--space-3` | `0.75rem` (12px) | Button padding (vertical) |
| `--space-4` | `1rem` (16px) | Default element gaps |
| `--space-6` | `1.5rem` (24px) | Card padding, group spacing |
| `--space-8` | `2rem` (32px) | Section internal padding |
| `--space-12` | `3rem` (48px) | Large section gaps |
| `--space-16` | `4rem` (64px) | Section vertical padding |
| `--space-24` | `6rem` (96px) | Hero/major section padding |
| `--space-32` | `8rem` (128px) | Page-level vertical rhythm |

### Named spacing aliases

For readability in component styles:

| Token | Maps to | Use |
|-------|---------|-----|
| `--section-padding-y` | `var(--space-24)` | Top/bottom padding on page sections |
| `--section-padding-x` | `var(--space-6)` | Left/right padding on page sections |
| `--card-padding` | `var(--space-6)` | Internal card padding |
| `--card-gap` | `var(--space-4)` | Gap between card children |
| `--inline-gap` | `var(--space-2)` | Horizontal gap between inline elements |
| `--stack-gap` | `var(--space-4)` | Vertical gap between stacked elements |

---

## Layout

### Content width

| Token | Value | Use |
|-------|-------|-----|
| `--max-width-content` | `72rem` (1152px) | Main content container |
| `--max-width-narrow` | `48rem` (768px) | Docs, text-heavy pages |
| `--max-width-wide` | `90rem` (1440px) | Full-width sections with padding |

### Breakpoints

Mobile-first. Breakpoints used in `@media (min-width: ...)` queries.

| Name | Value | Target |
|------|-------|--------|
| `sm` | `640px` | Large phones, landscape |
| `md` | `768px` | Tablets |
| `lg` | `1024px` | Small desktops |
| `xl` | `1280px` | Large desktops |

CSS custom properties cannot be used in media queries. Use these literal values in `@media` rules and add a comment referencing the breakpoint name:

```css
@media (min-width: 768px) { /* md */ }
```

### Grid

No grid framework. Use CSS Grid and Flexbox directly.

- **Page-level layout:** Single column, centered with `max-width` + `margin-inline: auto`.
- **Feature grids:** `display: grid; grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));` with `gap: var(--space-6)`.
- **Two-column splits:** `grid-template-columns: 1fr 1fr;` at `lg`, stacked below.

---

## Components

### Code blocks

```css
.code-block {
  background: var(--surface-2);
  border: 1px solid var(--border-accent);
  border-radius: 6px;
  padding: var(--space-6);
  font-family: var(--font-mono);
  font-size: var(--text-sm);
  line-height: 1.7;
  overflow-x: auto;
}
```

- Accent-colored left border variant: `border-left: 3px solid var(--accent);` with other borders using `--border-default`.
- No background gradients or shadows on code blocks.

### Inline code

```css
code {
  background: var(--surface-2);
  padding: 0.15em 0.4em;
  border-radius: 3px;
  font-family: var(--font-mono);
  font-size: 0.875em;
}
```

### Buttons

Primary style is ghost/outline. No filled buttons by default.

```css
.btn {
  display: inline-flex;
  align-items: center;
  gap: var(--space-2);
  padding: var(--space-3) var(--space-6);
  border: 1px solid var(--border-accent);
  border-radius: 6px;
  background: transparent;
  color: var(--accent);
  font-family: var(--font-mono);
  font-size: var(--text-sm);
  cursor: pointer;
  transition: background var(--transition-fast), border-color var(--transition-fast);
}

.btn:hover {
  background: var(--accent-muted);
  border-color: var(--accent);
}

.btn:active {
  background: rgba(56, 189, 248, 0.2);
}
```

Monospace font on buttons reinforces the technical aesthetic.

### Cards

```css
.card {
  background: var(--surface-1);
  border: 1px solid var(--border-default);
  border-radius: 8px;
  padding: var(--card-padding);
}
```

- No box shadows. Borders provide the only elevation cue.
- Hover state (if interactive): shift border to `--border-accent`.

### Links

```css
a {
  color: var(--accent);
  text-decoration: underline;
  text-underline-offset: 3px;
  text-decoration-color: rgba(56, 189, 248, 0.4);
  transition: text-decoration-color var(--transition-fast);
}

a:hover {
  text-decoration-color: var(--accent);
}
```

---

## Motion

| Token | Value | Use |
|-------|-------|-----|
| `--transition-fast` | `120ms ease-out` | Hover states, small interactions |
| `--transition-normal` | `200ms ease-out` | Expanding elements, tab switches |
| `--transition-slow` | `350ms ease-out` | Page-level transitions, modals |

### Reduced motion

```css
@media (prefers-reduced-motion: reduce) {
  *, *::before, *::after {
    transition-duration: 0.01ms !important;
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
  }
}
```

---

## Conventions

### Custom property naming

Pattern: `--{category}-{variant}`

| Category | Examples |
|----------|----------|
| `surface-*` | `--surface-0`, `--surface-1`, `--surface-2`, `--surface-3` |
| `text-*` | `--text-primary`, `--text-secondary`, `--text-muted` |
| `accent*` | `--accent`, `--accent-hover`, `--accent-active`, `--accent-muted` |
| `color-*` | `--color-error`, `--color-success`, `--color-warning` |
| `border-*` | `--border-default`, `--border-subtle`, `--border-accent` |
| `space-*` | `--space-1` through `--space-32` |
| `text-*` (size) | `--text-xs` through `--text-4xl` |
| `font-*` | `--font-sans`, `--font-mono` |
| `transition-*` | `--transition-fast`, `--transition-normal`, `--transition-slow` |
| `max-width-*` | `--max-width-content`, `--max-width-narrow`, `--max-width-wide` |



Scoped component styles should reference tokens via `var(--token-name)` and avoid hardcoded color/spacing values.
