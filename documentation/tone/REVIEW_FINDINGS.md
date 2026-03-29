# Documentation Review Findings

Weaknesses identified by comparing ModulaCMS documentation against the patterns in DOCUMENTATION_VOICE.md (derived from Next.js, Stripe, and Figma docs).

## Critical

### 1. Documentation is written from the system's perspective, not the user's

This is the fundamental problem that underlies most other findings. The documentation explains **how ModulaCMS works** rather than **how to use ModulaCMS**. It reads like internal architecture docs that were published externally.

The `documentation/` directory is for end users -- developers integrating ModulaCMS into their projects. Internal system design belongs in the `ai/` directory. Users care about "how do I do X" not "here's how the system implements X under the hood."

**Examples of system-perspective writing (current):**

- content-model.md explains database table relationships, Go struct field names, ULID encoding details, and column types -- information a contributor needs, not an integrator
- tree-structure.md teaches the sibling-pointer algorithm, O(1) complexity analysis, and the tree assembly query -- a user just needs to know content is hierarchical and how to nest/reorder it
- audit-trail.md describes the `change_events` table schema, transaction atomicity guarantees, and recorder interfaces -- a user needs to know that changes are tracked and how to query the audit log
- observability.md lists internal metric counter names, middleware chain ordering, and instrumentation implementation -- a user needs to know how to connect their monitoring tool
- rbac.md shows the `PermissionCache` struct, lock-free read patterns, and middleware function signatures -- a user needs to know how to create roles and assign permissions
- publishing-lifecycle.md details the `content_versions` table structure and snapshot JSON format -- a user needs to know how to publish, schedule, and restore

**What industry leaders do:**

Next.js "Layouts and Pages" says: "A **page** is UI that is rendered on a specific route. To create a page, add a page file inside the `app` directory." It never mentions React fiber reconciliation, the internal component tree, or how the framework resolves file-system routes to handlers. The user learns what to do, not how the framework does it.

Stripe "Payments" says: "Build a payment form or use a prebuilt payment page to accept online payments." It never explains Stripe's internal payment processing pipeline, database sharding strategy, or fraud detection algorithms.

**The test for every paragraph:** "Does the reader need to know this to USE the feature, or does this explain how the feature is IMPLEMENTED?" If it's implementation, move it to `ai/` or remove it.

**Recommendation:** Rewrite concept pages around the user's mental model. Each concept page should answer three questions: (1) What is this? (one sentence, user-facing), (2) What can I do with it? (capabilities list), (3) How do I use it? (link to the relevant guide). Implementation details like table schemas, Go structs, algorithm complexity, and internal type systems should be stripped from external docs entirely.

**Pages most affected (ranked by severity):**

| Page | Problem |
|------|---------|
| concepts/tree-structure.md | Teaches the algorithm instead of the mental model |
| concepts/content-model.md | Shows database tables and Go structs instead of user workflows |
| concepts/audit-trail.md | Describes internal schema instead of "what you can audit and how" |
| concepts/observability.md | Lists internal metric names instead of "how to monitor your CMS" |
| concepts/rbac.md | Shows middleware internals instead of "how to manage access" |
| concepts/publishing-lifecycle.md | Details snapshot storage instead of "how to publish and version content" |
| concepts/media-pipeline.md | Less affected -- mostly user-facing already |
| concepts/localization.md | Less affected -- mostly user-facing already |
| guides/content-trees.md | Mixes user tasks with implementation details |
| guides/authentication.md | Leaks middleware function signatures and cache internals |
| guides/content-modeling.md | Shows Go field types alongside user-facing field creation |

### 2. PHILOSOPHY.md is a sales manifesto, not documentation

*Note: This page was written for fun and is exempt from voice rules. Kept here for completeness but not a revision target.*

This is the biggest tone mismatch in the entire documentation set. The page reads like a founder's blog post or a pitch deck -- competitive jabs ("whatever opinionated framework decided"), war stories ("from scars"), hyperbolic claims ("screaming fast"), and casual/aggressive language ("No exit interview. No data hostage negotiation.").

**Voice rule violations:**
- Uses "we" to refer to the software throughout ("we're giving them the choice", "we never force a pattern")
- Marketing language in technical docs ("fast-as-hell JSON pumping machine", "will make you double-check your latency metrics")
- Competitive comparisons ("The Node.js runtime alone -- before you install a single package -- is 80 MB")
- Exclamation-adjacent intensity without exclamation marks
- Sections run 20-30+ lines without heading breaks
- Mixes concept explanation with opinion pieces

**What industry leaders do instead:** Next.js explains what it is in 2 sentences. Stripe says "Use Stripe to start accepting payments." Neither attacks competitors or tells war stories. Philosophy pages in leading docs are short, value-driven, and factual.

**Recommendation:** Rewrite as a concise philosophy page. State the three values, explain each in 3-5 sentences with concrete examples (not competitor attacks), and keep the "Data Authority" section which is genuinely strong architectural documentation. Move the rest to a blog post or marketing site.

### 2. No standardized callout pattern

The voice guide defines two callout types: **"Good to know"** and **"Tip"**. Zero documentation pages use either pattern. Instead, pages dump miscellaneous information into a "Notes" bullet list at the bottom.

**What industry leaders do:** Next.js uses `> **Good to know**:` blockquotes inline, right after the relevant section. Readers see the supplementary info exactly when they need it, not buried in a catch-all section at the bottom they may never reach.

**Pages affected:** Nearly every guide, concept, and reference page ends with a "Notes" section.

**Recommendation:** Audit every "Notes" section. Move each bullet to an inline `> **Good to know**:` callout after the relevant section. Delete the catch-all "Notes" heading.

### 3. Passive voice is pervasive in concept docs

The voice guide says: "Use present tense and active voice." Concept pages rely heavily on passive constructions:

- "Change events are recorded atomically" -> "ModulaCMS records change events atomically"
- "Content is served through slug-based routing" -> "The API serves content through slug-based routing"
- "Filters are passed as query parameters" -> "Pass filters as query parameters"
- "Metrics are collected across three layers" -> "ModulaCMS collects metrics across three layers"
- "Variants are generated during image optimization" -> "The optimization pipeline generates variants"

This pattern repeats across every concept page (audit-trail, content-model, media-pipeline, observability, publishing-lifecycle, rbac, tree-structure, localization). The guides are better but still have frequent passive constructions.

**Recommendation:** Rewrite concept pages to lead with the actor (ModulaCMS, the API, the pipeline, you) rather than the object being acted upon.

## High

### 4. README.md is a wall of links, not a landing page

The README has 72 links organized by section. Compare to Next.js docs landing page: a 3-section overview ("Getting Started", "Guides", "API Reference") with one sentence per section explaining what it's for, then one or two links.

The current README:
- Opens with a dense paragraph instead of a one-sentence summary
- Lists every single documentation page (72 links)
- No visual hierarchy distinguishing "start here" from "deep reference"
- No quick start or "first 3 steps" pattern

**Recommendation:** Cut the README to essential navigation. Keep 3-5 sections with 2-4 links each. Add a "Quick Start" section with 3 numbered steps (clone, init, serve). Move the comprehensive index to a separate SITEMAP.md or let readers discover pages through cross-references.

### 5. Dense sections without heading breaks

Several pages have sections exceeding 15 lines of prose without sub-headings:

| Page | Section | Approx. lines |
|------|---------|---------------|
| guides/deploy-sync.md | Importing Content | ~70 |
| guides/building-admin-interfaces.md | Content Management | ~50 |
| plugins/security.md | Lua Sandbox | ~50 |
| guides/authentication.md | Bootstrap Roles | ~30 |
| guides/media-management.md | Configuration | ~45 |
| guides/plugins.md | Configuration | ~45 |

**Recommendation:** Break these into sub-sections with H3 headings. Readers scan headings, not paragraphs.

### 6. Quickstart and Installation overlap significantly

Both pages contain the same build-from-source instructions (git clone, just build, cp to PATH, modula version). A reader following the recommended path (Quickstart first) hits the same content again in Installation.

**What industry leaders do:** Next.js has one Installation page with a "Quick start" section at the top (3 steps), then detailed options below. No separate quickstart page.

**Recommendation:** Either merge into one page with a quick start section at the top, or make the Quickstart truly minimal (link to Installation for build steps, focus only on the 3-step path).

### 7. MEDIA.md is a standalone duplicate

`documentation/MEDIA.md` exists at the documentation root alongside `guides/media-management.md`. Both cover media uploads, dimensions, focal points, and API endpoints. The README links to both. This creates confusion about which is canonical.

**Recommendation:** Pick one canonical location. If MEDIA.md is the comprehensive reference, move it to `reference/media.md`. If media-management.md is the guide, make MEDIA.md redirect or remove it. One source of truth per topic.

## Medium

### 8. Multi-sentence summaries instead of single-sentence

The voice guide says: "Start every page with an H1 title followed by a single-sentence summary." Most pages have 2-3 sentence summaries:

- content-model.md: "ModulaCMS uses a schema-first content model. You define **datatypes**... then create **content data** nodes..."
- observability.md: "ModulaCMS collects metrics across three layers... Observability is opt-in: metrics are always collected internally, but external reporting requires explicit configuration."

These aren't bad -- they're informative. But compare to Next.js: "Create a new Next.js app and run it locally." One sentence. The reader knows immediately what this page is for.

**Recommendation:** Tighten each opening to one sentence that answers "what does this page help me do?" Follow with a second paragraph if needed for context.

### 9. Configuration page is a 393-line monolith

configuration.md is a single page with 12+ table sections covering every config field. This is comprehensive but not scannable. A reader looking for "how do I set up OAuth" has to scroll past server settings, database settings, S3, CORS, and cookies to find it.

**What industry leaders do:** Stripe breaks configuration into separate pages per topic. Next.js has `next.config.js` as a top-level reference but links to individual feature pages for details.

**Recommendation:** Keep the single page as a quick-reference index (field name, type, default, one-line description). Link each section heading to a dedicated page or the relevant guide that explains the feature in context.

### 10. Code examples sometimes lack file path labels

The voice guide says: "Label every code block with its file path or context." Most pages do this well, but some concept docs show Go struct definitions or JSON without labeling the source:

- content-model.md shows Go structs without file paths
- tree-structure.md shows JSON tree structures without context labels
- Some curl examples omit the "Terminal" label that Next.js consistently uses

**Recommendation:** Add a label line before every code block: the file path for source code, "Terminal" for shell commands, "Response" for API responses, the config filename for JSON configs.

### 11. No "What to use and when" decision guides

Next.js includes "What to use and when" sub-sections that help readers choose between options (e.g., `searchParams` vs `useSearchParams`). ModulaCMS has similar choice points that lack this pattern:

- SQLite vs MySQL vs PostgreSQL (mentioned in config, never compared)
- Session cookies vs API tokens vs SSH keys (all in authentication, no decision guide)
- Admin content vs public content (explained separately, never compared side-by-side)
- Read-only SDK vs Admin SDK (overview page lists both but doesn't help you choose)

**Recommendation:** Add "When to use what" sub-sections at each decision point, with a short bullet list explaining the tradeoff.

### 12. Heading style inconsistency between content types

The voice guide distinguishes task-oriented headings for guides ("Create a datatype") from noun-phrase headings for reference ("Configuration options"). Most pages follow this, but some mix styles:

- authentication.md uses "Password Login" (noun) alongside "Create custom roles" (verb)
- admin-panel.md uses "Notes" (vague) as a heading

**Recommendation:** Audit headings. Guides get verb-first headings. Reference pages get noun-phrase headings. Never use "Notes", "Miscellaneous", or "Other".

## Low

### 13. fetching-content.md lacks next steps

The examples/fetching-content.md page ends without cross-references or next steps. Every other example page has them.

### 14. No explicit "Prerequisites" callout pattern

Installation and quickstart list prerequisites but not in a visually distinct format. Next.js uses a clear "System requirements" section with a bulleted list. ModulaCMS has this but buries it differently in each page.

### 15. Glossary could be more discoverable

The glossary exists but is buried under Reference. Terms defined in the glossary are not linked from the pages where they first appear. Industry-leading docs hyperlink terms to their glossary definitions on first use per page.

---

## Strengths (what you're already doing well)

These patterns are consistent with or better than the industry references:

- **Heading hierarchy** -- no level skipping across all 72 files
- **No hedging language** -- "might", "perhaps", "maybe" are absent
- **No filler transitions** -- "As mentioned above", "It is worth noting" not found anywhere
- **No "please"** -- imperative mood used correctly
- **No "simply/just/easy"** -- absent across all files
- **Realistic code examples** -- real slugs, real IDs, real config values (not "foo"/"bar")
- **Consistent structure within content types** -- guides follow the same pattern, SDK docs mirror each other across Go/TypeScript/Swift
- **Tables for reference data** -- used effectively and consistently
- **Cross-references at natural points** -- most pages link to related content inline
- **SDK parity** -- all three SDKs (Go, TypeScript, Swift) follow identical documentation structures
