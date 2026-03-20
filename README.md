# modulacms.com

Marketing and documentation website for [ModulaCMS](https://github.com/modulacms/modulacms) -- a free, open-source, single-binary headless CMS written in Go.

Go server that fetches content from the ModulaCMS REST API and renders HTML server-side with [templ](https://templ.guide) components. Vanilla CSS, no JS build step.

## Prerequisites

- Go 1.25+
- [templ CLI](https://templ.guide/quick-start/installation)
  ```sh
  go install github.com/a-h/templ/cmd/templ@v0.3.977
  ```
- A running ModulaCMS instance to serve content from

## Setup

Create a `.env` file:

```sh
CMS_BASE_URL=https://api.modulacms.com
CMS_API_KEY=your-api-key
```

## Run

```sh
templ generate
source .env && go run .
```

Server starts at `localhost:5050`.

## Build

```sh
templ generate
go build -o modulacms.com .
./modulacms.com
```

## Docker

```sh
docker build -t modulacms-site .
docker run -p 5050:5050 -e CMS_BASE_URL=... -e CMS_API_KEY=... modulacms-site
```

## Environment Variables

| Variable | Required | Default | Description |
|:---------|:---------|:--------|:------------|
| `CMS_BASE_URL` | Yes | -- | ModulaCMS API base URL |
| `CMS_API_KEY` | No | -- | API key for authenticated requests |
| `PORT` | No | `5050` | HTTP listen port |

## Content Source

Content is fetched from the ModulaCMS REST API in clean format:

```
GET /api/v1/routes/{route}?format=clean
```

Each route returns a nested content tree. See `SCHEMA.md` for the content model and `SCHEMA_EXAMPLES.json` for example API responses.

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
- `SCHEMA_EXAMPLES.json` -- example API responses (if present)
- `STYLE_GUIDE.md` -- visual design tokens and CSS conventions

## License

AGPL-3.0
