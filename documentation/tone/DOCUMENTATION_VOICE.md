# Documentation Voice

Rules for writing ModulaCMS documentation and user guides. Derived from Next.js, Stripe, and Figma documentation patterns.

## Perspective

This is the most important rule. Everything else is secondary.

### Do

- Write from the user's perspective. Every page should answer "how do I use this?" not "how does this work internally?"
- Frame concepts around what the user can do, not how the system implements it.
  - "Organize your content into hierarchical trees. Nest pages under pages, reorder them freely."
  - NOT: "Content uses sibling pointers (`next_sibling_id`, `prev_sibling_id`) for O(1) navigation. The tree assembly algorithm queries all nodes, sorts by pointer, and reconstructs the hierarchy."
- Apply this test to every paragraph: "Does the reader need to know this to USE the feature, or does this explain how the feature is IMPLEMENTED?" If it's implementation, it belongs in internal docs (`ai/`), not here.

### Don't

- Don't expose database table names, column names, or schema details. The user interacts through the API and SDKs, not the database.
- Don't show Go struct definitions, internal type systems, or package paths. These are contributor concerns.
- Don't explain algorithm complexity, internal caching strategies, or middleware ordering. The user cares about behavior and capability, not mechanism.
- Don't describe how the system stores data. Describe what the user can store and retrieve.

## Page Structure

### Do

- Start every page with an H1 title followed by a single-sentence summary of what the page covers.
  - "Create a new ModulaCMS project and run it locally."
  - "Organize content into hierarchical trees with drag-and-drop reordering."
- Use progressive disclosure: quick start first, then details, then advanced topics.
- End pages with a "Next steps" or "API Reference" section that points the reader forward.
- Use tables for reference material — file conventions, configuration options, field types, CLI flags.
- Use numbered lists for sequential steps. Use bullet lists for unordered items.
- Follow strict heading hierarchy: H1 > H2 > H3. Never skip a level.
- Keep sections short. If a section runs longer than a screen, split it.

### Don't

- Don't open with "In this section, we will discuss..." or "This document explains...". State what the thing is, not what the document is about.
- Don't put multiple concepts in a single section. One section, one idea.
- Don't bury the action. If the reader needs to run a command or create a file, that instruction comes first — explanation follows.
- Don't create deep heading hierarchies. If you need H4, the page is too dense — split into separate pages.

## Voice and Tone

### Do

- Write in second person. Address the reader as "you".
- Use present tense and active voice. "ModulaCMS creates a database" not "A database will be created by ModulaCMS".
- Use imperative mood for instructions. "Create a config file" not "You should create a config file".
- Make confident, declarative statements. "Routes map URL patterns to handlers" not "Routes can be thought of as a way to map URL patterns to handlers".
- Use contractions naturally. "you'll", "it's", "don't", "doesn't". This is documentation, not a legal contract.
- Be direct. Say what something is in one sentence, then say what it does.
  - "A **datatype** is a schema definition for a piece of content. It defines which fields appear when an editor creates or updates content."

### Don't

- Don't hedge. Remove "might", "perhaps", "maybe", "it is possible that", "generally speaking". If something is conditional, state the condition explicitly.
- Don't use "please". Just tell the reader what to do.
- Don't use "simply", "just", or "easy". What is simple to you may not be simple to the reader. These words add nothing.
- Don't use exclamation marks in body text. Save enthusiasm for release announcements, not documentation.
- Don't use "we" to refer to the software. "ModulaCMS supports three databases" not "We support three databases". Reserve "we" for recommendations: "We recommend starting with SQLite for development."
- Don't apologize. "This feature requires Docker" not "Unfortunately, this feature requires Docker".

## Definitions and Concepts

### Do

- Define a concept before showing how to use it. Lead with what the thing IS, then what it DOES, then HOW.
  - **What:** "A **field** is a single piece of data within a datatype."
  - **Does:** "Fields define the type, validation rules, and display behavior for content values."
  - **How:** "To create a field, use the `fields` endpoint or the TUI."
- Bold the term being defined on first use.
- When comparing two concepts, use a clear "X vs Y" structure with concrete examples showing when to use each.
- When a concept exists on a spectrum (e.g., rigid vs. flexible, simple vs. powerful), frame it that way. Readers understand tradeoffs better than absolute recommendations.

### Don't

- Don't assume the reader knows your terminology. Define terms on first use in each page. A reader may land on any page directly.
- Don't define the same concept differently in different places. Pick one canonical definition and reuse it.
- Don't explain internal implementation details unless they affect the reader's decisions. The reader cares about what they can do, not how the code works.

## Sentences and Paragraphs

### Do

- Write short sentences. One idea per sentence.
- Keep paragraphs to 1-3 sentences. Dense paragraphs are walls that readers skip.
- Lead with the most important information. Put conditions and caveats after the main point.
  - "Run `just sqlc` after modifying schema files. This regenerates the type-safe Go code for all three database backends."
- Use parallel structure in lists. If the first item starts with a verb, every item starts with a verb.

### Don't

- Don't write filler transitions. Remove "As you can see", "As mentioned above", "It is worth noting that", "It should be noted that", "In order to".
- Don't repeat yourself. If you explained it once, link to it rather than restating it.
- Don't write sentences longer than 25 words when a shorter version works. Long sentences are harder to scan.
- Don't stack multiple clauses with commas. Break them into separate sentences.

## Code Examples

### Do

- Show code immediately after explaining a concept. The example should appear within visual range of the explanation.
- Label every code block with its file path or context (e.g., `Terminal`, `modula.config.json`, `internal/db/db.go`).
- Write minimal, focused examples. Show only what is being discussed — strip unrelated code.
- Use realistic values in examples. `my-blog-post` not `foo`, `BlogPost` not `Thing`.
- Include comments in code only when the logic is not self-evident from the example context.

### Don't

- Don't show a code example without first explaining what it does and why.
- Don't show long examples with irrelevant lines. If the reader needs to find the important part, the example is too long.
- Don't mix multiple concepts in a single code block. One example, one concept.
- Don't show code without explaining what the reader should do with it — is this a file they create? A command they run? Output they should expect?

## Callouts and Asides

### Do

- Use **"Good to know"** blockquotes for supplementary information that doesn't interrupt the main flow — edge cases, compatibility notes, default behaviors.
- Use **"Tip"** for actionable suggestions that improve the reader's workflow but aren't required.
- Bold the callout label. Use blockquote formatting.
  - `> **Good to know**: SQLite is the default database. No additional setup is required for development.`
- Place callouts after the relevant section, not before.

### Don't

- Don't use callouts for critical information. If the reader must know it to succeed, put it in the main text.
- Don't stack multiple callouts in sequence. If you have that much supplementary info, restructure the section.
- Don't use "Note:", "Warning:", "Important:", "Caution:" — pick one consistent pattern and stick with it. We use "Good to know" and "Tip".

## Headings

### Do

- Use descriptive, scannable headings. A reader skimming headings should understand the page structure.
- Use task-oriented headings for guides: "Create a datatype", "Configure routes", "Set up media storage".
- Use noun-phrase headings for reference: "System requirements", "File conventions", "Configuration options".
- Keep headings short. Under 8 words is ideal.

### Don't

- Don't use question-format headings ("How do I create a route?") in guides. Save question headings for FAQ pages.
- Don't use vague headings like "Overview", "Details", "More information", "Miscellaneous".
- Don't number headings manually. Let the document structure speak for itself.

## Links and Cross-references

### Do

- Link to related pages at natural points in the text where a reader would want to go deeper.
- Use descriptive link text. "Learn more about content trees" not "Click here" or "Learn more".
- End pages with a curated list of related pages under "Next steps" or "API Reference".

### Don't

- Don't link the same page more than once in a section.
- Don't use "above" or "below" to reference content. Readers may not be reading linearly. Link directly.
- Don't create circular references where two pages just point at each other with no additional information.

## Tables

### Do

- Use tables for structured reference data: file lists, CLI flags, configuration keys, comparison matrices.
- Keep tables to 2-4 columns. More columns are hard to scan.
- Use consistent column headers. `| Name | Purpose |` or `| Option | Default | Description |`.
- Alphabetize or group table rows logically, not randomly.

### Don't

- Don't use tables for sequential instructions. Use numbered lists instead.
- Don't put long prose in table cells. If a cell needs more than one sentence, the content belongs in a section, not a table.

## Content Types

Different page types have different goals. Match the structure to the type.

**Getting Started / Quickstart:** Minimal steps to a working result. No theory. Link to concepts for readers who want depth.

**Concept / Explanation:** Define what a thing is and why it exists. Use examples and comparisons. No step-by-step instructions.

**Guide / Tutorial:** Walk through a task from start to finish. Show every step. Assume the reader is following along.

**Reference:** Exhaustive, scannable, structured. Tables, parameter lists, return types. No narrative. The reader is looking up a specific detail.
