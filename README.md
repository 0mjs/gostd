# gostd

An opinionated, searchable tour of the Go standard library: curated guides for the packages you will actually use, generated reference for the rest.

Try it out, [here](https://gostd.up.railway.app/)

## What it is

`gostd` is a small Go web app that turns the standard library into a browsable learning resource.

It combines two kinds of content:

- Curated, hand-written package guides for high-traffic packages such as `fmt`, `strings`, `os`, `net/http`, and `encoding/json`
- Generated overviews for the rest of the public standard library, built at startup from the local Go installation in `GOROOT`

The result is a site that gives you:

- A fast way to understand what a package is for
- Task-oriented examples instead of raw API dumps
- A "Start here" curriculum for learning the stdlib in a sensible order
- A recipe index for common "How do I...?" questions
- Search across packages, sections, examples, and recipes
- Clear warnings for deprecated, frozen, insecure, or otherwise discouraged packages

## Why this exists

The Go standard library is excellent, but the official docs are reference-first. That works once you already know what you are looking for. It is less helpful when you are still trying to answer:

- Which package should I reach for?
- What is the smallest useful example?
- What is the normal, idiomatic path here?
- Is this package still recommended?

`gostd` is built to answer those questions quickly.

## Running locally

This project has no third-party runtime dependencies.

Requirements:

- Go `1.26` or newer

Start the server:

```bash
go run .
```

By default it listens on:

```text
http://localhost:8080
```

You can override the port with `PORT`:

```bash
PORT=3000 go run .
```

## How it works

At startup, the app:

1. Registers all curated content from the `content/` package
2. Scans the local Go standard library under `GOROOT/src`
3. Generates package summaries, overview sections, and reference sections for public stdlib packages that do not have curated entries yet
4. Merges generated reference sections into curated packages so hand-written guides can still include broader API coverage
5. Builds a search index served at `/search.json`

The generated content is intentionally reference-oriented. The curated content is where the teaching voice, recipes, cheatsheets, and opinionated guidance live.

## Routes

- `/` — homepage, curriculum, and category browse view
- `/pkg/{name}` — package page
- `/recipes` — task-oriented recipe index
- `/search.json` — search payload for the command palette

## Project structure

```text
.
├── main.go           # HTTP server, routes, page assembly
├── content/          # curated package guides, recipes, search, generation logic
├── templates/        # HTML templates
├── static/           # CSS and static assets
└── go.mod
```

Important files:

- `main.go` wires the server, embedded templates/static assets, and page rendering
- `content/content.go` defines the core content model and registry
- `content/generated.go` scans `GOROOT` and generates fallback package coverage
- `content/curriculum.go` defines the "Start here" learning path
- `content/recipes.go` contains the task-oriented recipe library
- `content/search.go` builds the compact search index used by the UI

## Content model

Each package page is built from a `content.Package` with:

- package metadata
- a summary
- one or more sections
- examples per section
- optional advisories
- optional cheatsheet rows and TL;DR content

That keeps authored content simple: each `content/*.go` file can register one package and describe it in terms of real tasks and examples.

## Writing curated content

Curated package guides live in `content/*.go` and register themselves during `init()`.

In practice, a strong curated page should include:

- A short summary of when to use the package
- A few sections grouped by real tasks
- Small, idiomatic code samples
- Notes about common mistakes or tradeoffs
- Advisory flags when a package is deprecated, frozen, insecure, or niche

Recipes are separate from package pages and are grouped by task domain on `/recipes`.

## Design approach

The site is intentionally:

- Fast
- Mostly static at runtime
- Searchable without a server round-trip per keystroke
- Useful for both beginners and working Go developers

It is not trying to replace [pkg.go.dev](https://pkg.go.dev/). It sits in front of it: teach first, reference second.

## Limitations

- Generated coverage depends on the local Go toolchain present in `GOROOT`
- Generated package documentation is useful, but it is not a substitute for hand-written examples
- Search is client-side and backed by a prebuilt JSON index
- The project currently uses embedded templates and static assets rather than a separate frontend build

## Future directions

Likely improvements:

- richer recipe coverage
- better ranking in search
- package cross-links based on common workflows
- tests around generated content shape and rendering
- export or snapshot support for offline reading
