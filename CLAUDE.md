# gibson-tool-runner — CLAUDE.md

> **Workflow rules:** see [`zeroroot-ai/.github` → `AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md) — canonical for branching / commits / PRs / releases / merging. Conventional Commits MANDATORY. Never push to main. Never force-push.

This file is the per-repo addendum. Workspace-wide concerns live in [`~/Code/zeroroot.ai/CLAUDE.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md); architectural decisions in [`docs/adr/`](https://github.com/zeroroot-ai/docs/tree/main/adr).

## TL;DR

One microVM image, one Go binary, N parsers for CLI security and ops tools. The Gibson daemon dispatches sandboxed tool calls into a Setec microVM running this image. Entry point: `make test` then `make image`.

## Architecture

The binary reads a typed proto request from `GIBSON_TOOL_INPUT_B64` (base64 of a protojson-serialised request), dispatches to the matching `registry.Parser`, shells out to the installed CLI, and emits the response on stdout as `===GIBSON_TOOL_OUTPUT===<base64(protojson(response))>` (exit 0) or `===GIBSON_TOOL_ERROR===<message>` (exit 2). Parsers live under `parsers/<tool>/`; each registers via an `init()` call imported by `cmd/gibson-runner/main.go`.

Each parser translates raw CLI output into `gibson.graphrag.DiscoveryResult` nodes — the taxonomy-aligned graph objects the daemon writes to Neo4j. This is the Apache-licensed open execution layer (ADR-0054): it depends only on the public OSS `sdk` plus community libraries (connectrpc/otel) and must **not** import `platform-clients`, `gibson`, or any ELv2/closed module. The transport/observability/readiness primitives it once consumed from `platform-clients` now live in-repo under `internal/` (severed per #98).

Full ABI contract: `TOOLS.md`. Security model: `SECURITY.md`.

## Regen commands

```bash
make bin            # build the gibson-runner binary to bin/
make image          # build the Docker image (debian:trixie-slim base)
make list-tools     # run bin/gibson-runner --list-tools
```

## Gotchas

- **No `platform-clients` / `gibson` / closed-module imports.** This is the Apache execution layer; it builds against the public `sdk` + community libs only. Transport, observability, and readiness helpers live under `internal/`. A new ELv2/closed dependency would break the license boundary — keep it out.
- **Tool ABI sentinel line.** Output parsing depends on the exact prefix `===GIBSON_TOOL_OUTPUT===`; do not alter it without updating the daemon catalog.
- **Adding a parser**: create `parsers/<tool>/<tool>.go`, implement `registry.Parser`, register in `init()`, add `testdata/` golden files, add a blank import in `cmd/gibson-runner/main.go`, add the apt/go/pip install step to `Dockerfile`.

## Links

- Org-level workflow: [`AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md)
- Workspace map: workspace `CLAUDE.md`
- Per-repo ADRs: [`docs/repos/gibson-tool-runner/adr/`](https://github.com/zeroroot-ai/docs/tree/main/repos/gibson-tool-runner/adr)
- Domain glossary: [`docs/glossary.md`](https://github.com/zeroroot-ai/docs/blob/main/glossary.md)
- PR checklist: [`docs/agents/pr-checklist.md`](https://github.com/zeroroot-ai/docs/blob/main/agents/pr-checklist.md)
