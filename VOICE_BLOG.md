# VOICE_BLOG.md — Case Studies, Blog Posts, Product Updates

## Identity

Blog and case study content sounds like a developer writing about something they built, used, or learned. First person, real problems, real outcomes. The reader should feel like they're hearing from someone who actually did the work, not someone summarizing it from a press release.

## Core Voice

**Narrative.** Blog posts tell stories. Start with a problem or situation, walk through what happened, end with what came out of it. Don't structure posts like feature lists.

**Honest.** Include what didn't work. Include what was harder than expected. Include what you'd do differently. Developers trust writers who show the full picture, not just the highlight reel.

**Personal.** Use "I" and "we" freely. Share specific decisions and why they were made. "We chose SQLite for local dev because it eliminates the Docker-compose-just-to-write-content problem" is better than "ModulaCMS supports multiple databases."

**Opinionated.** Pick a side. Don't say "both approaches have their merits." If you think WordPress's plugin model is a security risk, say so and explain why.

## Anti-AI Rules

### Max one bulleted list per article
Bullet lists are the biggest tell for AI-generated content. Convert everything to prose. If you absolutely need one list for a collection of items, use it once and make it count.

### No numbered step lists
If you're walking through a process, write it as prose with transitions. "First, install the binary. Then run the init wizard — it creates your config and database in one pass." Not a numbered list.

One exception: if the article is explicitly a tutorial and steps need to be referenced by number.

### No em dashes for pauses
Use commas, periods, or ellipsis. Em dashes in casual technical writing are an AI tell.

### No groups of three
Don't list three adjectives, three benefits, three examples in parallel. Two or four. Break the pattern.

### No summary sections
Don't end with "In conclusion" or "To summarize." When you're done explaining, stop. End with the last useful thing, not a recap.

### Vary paragraph length
Mix it up. Single sentence for emphasis. Longer paragraphs when building an argument. Two-sentence transitions. Uniform paragraph blocks read as generated.

### No emoji, no decorative symbols
Technical writing doesn't need decoration.

### No duplicated talking points
If you've made a point, don't restate it in different words later. Say it once, move on. AI loves to rephrase the same idea across multiple sections.

## Tone Rules

### Problem first, always
Every post starts with a concrete problem. Not "Today we're excited to announce..." but "We kept running into the same issue when..."

A reader should understand the pain point within the first two paragraphs. If they've felt that pain, they'll keep reading.

### Show the messy parts
Don't just present the polished outcome. Show the first attempt that didn't work. Show the constraint that forced a design change. Show the trade-off you wrestled with.

```
My first thought was to add caching at the CMS layer. Built it,
benchmarked it, found out it conflicted with how Next.js handles
ISR. Ripped it out. The CMS shouldn't cache because the framework
already does. Lesson learned.
```

### Admit uncertainty
Say "I don't know" or "I haven't tested this at that scale" instead of hedging with formal qualifiers. "While this approach may potentially work in certain scenarios" is the sound of an AI covering its bases. "This worked for us. Might not work at 10x the traffic — haven't tried it" is a human being honest.

### Have actual opinions
Don't hedge. If WordPress's ACF approach is a bad architecture, explain why.

**Not this:** "Both WordPress and ModulaCMS are viable options, each with their own strengths and weaknesses."

**This:** "WordPress was never designed for structured content. ACF is a workaround bolted onto a blog engine's database, and it shows at scale."

### Skip perfect transitions
You don't need smooth bridges between every section. Sometimes just jump to the next point. "Now that we understand the problem, let's explore potential solutions" is AI filler. Just start the next section.

### Break grammar rules sparingly
Start sentences with "And" or "But" when it feels natural. Use fragments for emphasis. But don't overdo it — too much and it sounds forced.

### Use asides and tangents
Brief parenthetical thoughts add personality. "(Which, honestly, we should have caught sooner.)" These feel human because they are human.

## Content Type Guidelines

### Case Studies
Structure around the customer's problem, not ModulaCMS's features. What were they using before? What broke? What changed after migration? Include specific numbers if available — page load times, API response times, developer hours saved.

Don't write case studies that read like feature tours wearing a customer's name.

### Product Updates / Changelog Posts
What changed and why. Lead with the motivation, not the feature name. "Field-level localization meant duplicating entire content entries per language. Now locale variants live on the field itself" tells a better story than "We added field-level localization."

### Technical Blog Posts
Follow the narrative structure: problem, what we tried, what worked, what we learned. Code examples are appropriate here (unlike marketing pages) but should illustrate a point, not demonstrate an API.

## Checklist

Before publishing blog content:

- [ ] Does it start with a problem or situation, not a feature announcement?
- [ ] More than one bulleted list? Convert extras to prose.
- [ ] Any groups of three? Break the pattern.
- [ ] Any em dashes used as pauses? Replace with commas or periods.
- [ ] Are all paragraphs the same length? Vary them.
- [ ] Is the same point made twice in different words? Cut one.
- [ ] Any hedging where an opinion should be? Pick a side.
- [ ] Does it end with a summary section? Delete it — end with the last useful thing.
- [ ] Read it out loud. Does it sound like a person talking about their work?
