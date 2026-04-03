# I built a CMS that doesn't fight your framework

**Author:** Michael Hegner
**Reading time:** ~10 min read

---

Every CMS I've used in the past 5 years has had problems that make working with them frustrating.

WordPress is a Jenga tower of code. Every request rebuilds the entire stack, and a bad line in a translation file crashes the site just as hard as a broken plugin. Umbraco is built with the same corporate Windows mentality as the ecosystem it runs on. Uselessly complex abstractions and error states that don't help you get unstuck, just block you from working. Contentful wants to host your data and charge you per user for the privilege. Payload wants you to buy the Pro tier before you can use features that should ship with every CMS.

I kept running into the same friction: the CMS and the frontend framework both trying to own routing, both trying to own rendering, both trying to own caching. Two systems designed to do the same job, stepping on each other's toes. Pick your framework, then spend weeks undoing what the CMS insists on doing for you.

Modula exists because I got tired of that fight.

## What it actually is

Modula is a headless CMS written in Go. It installs as a single binary, containing the REST API server, an admin panel, a TUI interface you can SSH into, and most of the same features as many modern CMS's. No Node.js. No npm. No runtime dependencies.

```
mkdir mysite && cd mysite
modula init
modula serve
```

That's zero to running. The init wizard creates your config, your database, seeds the bootstrap data, and registers the project in the global configs.json, letting you start any project's server from any directory. The serve command starts everything.

Modula was built for developers. I didn't compromise complexity for less performance, less flexibility, or less transparency. If you're unfamiliar with programming, Modula is not going to hold your hand. But it's simple enough to deploy anywhere, manage quickly, extend reasonably, and handle the build-locally-deploy-globally cycle that real projects demand. Most importantly, it's faster than the vast majority of open source CMS's.

Every feature was built on three pillars. Performance: nothing ships that makes the system slower without justification. Flexibility: nothing is hardcoded to a single environment, database, or deployment target. Transparency: everything is observable in the config file or the database. I didn't want people rooting around menu pages trying to find how to change the homepage.

## The content model

WordPress was built as a blog engine in 2003. If you wanted structured content, like a product catalog, an event calendar, a course curriculum, you installed Advanced Custom Fields and crammed everything into the posts table. ACF is a workaround bolted onto a blog engine's database, and it shows at scale.

Modula ships with opinionated defaults you can quickstart from and get to building immediately. But every datatype, every field, every route is just as customizable as anything you create yourself. Nothing is special. Nothing is locked.

Modula ships with a built-in admin panel that lets you manage everything in the CMS. But every way to manage content is just an API endpoint. The admin panel is a consumer of the API, not the API itself.

With the V1 release we're shipping our first admin-panel product: a Next.js admin panel built to work with Modula's admin tables. Import the admin datatypes, admin fields, and the default instances and routes for our default panel, and you have a working admin UI. But then you can manipulate the admin panel just like your client's content. The same content model, the same API, the same tools.

This is where it gets interesting for agencies. Build one admin panel and enable or disable features at runtime. Hide complexity for clients that don't need it. Tailor the content editing experience to each client without changing the underlying code. The admin panel is not sacred. It's just another application sitting on top of Modula's API.

## Plugin security, not plugin roulette

WordPress's plugin problem is well documented. 96% of WordPress security vulnerabilities originate in plugins. A WordPress plugin is a PHP script that hooks directly into the core with access to the database and filesystem. You install it and hope it was written by someone careful.

I didn't want that for Modula.

Plugins in Modula are Lua scripts that run in sandboxed virtual machines. Each plugin gets its own VM from a pool, with an operation budget and a circuit breaker. A plugin can register HTTP endpoints, hook into content lifecycle events, and use isolated database storage. It cannot access your main database tables. It cannot touch the filesystem. It cannot make arbitrary network requests.

Before a plugin's routes or hooks go live, they require explicit approval. You see exactly what the plugin is asking to do, approve it, and only then does it activate. If a plugin misbehaves, the circuit breaker kills it.

We chose Lua over JavaScript for the plugin runtime because Lua VMs are tiny, fast to spin up, and easy to sandbox. Node.js plugins would mean shipping a Node runtime, which defeats the purpose of a single binary with no dependencies.

## The TUI and project registry

I built the TUI because I wanted to manage content without opening a browser tab.

Run `modula tui <project> <environment>` and you get a full terminal interface. Content management, schema configuration, database operations, schema template installation. It's the only CMS that you can manage entirely from a terminal session. Every operation available in the TUI is also available as a CLI command, so anything you can do interactively, you can script.

The project registry is the other piece that changed how I work. When you run `modula init`, the project gets registered in a global `configs.json`. After that, you can start any project's server from any directory. You don't need to `cd` into the project folder. You don't need to remember port numbers or config paths. `modula serve myproject` from anywhere and it's running.

This sounds small. It isn't. When you're juggling four client projects and a personal site, being able to spin up any of them from wherever you happen to be in the terminal is the difference between flow and friction.

## Everything is observable

Most CMS platforms bury configuration in admin panel settings pages. Change the site title here. Toggle localization there. Manage API keys on this other page. You end up clicking through menus trying to find where some setting lives.

In Modula, there's one config file. `modula.config.json`. That's where the database connection lives, where S3 credentials live, where SSL settings live, where feature flags live. Open it, read it, change it. The CMS picks up changes live for most and on restart for the rest.

Everything else is in the database, and the database schema is straightforward enough that you can query it directly if you want to. No hidden state. No magic tables that the docs don't mention. If something is configured, you can find it, and you can find it without the admin panel.

## Built-in MCP

Modula ships with a built-in MCP (Model Context Protocol) server. Every instance exposes over 170 tools that AI agents can use to manage content, configure schemas, upload media, handle permissions, and do the tedious migration work that nobody enjoys doing manually.

The most tedious part of any CMS is bulk operations. Find and replace across 500 content entries. Migrate a custom field from one format to another. Rename, reorder, and restructure. This is exactly the kind of work that AI agents handle well, and Modula gives them the tools to do it without fragile screen-scraping or undocumented API abuse.

The MCP server isn't a bolt-on. It ships in the binary, runs on the same port, uses the same auth. Point Claude, Cursor, or any MCP-compatible agent at your Modula instance and it can do real work immediately.

## What it costs

Nothing. Modula is free, AGPL-3.0 licensed. All features ship in the binary. No tiers, no user caps, no upgrade prompts, no "contact sales for this feature."

The binary is the product. Download it, run it, own your infrastructure.

## Try it

```
modula init && modula serve
```

Two commands. Your CMS is running. Define a schema, create some content, point your frontend at the API.

If you're currently fighting your CMS, spending time working around it instead of with it, give Modula a look. The documentation is at [modulacms.com/docs](/docs). The source code is on GitHub. If you find bugs or have feedback, open an issue.
