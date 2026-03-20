# Why ModulaCMS Exists

ModulaCMS was born from years of agency work fighting traditional CMSes. You know the drill. Monday morning standup. The orthodontist needs a before-and-after slider with a booking widget. The dog treat company wants a "build your own box" configurator with real-time pricing. The divorce attorney, the one that took four months to approve a shade of blue, now needs a live chat integration by Thursday because the partner's nephew said AI is the future.

And you're sitting there, four browser tabs into a WordPress plugin marketplace that smells like 2014, trying to jury-rig a page builder into doing things that would make its original developers file a restraining order. You've got a `functions.php` file three thousand lines deep, `add_filter` calls stacked like Jenga in an earthquake. Every time you look at the plugins update page is like looking into the abyss, not knowing if any one update is going to consume your next 3 hours and require a phone call to the client about why they got an email from wp-engine that a rollback occured.

Every feature in ModulaCMS exists because over the last five years as an agency dev I have stared at my screen and said "there has to be a better way." If you've never said that, this product probably isn't for you. If you've said it so many times it's become a mantra, keep reading.

## The WordPress Problem

WordPress started as a blogging platform. Then it became a page builder, an e-commerce platform, a learning management system, a membership site, a real estate listing engine, a patient intake portal. Somewhere along the way it forgot how to do the one thing a headless CMS actually needs to do: serve JSON fast.

The content model is the core issue. Under the hood it's one god-table, `wp_posts`, and every content type is metadata duct-taped onto the side of it. Advanced Custom Fields is the standard workaround, and it makes the problem worse. Every ACF field becomes a row in `wp_postmeta`, a key-value table designed for post metadata, not complex content. A content type with 30 fields means 30+ joins per query running through WP_Query, which was built for blog posts. At scale the admin panel lags, queries take seconds, and the browser tab begs for mercy.

Then there's the plugin ecosystem. WordPress plugins run with full PHP and filesystem access. No sandboxing. No approval workflow. No resource limits. The premium SEO plugin phones home with your entire sitemap. The contact form widget has full write access to your users table for reasons known only to God and a developer in 2011. A plugin updates, your site goes white-screen, and you're in an emergency SSH session deactivating plugins by renaming folders like a bomb disposal technician cutting wires.

And the settings. WordPress scatters configuration across seventeen different admin screens. Your site title is under Settings > General. Your permalink structure is Settings > Permalinks. Your homepage display is Settings > Reading, sandwiched between "posts per page" and "search engine visibility," because apparently someone decided that what the homepage shows is a reading comprehension issue. SMTP settings? That's a plugin. CORS headers? Different plugin. Cron jobs? That's `wp-config.php`, a PHP file sitting in your web root with your database password in it. 

ModulaCMS stores content fields in dedicated tables with proper relationships. The data model was designed for complex content from the start, not retrofitted onto a blog engine. Everything configures through one `config.json` file. One place. One format. Every text editor on earth can read it.

## The SaaS Pricing Trap

Contentful and Sanity are good products with pricing models that create problems as your team grows.

Contentful caps at 5 users and 25,000 records on its free tier. Sanity caps at 3 users and 500,000 API calls per month. For a small team starting out, it's manageable. For a growing company, costs compound fast and unpredictably. And then there's the lock-in. Your content lives on their infrastructure. Migrating away means extracting everything through their API, hoping the export is complete, and rebuilding your content model in whatever you move to. If you want to discuss your export options, please contact their enterprise sales team.

ModulaCMS is self-hosted. All features, no user limits, no API call caps. Free under AGPL-3.0. Your data lives on your infrastructure. The database is standard SQL. The media is in your S3 bucket. The backups are ZIP files. If you decide tomorrow that ModulaCMS isn't for you, you take everything with you. No exit interview. No data hostage negotiation.

## The Node.js Dependency Chain

Strapi and Payload are capable headless CMSes that carry the Node.js ecosystem with them. That means npm, package.json, node_modules, and everything that comes with it. Dependency resolution, version conflicts, security advisories on transitive dependencies you didn't choose. A fresh Strapi install pulls hundreds of packages. The Node.js runtime alone, before you install a single package, is 80 MB. Your JavaScript runtime is three times the size of the entire ModulaCMS binary and it hasn't done anything yet.

Payload recently added paid tiers that gate features behind enterprise pricing. And since Figma acquired them, the product direction is tied to acquisition-driven priorities. That may or may not matter to you, but it's worth knowing.

ModulaCMS ships as a single Go binary. 27-29 MB. No runtime. No package manager. No dependency tree. No "install nvm to install node to install npm to install yarn to install the thing that installs the other things." One file. Copy it to the server. It runs.

## Why We Chose Go

Go isn't a trendy choice. It's a boring choice, and boring is exactly what you want from the language running your production infrastructure.

Compiled to native machine code. Goroutines for concurrency without callback spaghetti. Garbage collection that doesn't stop the world. The result is a CMS that starts faster than most CMSes finish loading their configuration, serves thousands of concurrent requests without flinching, and runs three servers in less memory than a typical Node process uses to import its dependencies.

Here's the part that sounds like marketing but isn't: a $5 shared-CPU Linux box running ModulaCMS will outperform a Vercel deployment. Not compete with. Outperform. A single-core VPS with 1 GB of RAM will serve content faster than a globally distributed edge network backed by a company valued at billions. Because ModulaCMS is a compiled binary doing one thing well on hardware it has to itself, and Vercel is spinning up cold-start functions, routing through edge middleware, and adding latency at every layer of abstraction between your content and your user. The $5 box doesn't have a CDN. It doesn't need one. It responds before the CDN would have finished its TLS handshake.

Go is the language that runs behind every microservice powering every Node server on the planet. Kubernetes, Docker, the API gateways, the load balancers, the service meshes. The companies that run the internet at scale tried other languages first and rewrote everything in Go when the traffic got real. ModulaCMS puts that same language directly behind your content API.

## Why No Server-Side Rendering

ModulaCMS does not render HTML. That's intentional, and it's the design decision people ask about most.

Modern frameworks handle rendering. Next.js, Nuxt, SvelteKit, Astro, Go templates, Laravel Blade. They all have well-tested opinions about how to turn data into pages.

When a CMS also does rendering, you get conflicts. Your CMS and your framework both want to own URL structure. Two session systems that need bridging. Two caching strategies that fight each other. Two image optimization pipelines doing the same work. ModulaCMS eliminates the overlap by refusing to participate. The CMS provides data. The framework handles everything else.

ModulaCMS is a data authority. It stores, organizes, versions, and serves structured content. The CMS answers "what is at this URL?" Your framework decides "what do I do about it?" Your framework already solves these problems. Your CMS shouldn't solve them again, differently.

## Why Sandboxed Lua Plugins

ModulaCMS took one look at the WordPress plugin ecosystem and had a religious experience. Not the good kind.

The mass of `wp_options` rows injected by plugins uninstalled years ago. jQuery versions fighting each other in the `<head>` tag. The premium SEO plugin quietly phoning home with your sitemap. The entire model is "here's the keys to everything, we trust you, please don't burn the house down." And then the house burns down. Every single time.

ModulaCMS plugins run in sandboxed Lua VMs with the opposite philosophy. A plugin declares what it is. It declares what it wants, which tables it needs, which content hooks it wants to intercept, which HTTP routes it wants to register. And then it asks for approval. You review. You approve or deny. The plugin gets access to its own isolated tables and nothing else. It cannot touch core CMS tables. It cannot touch other plugins' tables. It cannot silently inject ads into your footer.

If it misbehaves anyway, infinite loop or runaway memory, a circuit breaker trips automatically and shuts it down. No white screen. No emergency SSH session. Just a notification that the plugin broke itself and a button to re-enable it when the author fixes their code.

Lua was chosen because it embeds cleanly, sandboxes reliably, and is fast enough for content lifecycle hooks without introducing a full runtime. It's not as flexible as Node, and that's the point.

## Why You Should Know Where Your Stuff Is

Here's a real story. A team is running a headless CMS on an Azure compute instance. Healthcare. The uploads need to be HIPAA-compliant, organized in a specific bucket structure, access-logged. Someone built the upload pipeline. The files are going up. Thumbnails render in the admin panel. The client is uploading patient-facing documents.

Except nobody verified that the upload pipeline was connected to the storage bucket. The CMS was quietly writing files to its own local filesystem, the local filesystem of the Azure compute unit, a compute unit you can't SSH into and can't browse the filesystem of. HIPAA-sensitive media disappearing into a sealed box in a Microsoft data center.

They found out when the compute instance recycled and the files vanished. The backups were also on the local filesystem. Of the same compute instance. That just recycled.

ModulaCMS doesn't let this happen because it doesn't make storage decisions for you. You configure where media goes. You configure where backups go. You configure the database connection string. Every file has a location you specified. Every decision is logged in an audit trail you can query. There are no black boxes.

If you decide tomorrow that ModulaCMS isn't for you, you take everything with you. The database is standard SQL. The media is in your bucket. The backups are ZIP files you can open with any computer. No "please contact our enterprise sales team to discuss your export options."

## Why Opt-In Everything

ModulaCMS ships with all features available, but most are off by default.

This isn't minimalism for its own sake. It's predictability. When you enable localization, you know it's there because you chose it. When you don't need webhooks, they don't consume resources or add complexity to your mental model. Most CMSes want to be your entire world. They want to manage your content, your users, your auth, your media, your emails, your analytics, your deployment. They get increasingly passive-aggressive when you try to use only part of them.

ModulaCMS has no such ego. If all you ever use is the media pipeline, S3 storage with automatic WebP conversion and responsive dimension presets, then that's your entire integration surface. Hit the media endpoints. Get optimized images. Ignore everything else. The content tree doesn't mind. The RBAC system isn't offended. Use one endpoint or use all 47. ModulaCMS is a toolbox, not a religion.

## How It Scales

Every CMS makes you choose upfront. WordPress is for small sites until it isn't, and then you migrate. Contentful is enterprise from day one and charges accordingly, even when you're three people with twelve pages. Every CMS is built for a specific size, and the moment you outgrow it, you're looking at another migration. The third one this decade.

ModulaCMS doesn't make you choose because it works at every scale.

Day one, you're a small agency. One binary on one server. SQLite for the database because you don't feel like configuring Postgres on a Tuesday. Total infrastructure cost: a $5 VPS.

Year two, the agency is growing. Fifteen clients, real traffic, SQLite isn't cutting it. You change one string in `config.json` from `"sqlite"` to `"postgres"`, point it at a managed database, and you're done. Same binary, same admin panel, same everything.

Year five, you're running a distributed network. Multiple instances behind a load balancer, spread across regions. The audit trail handles distributed ordering. The plugin coordinator syncs state across instances via the database. The deploy sync system pushes content between environments with schema validation. None of this required an enterprise edition upgrade. It was all there from the first commit, waiting for you to need it, not charging you while you didn't.

## What ModulaCMS Does Not Do

No server-side rendering. No caching layer. No native forms. No cron jobs. No frontend opinions.

These are boundaries, not gaps. ModulaCMS does not do everything WordPress does because doing everything WordPress does is the problem. WordPress spent two decades bolting on every feature anyone ever asked for, and the result is an application that can technically do everything and does none of it particularly well.

ModulaCMS does one thing. It serves structured content over HTTP at speed. Content in, data out. Everything else is yours.

## The Comparison

| | ModulaCMS | WordPress | Contentful | Strapi | Payload |
|---|---|---|---|---|---|
| License | AGPL-3.0 | GPL-2.0 | Proprietary | MIT / EE | MIT / EE |
| Runtime | None (Go binary) | PHP | Cloud | Node.js | Node.js |
| Binary size | 27-29 MB | N/A | N/A | ~200 MB+ installed | ~150 MB+ installed |
| User limits | None | None | 5 (free) | None | None |
| API call limits | None | None | Tiered | None | None |
| Paid tiers | No | No | Yes | Yes | Yes |
| SSH TUI | Yes | No | No | No | No |
| Plugin sandboxing | Yes (Lua VMs) | No | N/A | No | No |
| Self-hosted | Yes | Yes | No | Yes | Yes |
| Multi-format API | Yes | No | No | No | No |
