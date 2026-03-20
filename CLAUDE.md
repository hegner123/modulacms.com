# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Is

Go web server for the modulacms.com marketing site. Fetches content from the ModulaCMS REST API (`?format=clean`) and renders HTML server-side using templ components. Uses Tailwind CSS 4.2 for styling.

## Build & Run

```bash
# Build CSS (styles/app.css → static/app.css)
tailwindcss -i styles/app.css -o static/app.css

# Generate templ code (required before build — *.templ → *_templ.go)
templ generate

# Build binary
go build -o modulacms.com .

# Run (requires CMS_BASE_URL env var)
source .env && go run .

# Or use just (runs css → generate → build):
just build
```

Port defaults to `5050` (override with `PORT` env var).

Required env: `CMS_BASE_URL`. Optional: `CMS_API_KEY`, `PORT`.

### Docker

```bash
docker build -t modulacms-site .
docker run -p 5050:5050 -e CMS_BASE_URL=... -e CMS_API_KEY=... modulacms-site
```

## Architecture

**Request flow:** HTTP request → `server.go` route → `render.go` handler → `fetch.go` (CMS API call + media resolution) → `components/*.templ` (HTML rendering) → response.

### Key Files

- `main.go` — Entry point. Creates ModulaCMS SDK client, HTTP server, graceful shutdown.
- `server.go` — Mux with route registrations. Currently only `GET /{$}` (home page).
- `render.go` — Handler functions. Fetches content, renders templ components.
- `fetch.go` — CMS API calls via SDK. `resolveMediaURLs` walks the tree to resolve Image media IDs to URLs.
- `content/content.go` — Full content model. This is the most important file to understand.

### Content Model (`content/content.go`)

The CMS API returns a nested JSON tree. The `Child` struct is a **discriminated union** — exactly one pointer field is non-nil per instance. The `"type"` JSON field determines which Go struct to unmarshal into.

Type names from the API map to Go types via `unmarshalChild()`:
- `"Row"` → `Row`, `"Column"/"Columns"` → `Columns`, `"Grid"` → `Grid`, `"Area"` → `Area`
- `"CTA"` → `CTA`, `"Card"` → `Card`, `"Rich Text"` → `RichText`, `"Text"` → `Text`, `"Image"` → `Image`, `"Button"` → `Button`
- `"Content"` → `PostContent`, `"Section"` → `Section`, `"Code Block"` → `CodeBlock`, etc.
- `"Settings"`, `"Animation"` — silently skipped (metadata, not rendered)

Layout types (`Row`, `Columns`, `Grid`, `Area`) have `RawChildren` that get recursively resolved into `Resolved []Child` during parsing. This builds the full tree in memory before rendering.

### Templ Components (`components/`)

Each content type has a corresponding `.templ` file. The dispatcher is `node.templ:ChildNode()` — it checks each union field and calls the matching component. `page.templ:HomePage()` iterates top-level children. `layout.templ:Layout()` is the HTML document shell.

**Generated files:** `*_templ.go` files are gitignored. Run `templ generate` after editing any `.templ` file.

### CSS (Tailwind 4.2)

Tailwind CSS 4.2 with custom theme. Built from `styles/app.css` → `static/app.css`.

- `styles/app.css` — Tailwind input: `@theme` (design tokens), `@layer base` (element styles), `@layer components` (btn, rich-text, code-block, menu)
- `static/app.css` — Generated output (gitignored)

Most components use Tailwind utility classes inline in `.templ` files. Complex multi-variant or descendant-selector patterns (buttons, rich-text, code-block, menu) use `@layer components` with `@apply`.

**Theme color naming:** `surface-*` (backgrounds), `fg`/`fg-muted`/`fg-faint` (text), `accent`/`accent-*` (brand), `edge`/`edge-*` (borders). Font sizes lg–4xl are larger than Tailwind defaults.

Dark theme only. System fonts only. See `STYLE_GUIDE.md` for design intent.

## Adding a New Content Block Type

1. Add the Go struct to `content/content.go`
2. Add a pointer field to the `Child` union struct
3. Add a `case` in `unmarshalChild()` mapping the API type name to the struct
4. If it has children, add a `resolve*` function and call it after unmarshal
5. Create `components/<name>.templ` with the rendering template (use Tailwind utility classes)
6. Add an `if child.<Name> != nil` branch in `components/node.templ:ChildNode()`
7. Run `just generate` (rebuilds CSS + templ)

## Adding a New Route/Page

1. Add a handler function in `render.go`
2. Add a fetch function in `fetch.go` (call `client.Content.GetPage` with the route slug)
3. Register the route in `server.go:newMux()`

## Dependencies

- `github.com/a-h/templ` — HTML templating (templ CLI must be installed: `go install github.com/a-h/templ/cmd/templ@v0.3.977`)
- `github.com/hegner123/modulacms/sdks/go` — ModulaCMS Go SDK for API calls
- `tailwindcss` (npm) — CSS framework. CLI at `/usr/local/bin/tailwindcss` (globally installed `@tailwindcss/cli`)
- `htmx` — included as static JS, used client-side
