# Home Page

---

## Hero
<!-- Component: CTA block, centered, full-width -->
<!-- Background: surface-0, generous vertical padding -->

**Heading:** ModulaCMS handles the infrastructure. You define the content.

**Subheading:** Self-hosted headless CMS. Single binary. No runtime. Starts in milliseconds.

**Button:** Get Started → /docs/quickstart

---

## Get Started
<!-- Component: CTA block, centered, narrow-width -->
<!-- Background: surface-1 or subtle accent border to break from previous section -->

**Heading:** Two commands. Running.

**Subheading:** `modula init && modula serve`

SQLite database, config file, all tables, bootstrap data. Open-source under AGPL-3.0. All features included. No tiers, no usage caps, no upgrade prompts.

**Button:** View on GitHub → https://github.com/hegner123/modulacms
**Button (secondary):** Read the Docs → /docs

---

## Features
<!-- Component: Grouped feature tables or card grid -->
<!-- Layout: Category headers with feature tables beneath each -->

### Core

| Feature | Description |
|---|---|
| Single binary | 27-29 MB. Three servers (HTTP, HTTPS, SSH), admin panel, TUI, plugin runtime, media pipeline, deploy engine. No runtime dependencies. |
| Multi-database | SQLite for local dev, MySQL or PostgreSQL for production. Switch with one config change. Stateless horizontal scaling. |
| SSL built in | Self-signed certs for dev, Let's Encrypt for production. No external tools. |
| One config file | Database, servers, auth, S3, email, CORS, plugins, deploy, observability. One JSON file. Diffable. Version-controllable. |

### Content

| Feature | Description |
|---|---|
| User-defined schema | Define datatypes and fields at runtime. Posts, products, calendars, anything. Add or remove fields without breaking existing content. |
| Content versioning | Every publish creates an immutable snapshot. Schedule publication. Restore any previous version. Full history from day one. |
| Localization | Field-level locale variants with BCP 47 codes and fallback chains. Translate fields, not entire entries. |
| Multi-format API | Output as clean JSON, or in formats compatible with Contentful, Sanity, Strapi, and WordPress. Migrate without rewriting your data layer. |

### Developer Tools

| Feature | Description |
|---|---|
| Four interfaces | CLI for automation, TUI over SSH for developers, admin panel for content editors, REST API for applications. Same data, pick your environment. |
| SDKs | TypeScript, Go, Swift. Typed responses, branded IDs, zero external dependencies. |
| Sandboxed plugins | Lua VMs with isolated database storage, operation budgets, circuit breakers, admin approval. Extend the CMS without compromising it. |
| Webhooks | 12 content lifecycle events. HMAC-SHA256 signed payloads. Automatic retry. |

### Operations

| Feature | Description |
|---|---|
| Media pipeline | S3-compatible storage. Automatic WebP conversion, focal point cropping, responsive dimension presets. Optimization at upload, not request time. |
| Deploy sync | Export and import content between instances. Dry-run preview. Hash-validated payloads. Schema version verification. |
| Audit trail | Every mutation logged: what changed, who, from where, when. |

---

## Schema
<!-- Component: Row > Two columns (text 50%, visual 50%) -->
<!-- Right column: Diagram of the five-table model or visual showing content composition chain -->

### How Content Works

ModulaCMS ships with one datatype (Pages) and one field (Title). Everything else is yours.

Five core tables model any content structure. Datatypes define what kinds of content exist. Fields define their properties. Content data holds instances. Content fields hold values. Routes assign content to URLs. All user-defined, all at runtime, all backed by relational tables with proper indexes.

A blog: define a "Blog" datatype, add fields to identify it (title, slug, author, publish date), then create child datatypes for the content that goes inside it. A rich text block, a featured image, a pull quote. Assign fields to each. Now you have a parent type with composable children you arrange however you want.

Same idea for a portfolio. Define a "Project" datatype with title, client name, year. Then assign child datatypes for case study sections, image galleries, testimonials. Define whatever the client asks for six months after launch. The schema adapts because it was designed to.

### Composable Content

Define content once, reference it everywhere. Menus, shared sections, testimonial blocks, any reusable content is authored in one place and composed into any page that needs it. Every reference uses a published snapshot, so you always know exactly what version is live. Change the source, and every page picks up the update on next publish.

---

## Schema Templates
<!-- Component: Row > Two columns (text 50%, list/grid 50%) -->
<!-- Right column: Grid or list of available template names with icons or short descriptions -->

Flexibility also means decision fatigue. Total freedom is great for about forty-five seconds, then you're staring at an empty datatypes list wondering where to start.

So we made some decisions for you. ModulaCMS ships with pre-built schema templates: blog, portfolio, e-commerce, documentation, FAQ, and more. Embedded in the binary, installable through the TUI. Pick one and you've got a working content structure with sensible fields in seconds.

But these aren't locked-in templates baked into the core. They're just datatypes and fields, the same as anything you'd build yourself. Don't need the excerpt field on your blog post? Remove it. Want to add a custom "mood" dropdown because the client's content team insists on tagging every article with a vibe? Add it. Gut the entire template and rebuild from scratch if you have opinions. The starter schemas get you to "something works" fast, then get out of your way.

---

## When the Schema Isn't Enough
<!-- Component: Row > Two columns (text 60%, image 40%) -->
<!-- Image: Diagram showing a Lua plugin with its own database tables alongside the CMS content tree. Or a terminal showing plugin approval in the TUI. -->

The content schema is intentionally flexible. But some problems need more than flexible content trees. They need structured data with direct relational mapping.

E-commerce is one. Orders, customers, products, inventory, shipping. These aren't website content. They're interconnected records with transactional integrity requirements that a content tree was never designed to handle. Forms are another. You need a schema for the form layout on the page, but you also need workflows that store and process submissions. The best place for that isn't inside your website content.

So we gave you the tools to build the services you need while still benefiting from the infrastructure. Sandboxed Lua plugins run inside the CMS with their own isolated database tables, operation budgets, circuit breakers, and admin approval. Build an order management system, a form processor, a booking engine. Your plugin gets its own storage, its own API endpoints, and the full protection of the CMS runtime. No separate service to deploy. No integration to maintain.

---

## AI-Ready
<!-- Component: Row > Two columns (text 60%, image 40%) -->
<!-- Image: Screenshot of a Claude/AI conversation managing content through the MCP server. Or a terminal showing the modula-mcp binary starting up. -->

Connect any AI assistant to your CMS. Draft content, build schemas, manage media, configure permissions. All through natural language.

ModulaCMS includes `modula-mcp`, a Model Context Protocol server that works with any compatible AI assistant. Point it at any environment and control how much access the AI has. Let it draft blog posts on your local instance but only read content in production. Give it full schema access in development so you can describe the content model you want and let the model build it. Or lock it down to read-only. The permission model is the same RBAC system that controls every other interface.

Content, schema, routes, media, users, roles, plugins, deploy sync, configuration. Everything the TUI and admin panel can do, an AI assistant can do through MCP. Two environment variables to configure: `MODULACMS_URL` and `MODULACMS_API_KEY`.

---

## Developer-Friendly Auth
<!-- Component: Row > Two columns (text 60%, image 40%) -->
<!-- Image: Terminal showing `ssh modula` connecting directly to the TUI. Or a split showing auth methods side by side. -->

`ssh modula` and you're in. No password. No 2FA prompt. Your SSH key is your key to every instance you manage.

No more keeping a spreadsheet of thirty client logins with different username conventions and expired passwords. If a client needs a content update, SSH drops you into the TUI. Make the change, get back to your work.

Native email and password authentication works out of the box. Full spec-compliant OAuth connects to any provider: Google, GitHub, Azure, Okta, Auth0, your company's homegrown identity server. If it speaks OAuth, ModulaCMS speaks back.

API tokens are just as simple. Generate one from the TUI, drop it in your local `config.json`, and use `modula connect` to run the TUI on your machine against any deployed ModulaCMS API. Manage content and media, configure schema. All from your terminal. No browser, no VPN.

---

## Security
<!-- Component: Row > Two columns (text 60%, image 40%) -->
<!-- Image: Diagram showing the layered security model: rate limiter → auth middleware → permission guards → parameterized queries. Or a terminal showing a 403 response with audit log output. -->

Every query is parameterized. ModulaCMS uses sqlc to generate all database access from typed SQL files. No query is ever constructed from string concatenation. The generated code enforces parameterized placeholders across all three database backends (SQLite, MySQL, PostgreSQL). SQL injection is not a category of bug that exists in this codebase.

SSH keys are stored with SHA256 fingerprints and a UNIQUE constraint at the database level. Key validation runs through the same cryptographic library that powers the SSH server. One key per identity, verified on every connection.

Every protected endpoint passes through authentication middleware before any handler runs. Session cookies are HTTP-only with 24-hour expiration. API requests fall back to Bearer token validation. If neither is present, the request is denied. Fail-closed by default.

Behind authentication sits a permission guard system with 47 discrete permissions mapped to roles. `RequirePermission` enforces a single permission. `RequireAnyPermission` and `RequireAllPermissions` handle OR and AND logic. `RequireResourcePermission` auto-maps HTTP methods to operations: GET becomes read, POST becomes create, PUT becomes update, DELETE becomes delete. Permission lookups are O(1). If no permission set is found in the request context, the response is 403. No fallthrough. No defaults.

Auth endpoints are rate-limited at 10 requests per minute per IP using a token bucket algorithm. The limiter is proxy-aware, checking forwarded headers before falling back to the remote address. Stale entries are evicted every 10 minutes.

Docker deployments run as a non-root user (UID 1000:1000). The multi-stage build compiles with CGO for native SQLite and WebP support, then strips the toolchain from the runtime image. Persistent volumes isolate data, certificates, SSH keys, backups, and plugins. Health checks run on all services. The production compose includes Caddy as a reverse proxy, PostgreSQL, and MinIO object storage.

---

## Footer CTA
<!-- Component: CTA block, centered, full-width -->
<!-- Background: surface-1 or accent-tinted -->

**Heading:** Two commands. See for yourself.

**Subheading:** `modula init && modula serve`

**Button:** Get Started → /docs/quickstart
**Button (secondary):** Read the Story → /story
