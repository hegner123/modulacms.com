# Blog

Product updates, technical decisions, and lessons from building ModulaCMS.

<!-- This page serves as the blog index. Individual posts will live as separate content entries in the CMS, rendered on /blog/:slug routes. The content below defines the blog landing page itself. -->

## Voice

All blog content follows VOICE_BLOG.md. Problem-first, opinionated, honest about trade-offs.

## Planned Posts

### Launch and Product

- Why we built another CMS (condensed version of the story page, focused on the personal motivation)
- The single binary decision and what it costs
- How the SSH TUI works under the hood (Bubbletea + Wish architecture)

### Technical Deep-Dives

- Content versioning with immutable snapshots: how and why
- Sandboxed Lua plugins: building a safe extension system
- Multi-format API output: making migration painless
- Field-level localization without content duplication

### Migration Guides

- Moving from WordPress to ModulaCMS
- Moving from Contentful to ModulaCMS
- Moving from Strapi to ModulaCMS

### Case Studies

- modulacms.com: building the product site with its own CMS
- Real-world usage stories (as users adopt the product)

## Post Structure

Every post follows this structure:

1. Problem or situation that motivated the work
2. What we tried, what worked, what didn't
3. The solution and how it works
4. What we learned or would do differently

No feature announcements disguised as blog posts. If a post doesn't have a real problem at its core, it's a changelog entry, not a blog post.
