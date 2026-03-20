# SITE_ROADMAP.md — Page Structure and Content Strategy

## Pages

### Home Page (`/`)
**Voice:** VOICE_MARKETING.md
**Tone:** Neutral

Explain what ModulaCMS is and what it does. No comparisons to competitors. No opinions about the CMS industry. Just features, capabilities, and what the developer gets.

Sections:
- Hero: What it is, one-liner, get started CTA
- Feature overview: What it does (content model, API, interfaces, media, plugins, etc.)
- How it works: The developer experience from install to serving content
- Get started CTA

Don't try to explain everything. Give the reader enough to understand the product and a reason to explore further.

### Story Page (`/story`)
**Voice:** VOICE_BLOG.md
**Tone:** Opinionated, direct

This is where ModulaCMS has opinions. Name specific competitors. Call out specific pain points. Explain the WHY behind every design decision.

Content:
- Why ModulaCMS exists: the problems with the current CMS landscape
- Specific pain points in specific products (WordPress ACF bloat, Contentful pricing caps, Strapi's Node.js dependency, Payload's tier-gating)
- Design decisions and the reasoning behind them (why single binary, why no SSR, why Lua not Node for plugins, why stateless, why opt-in features)
- What ModulaCMS intentionally does not do, and why those boundaries make it better
- Comparison tables with specific numbers (pricing, binary size, dependencies, user limits)

This page answers: "Why should I use this instead of what I'm already using?"

### Templates Page (`/templates`)
**Voice:** VOICE_MARKETING.md
**Tone:** Neutral

Starter templates for different frameworks. Show what's available, what stack each uses, and how to get started with each one.

### Documentation (`/docs/*`)
**Voice:** VOICE_DOCUMENTATION.md

All technical reference material: getting started guides, API reference, CLI reference, configuration, content model, plugin development, SDK usage, TUI guide.

### Blog (`/blog/*`)
**Voice:** VOICE_BLOG.md

Product updates, case studies, technical deep-dives, changelog narratives. Problem-first, narrative structure.

### Changelog (`/changelog`)
**Voice:** VOICE_DOCUMENTATION.md

Factual, structured. Added/Changed/Fixed/Removed. No editorializing.
