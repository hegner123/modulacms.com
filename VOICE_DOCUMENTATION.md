# VOICE_DOCUMENTATION.md — Guides, References, Tutorials, Changelogs

## Identity

Documentation sounds like a knowledgeable colleague answering your question directly. No selling, no narrative, no personality. The reader is here to accomplish something specific. Respect their time by getting to the answer fast.

## Core Voice

**Direct.** Lead with the answer or the command. Explain after. The reader came here with a question — answer it before providing context.

**Precise.** Use exact terminology. Don't simplify names to be "friendlier." If the flag is `--admin-password`, write `--admin-password`, not "the admin password flag." If the endpoint returns a 404, say 404, not "an error."

**Assume competence.** The reader is a professional developer. Don't explain what a REST API is. Don't explain what JSON is. Don't explain what a binary is. If a concept is specific to ModulaCMS, explain it once and link back to that explanation.

**Complete.** Show full, runnable examples. Don't show a function call without showing the import. Don't show a config snippet without showing where the file lives. Incomplete examples waste more developer time than verbose ones.

## Anti-AI Rules (Relaxed)

Documentation has different needs than marketing or blog content. Some AI-pattern rules are relaxed here because lists, steps, and structured formatting are natural in technical docs.

### Bulleted lists are acceptable
Lists are a natural documentation format. Use them freely for enumerating options, flags, parameters, and features. But don't use bullets when a table or prose would be clearer.

### Numbered lists are acceptable for sequential steps
If steps must happen in order, number them. But prefer prose for short sequences (under four steps).

### No em dashes for pauses
Still applies. Use commas or restructure.

### No groups of three (relaxed)
Less critical in docs than marketing, but still watch for artificial-feeling triads in introductory prose. Lists of three parameters or three config options are fine — that's just what exists.

### No summary sections
Still applies. Don't recap at the end of a doc page. If the reader needs a summary, the page is too long — split it.

### No emoji, no decorative symbols
Still applies. Use formatting (bold, code, headings) for hierarchy instead.

### Vary paragraph length (relaxed)
Docs naturally have shorter paragraphs. Uniform short paragraphs are fine here — that's just how reference material reads.

## Tone Rules

### Code first, explain after
Show the command or code example, then explain what it does and why. Don't make the reader wade through three paragraphs of context before seeing the thing they came for.

**Not this:**
```
To initialize a new ModulaCMS project, you'll need to run the
initialization command. This command creates a configuration file,
sets up the database, and seeds initial data. The command is:

modula init
```

**This:**
```
modula init

Creates `modula.config.json` and `modula.db`, builds all tables,
seeds bootstrap data, and registers the project.
```

### Show complete, runnable examples
Every code example should work if copied and pasted into a terminal or file. Include the full command with flags. Show the expected output when it's not obvious.

### Don't over-explain
If something follows standard conventions (REST endpoints, JSON responses, SQL queries), don't explain the convention. Explain what's specific to ModulaCMS.

### Use second person
"Run `modula init` to create your project" not "The user runs `modula init` to create their project."

### Flag edge cases
If something behaves differently on SQLite vs PostgreSQL, or locally vs production, call it out explicitly. Don't bury compatibility notes in prose — use a callout or bold the distinction.

### No selling
Documentation is not marketing. Don't say "the powerful plugin system" — say "the plugin system." Don't say "effortlessly manage content" — say "manage content." Features don't need adjectives in docs.

## Content Type Guidelines

### Getting Started / Quickstart
Absolute minimum path from zero to running. Every step should be a command the reader can execute. Cut any step that isn't strictly necessary to get to "it works." Save configuration options and deeper explanation for dedicated pages.

### Concept / Architecture Guides
Explain the mental model first, then the implementation details. "Content is a tree. Pages contain rows, rows contain columns, columns contain blocks" gives the reader a framework before they see the schema.

Keep these focused on one concept per page. Don't explain the content model and the plugin system on the same page.

### API Reference
Endpoint, method, parameters, response shape, status codes. Use tables for parameters. Show a curl example and the response body. Note required vs optional parameters.

### CLI Reference
Command, flags, description, example. Show the output. Note which flags have defaults and what those defaults are.

### Configuration Reference
Show the full config file with comments explaining each field. Then break down non-obvious fields individually. Don't explain fields whose names are self-evident (`port`, `database_url`).

### Changelog
One line per change when possible. Group by category (Added, Changed, Fixed, Removed). No excitement, no editorializing. "Added field-level localization with BCP 47 codes" not "Exciting new localization support."

For breaking changes, bold the entry and include a migration path.

## Checklist

Before publishing documentation:

- [ ] Does every code example work if copy-pasted?
- [ ] Is the answer/command shown before the explanation?
- [ ] Any marketing language? ("powerful", "seamless", "effortless") Remove it.
- [ ] Any concepts explained that professional developers already know?
- [ ] Are edge cases and platform differences called out?
- [ ] Any em dashes used as pauses? Replace.
- [ ] Does the page cover one topic, or should it be split?
- [ ] If there's a getting-started guide, can someone go from zero to running in under five minutes following it?
