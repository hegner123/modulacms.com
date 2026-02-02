# modulacms.com

Marketing and documentation website for [ModulaCMS](https://github.com/modulacms/modulacms) -- a free, open-source, single-binary headless CMS written in Go.

Built with [Astro](https://astro.build). Content served by its own CMS backend via REST API.

## Setup

```sh
npm install
npm run dev
```

Dev server runs at `localhost:4321`.

## Commands

| Command | Action |
|:--------|:-------|
| `npm run dev` | Start dev server at `localhost:4321` |
| `npm run build` | Production build to `./dist/` |
| `npm run preview` | Preview production build locally |

## Content Source

Content is fetched from the ModulaCMS REST API using the clean format:

```
GET /api/v1/routes/{route}?format=clean
```

Each route returns a nested content tree. See `SCHEMA.md` for the full content model and `SCHEMA_EXAMPLES.json` for example API responses.

## Site Structure

| Route | Description |
|:------|:------------|
| `/` | Landing page -- hero, install, TUI demo, features, comparisons, CTA |
| `/templates` | Starter templates for different frameworks |
| `/docs/:slug` | Documentation articles |
| `/changelog` | Release history |

## Related

- [ModulaCMS backend](../modulacms/) -- the CMS binary this site consumes
- `PRODUCT_BRIEF.md` -- product strategy and marketing copy
- `SCHEMA.md` -- content model definition
- `SCHEMA_EXAMPLES.json` -- example API responses

## License

AGPL-3.0
