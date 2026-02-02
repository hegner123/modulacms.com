# ModulaCMS Product Brief

## One-liner

ModulaCMS lets developers work with their data in the way that best suits their workflow — advanced CI/CD pipelines, Dockerfiles, quick local development, client-friendly UI, or direct access via TUI.

## What ModulaCMS Is

A free, open-source (AGPL-3.0), single-binary headless CMS written in Go. It is a thin, performant layer between your database and your frontend. It handles content infrastructure so your frontend framework can do what it was designed to do — without conflict, duplication, or compromise.

ModulaCMS provides a performant backend that never competes with your frontend framework. Every feature is opt-in through a config file. Nothing is forced on you.

Start with a REST API. Enable what you need. Nothing more.

## Core Architecture

- **Single binary** — no runtime dependencies, no Node.js, no npm
- **Stateless** — deploy behind a load balancer with PostgreSQL + S3, it just works
- **Multi-database** — SQLite for local development, MySQL/PostgreSQL for production
- **REST API** — JSON responses, framework-agnostic, language-agnostic
- **Multi-format output** — transform API responses to match Contentful, Sanity, Strapi, WordPress, or clean JSON formats

## Content Model

ModulaCMS comes with one root datatype (Pages) and one field (Title). Everything else is user-defined.

**System-defined (ModulaCMS owns these):**

- Users, roles, sessions, tokens
- OAuth, SSH keys
- Media, media dimensions
- Backups, audit trails
- Change events

**User-defined (you own these):**

- Routes — URL structure and navigation
- Datatypes — what kinds of content exist (posts, products, calendars, quizzes, anything)
- Fields — what properties each datatype has
- Content Data — the actual content entries
- Content Fields — the field values per entry
- Admin counterparts of all the above — how the admin panel presents and organizes content

Posts, Comments, Catalogs, Calendars, Schedules, Promotions, Quizzes — all first-class datatypes. No "custom post types" bolted onto a blog engine.

## Four Interfaces

| Interface | Audience | Use Case |
|-----------|----------|----------|
| CLI | DevOps, CI/CD | Scripting, automation, Docker, pipelines |
| TUI (SSH) | Developers | Content management, schema configuration, database operations, schema template installation |
| Admin Panel | Content editors, clients | Browser-based content management |
| REST API | Applications | Programmatic access to everything |

The TUI is a superset of the CLI — every CLI command is available as a selectable action in the TUI. The admin panel is independently configurable from the content schema.

## Developer Experience

### Getting Started

```
modulacms serve
```

First run: auto-creates default config, initializes SQLite database, creates tables, starts server. One command.

### Interactive Setup

```
modulacms serve --wizard
```

Interactive configuration for database selection, S3 setup, OAuth, SSL.

### SSH Management

```
ssh modulacms
```

Full TUI for content management, schema configuration, database operations. SSH key authentication.

### CLI Operations

```
modulacms db init
modulacms db wipe
modulacms db reset
modulacms db backup
```

Infrastructure commands for scripting and automation.

### Built-in SSL

- Local development: automatic self-signed certificate generation
- Production: built-in Let's Encrypt integration

### Shared Media

Native S3 support means teams share a media bucket. No downloading assets for local development.

### Schema Templates

Pre-built datatype/field configurations (blog, portfolio, e-commerce, docs) embedded in the binary via Go embed. Installed through the TUI — browse, preview, select, populate.

## What ModulaCMS Does NOT Do (By Design)

- No server-side rendering
- No caching layer
- No native forms
- No webhooks
- No cron jobs
- No native extensions (Lua plugins planned)
- No frontend opinions

These are not missing features. Your frontend framework handles caching, SSR, forms, and scheduled tasks. ModulaCMS does not duplicate or conflict with these capabilities.

**Your framework already solves these problems. Your CMS shouldn't solve them again, differently.**

## The Problem ModulaCMS Solves

### CMS Rigidity

The hardest part about working with a CMS is how rigid their database structure is. If you have a clear understanding of your site's needs, picking the right CMS is simple. But if you're growing and your needs grow, no other CMS gracefully handles a growing company's needs.

With ModulaCMS, turn around new features with new admin panels weeks faster than custom plugin building, or writing complex adapters to fit your feature into a data schema it wasn't designed for.

It's no longer a question of "how can we make X do Y." It's a question of "what's the best way to adapt ModulaCMS to Y."

### CMS/Framework Conflict

Modern CMSes are too full-featured. They duplicate what your frontend framework already does, but differently and incompatibly:

- **Routing** — your CMS and your framework both want to own URL structure
- **Auth** — two session systems that need bridging
- **Caching** — two invalidation strategies that conflict
- **Forms** — server-side processing split across two systems
- **Image optimization** — duplicate pipelines, wasted effort

ModulaCMS eliminates the overlap. The CMS provides data. The framework handles everything else. No conflicts, no duplication.

### The ACF Problem

WordPress with heavy Advanced Custom Fields usage becomes unusable. Every ACF field is a row in wp_postmeta — a key-value table. A content type with 30 fields means 30+ joins per query through WP_Query, which was designed for blog posts, not complex structured content. Building an entire admin experience on ACF causes massive lags and editing delays.

ModulaCMS stores content fields in dedicated tables with proper relationships. The data model was designed for complex content from day one.

## Competitive Position

### vs SaaS (Contentful, Sanity)

Free, self-hosted, no usage caps. No per-user or per-API-call pricing. Contentful caps at 5 users and 25K records on free tier. Sanity caps at 3 users and 500K API calls/month. ModulaCMS: all features, no limits.

### vs Payload

No paid tiers — features are not gated behind enterprise pricing. No Node.js runtime required. SSH TUI for terminal management. Not dependent on acquisition-driven product direction (Figma).

### vs Strapi

Single binary vs Node.js process. Go performance. SSH TUI. No runtime dependencies.

### vs WordPress

Modern architecture, truly flexible schema (not custom post types on a blog engine), stateless horizontal scaling, no plugin dependency hell.

### Unique to ModulaCMS

- SSH TUI — no other CMS offers terminal-native management
- Admin panel independence — admin UI is configured separately from content schema
- Single binary — no runtime, no package manager, no dependencies
- Multi-format API — output compatible with Contentful/Sanity/Strapi/WordPress formats
- All features, no tiers, free forever

## Business Model

- **Free:** ModulaCMS binary (all features), default admin panel, default client theme, documentation, starter templates
- **Paid:** Premium admin panels (industry-specific), premium client themes (polished, themeable frontends)
- **Distribution:** Source zip files with purchase codes
- **License:** AGPL-3.0

## Marketing Approach

Developer-first. Technical documentation, not CMS marketing fluff. Code examples front and center.

### Site Structure (modulacms.com)

1. **What it is** — one sentence, code example of hitting the API
2. **Install** — `modulacms serve`, done
3. **SSH TUI demo** — gif/video of the terminal UI (this is the hook)
4. **What you need to get started** — short list, emphasize simplicity
5. **What you can do with ModulaCMS** — progressive disclosure of features
6. **API reference** — endpoints, output formats
7. **Starter templates** — pick your framework, clone, connect
8. **Schema definition** — how to define datatypes and fields
9. **CLI reference** — every command, flags, exit codes
10. **TUI guide** — walkthrough with screenshots/gifs

Lead with the TUI. It's the thing that makes people stop scrolling.

### Starter Templates

| Project | Stack | Purpose |
|---------|-------|---------|
| modulacms.com | Astro | Product site, powered by own CMS, JS ecosystem demo |
| starter-php | PHP (Laravel or vanilla) | Server-rendered, WordPress migration audience |
| admin.modulacms.com | Next.js | Already built |

Two starters in different language ecosystems prove "any framework" more effectively than five JavaScript starters.

### Key Marketing Lines

- "ModulaCMS lets developers work with their data in the way that best suits their workflow."
- "Start with a REST API. Enable what you need. Nothing more."
- "Your framework already solves these problems. Your CMS shouldn't solve them again, differently."
- "All features. One binary. No tiers, no cloud lock-in, no upgrade prompts. Ever."
- "ModulaCMS handles the infrastructure. You define the content."
- "Single binary. No runtime. Stateless. Deploy it like infrastructure, not like a CMS."
- "It's no longer how can we make X do Y. It's what's the best way to adapt ModulaCMS to Y."

## Technical Details

- **Language:** Go 1.24
- **CLI Framework:** Cobra
- **TUI Framework:** Charmbracelet Bubbletea
- **SSH Server:** Charmbracelet Wish
- **Database Access:** sqlc (type-safe, no ORM)
- **IDs:** ULID (26-char strings)
- **Auth:** Session cookies + Bearer token API keys + OAuth (PKCE)
- **Media:** S3-compatible storage, image optimization (WebP)
- **Plugins:** Lua via gopher-lua (planned)
- **Schema Templates:** Go embed
- **SSL:** Self-signed (dev) + Let's Encrypt (production)
- **Observability:** Optional Sentry/Datadog integration

## Shipped vs Planned

### Shipped

- Single binary distribution
- User-defined datatypes, fields, routes
- Independent admin panel configuration
- SSH TUI management
- CLI database operations (init, wipe, reset, backup)
- Local SSL cert generation
- Production Let's Encrypt
- S3 media storage
- Image optimization
- Multi-database (SQLite, MySQL, PostgreSQL)
- Multi-format API output (Contentful, Sanity, Strapi, WordPress, clean, raw)
- OAuth + session auth
- Content import (WordPress, Contentful, Sanity, Strapi)
- Install wizard
- Role-based access control

### Planned

- Lua plugin system
- Cross-environment content migration (import/export between instances)
- Schema templates (Go embed, TUI-installable)
- One-command startup (auto-detect and initialize on first `modulacms serve`)
- Additional docker-compose configurations
