# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Marketing and documentation website for [ModulaCMS](../modulacms/) -- a free, open-source, single-binary headless CMS written in Go (AGPL-3.0). This site is built with Astro and will be powered by its own CMS backend via REST API.

The CMS backend lives at `../modulacms/` (sibling directory). API calls use the `?format=clean` query parameter for nested content trees.

## Commands

```bash
npm run dev       # Dev server at localhost:4321
npm run build     # Production build to ./dist/
npm run preview   # Preview production build locally
```

## Architecture

**Framework:** Astro 5.x (static site generator), TypeScript strict mode.

**Content source:** ModulaCMS REST API. Each route returns a single content tree where parent nodes contain children via `nodes[]` arrays. No relation IDs between datatypes -- parent/child is expressed by nesting.

**API endpoints** (all use `GET /api/v1/routes/{route}?format=clean`):

| Route | Endpoint path | Root type |
|-------|--------------|-----------|
| `/` | `landing` | Page > Section > Feature, Comparison, CodeExample |
| `/templates` | `templates` | Page > StarterTemplate |
| `/docs/:slug` | `docs/:slug` | Doc > CodeExample, APIEndpoint, CLICommand |
| `/changelog` | `changelog` | Page > Changelog |

**Response shape** (clean format):
```json
{
  "id": 1,
  "type": "Page",
  "title": "...",
  "slug": "/",
  "_meta": { "authorId": 1, "routeId": 1, "dateCreated": "...", "dateModified": "..." },
  "nodes": [
    { "id": 2, "type": "Section", "section_type": "hero", "nodes": [...] }
  ]
}
```

Fields are flattened as top-level keys (not nested in a `fields` object). System metadata lives in `_meta`. Children live in `nodes[]`.

## Key Reference Files

- `PRODUCT_BRIEF.md` -- Full product strategy, messaging, competitive positioning, and feature inventory for ModulaCMS. Use this as the source of truth for marketing copy and site content.
- `SCHEMA.md` -- Content model definition with all datatypes, fields, and parent/child relationships.
- `SCHEMA_EXAMPLES.json` -- Complete example API responses for every route in clean format.
- `STYLE_GUIDE.md` -- Visual design tokens, color palette, typography, spacing scale, component conventions. Source of truth for all CSS values.

## Content Model

Landing page sections are differentiated by `section_type` field: `hero`, `install`, `tui_demo`, `features`, `comparisons`, `cta`. Each section type renders differently and may have different child node types.

Datatypes: Page, Section, Feature, Comparison, CodeExample, Doc, APIEndpoint, CLICommand, StarterTemplate, Changelog. See `SCHEMA.md` for field definitions.

## Styles

Vanilla CSS with Astro scoped styles. Dark theme only. All design tokens as CSS custom properties.

**Global CSS files** (import order matters):

| File | Purpose |
|------|---------|
| `src/styles/tokens.css` | All custom properties on `:root` (colors, fonts, spacing, layout, motion) |
| `src/styles/base.css` | Reset, element defaults (body, headings, links, code, pre) |
| `src/styles/utilities.css` | `.sr-only`, `.container`, `.container-narrow`, `.container-wide` |

These are not yet imported -- requires a `Base.astro` layout to wire them up.

## Current State

Astro starter template with global CSS infrastructure. Single `src/pages/index.astro` with placeholder content. Schema, content strategy, and visual design tokens are defined. No layout, components, or pages built yet.
