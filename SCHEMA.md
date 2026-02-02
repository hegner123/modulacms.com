# modulacms.com Content Schema

Each route returns a single content tree. The root node owns child nodes via `nodes[]`. No relation IDs between datatypes — parent/child is expressed by nesting.

## Content Trees

```
/ (landing)
Page
├── Section (hero)
├── Section (install)
│   ├── CodeExample
│   └── CodeExample
├── Section (tui_demo)
├── Section (features)
│   ├── Feature
│   ├── Feature
│   └── ...
├── Section (comparisons)
│   ├── Comparison
│   ├── Comparison
│   └── ...
└── Section (cta)

/templates
Page
├── StarterTemplate
└── StarterTemplate

/docs/getting-started
Doc
├── CodeExample
└── CodeExample

/docs/api
Doc
├── APIEndpoint
├── APIEndpoint
└── APIEndpoint

/docs/cli
Doc
├── CLICommand
├── CLICommand
└── CLICommand

/changelog
Page
├── Changelog
└── Changelog
```

## Routes

| Route | Root Datatype | Children | Fetch |
|-------|---------------|----------|-------|
| `/` | Page | Section > Feature, Comparison, CodeExample | `GET /api/v1/routes/landing?format=clean` |
| `/templates` | Page | StarterTemplate | `GET /api/v1/routes/templates?format=clean` |
| `/docs/:slug` | Doc | CodeExample | `GET /api/v1/routes/docs/:slug?format=clean` |
| `/docs/api` | Doc | APIEndpoint | `GET /api/v1/routes/docs/api?format=clean` |
| `/docs/cli` | Doc | CLICommand | `GET /api/v1/routes/docs/cli?format=clean` |
| `/changelog` | Page | Changelog | `GET /api/v1/routes/changelog?format=clean` |

## Datatypes & Fields

### Page

Root datatype for landing, templates, and changelog.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | page title |
| slug | text | yes | URL path segment |
| description | text | no | meta description, OG tags |
| body | richtext | no | general page content |

**Children:** Section, StarterTemplate, Changelog (depends on page purpose)

### Section

Landing page content block. Renders differently based on `section_type`.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | internal label |
| section_type | text | yes | hero, install, tui_demo, features, comparisons, cta |
| heading | text | no | display heading |
| subheading | text | no | |
| body | richtext | no | |
| media_url | text | no | image, gif, video path |
| media_alt | text | no | |

**Parent:** Page
**Children:** Feature, Comparison, CodeExample (depends on section_type)

### Feature

Feature card nested under a features Section.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | feature name |
| description | text | yes | short explanation |
| icon | text | no | icon name or svg reference |
| category | text | no | architecture, dx, interfaces, content_model |

**Parent:** Section (section_type: features)
**Children:** none

### Comparison

Competitive comparison entry nested under a comparisons Section.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | competitor name |
| competitor_type | text | yes | saas, oss, legacy |
| pain_point | text | yes | what they get wrong |
| modulacms_answer | text | yes | how ModulaCMS solves it |

**Parent:** Section (section_type: comparisons)
**Children:** none

### CodeExample

Code snippet nested under a Section or Doc.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | label for the snippet |
| language | text | yes | bash, json, go, js, php, etc. |
| code | text | yes | raw code content |
| description | text | no | what this snippet demonstrates |

**Parent:** Section (section_type: install) or Doc
**Children:** none

### Doc

Documentation article. Root datatype for docs routes.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | article title |
| slug | text | yes | URL path under /docs/ |
| category | text | yes | getting-started, guides, reference, concepts |
| body | richtext | yes | article content |
| description | text | no | meta/preview text |
| prev_doc_id | number | no | previous article ID for navigation |
| next_doc_id | number | no | next article ID for navigation |

**Children:** CodeExample, APIEndpoint, CLICommand (depends on doc purpose)

### APIEndpoint

API reference entry nested under an API docs page.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | e.g. "List Content" |
| method | text | yes | GET, POST, PUT, DELETE |
| path | text | yes | /api/v1/content/:datatype, etc. |
| description | text | yes | what this endpoint does |
| parameters | text | no | query/path params description |
| request_body | text | no | example JSON request |
| response_body | text | yes | example JSON response |
| auth_required | text | yes | none, session, token |

**Parent:** Doc (category: reference, slug: api)
**Children:** none

### CLICommand

CLI reference entry nested under a CLI docs page.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | e.g. "db init" |
| command | text | yes | full command string |
| description | text | yes | what it does |
| flags | text | no | available flags and descriptions |
| examples | text | no | usage examples |
| exit_codes | text | no | non-zero exit code meanings |

**Parent:** Doc (category: reference, slug: cli)
**Children:** none

### StarterTemplate

Template showcase entry nested under the templates Page.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | e.g. "Astro Starter" |
| stack | text | yes | Astro, PHP, Next.js |
| repo_url | text | yes | GitHub URL |
| description | text | yes | what it demonstrates |
| language_ecosystem | text | yes | javascript, php |
| preview_url | text | no | live demo link |
| media_url | text | no | screenshot |

**Parent:** Page (slug: templates)
**Children:** none

### Changelog

Release entry nested under the changelog Page.

| Field | Type | Required | Note |
|-------|------|----------|------|
| title | text | yes | version or release name |
| date | text | yes | release date |
| body | richtext | yes | what changed |
| version | text | yes | semver |
| tag | text | no | major, minor, patch, beta |

**Parent:** Page (slug: changelog)
**Children:** none
