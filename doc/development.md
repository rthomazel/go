# Development Guide

## Prerequisites

- **Go 1.26.1+** (managed via [mise](https://mise.jdx.dev/))
- **gofumpt** — formatter
- **golangci-lint** — linter
- **prettier** — YAML/JSON formatter
- **Node.js** — required for prettier and commitlint
- **bash** 4+

Run setup to verify all tools are present:

```bash
./run setup
```

This checks for required binaries and prints install commands for anything missing.

## Common tasks

```bash
# Build a module
./run build logging

# Run tests for a module
./run test misc

# Lint a module
./run lint cmd

# Format a module
./run format httpmisc

# Format all YAML/JSON configs
./run format-configs

# Fix copyright headers
./run copyright-fix-go
./run copyright-fix-sh

# Generate mock files
./run generate-mocks

# Regenerate go.work (after adding/removing modules)
./run generate-go-work
```

## Adding a new library module

```bash
./run new-module
```

The script (`sh/new_module.sh`) scaffolds the module directory, `go.mod`, and initial source files, then regenerates `go.work`.

## Adding a new CLI command

```bash
./run new-command
```

The script (`sh/new_command.sh`) copies the `cmd/template` directory, renames it, and wires it into the workspace.

## Module versioning & releases

Each sub-module is versioned independently using Git tags in the form `<module>/v<semver>`, e.g. `logging/v0.3.0`.

The release workflow:
1. Commits follow [Conventional Commits](https://www.conventionalcommits.org/) format
2. `t0changelog` generates the changelog section from `git log` since the last tag
3. `t0copyright` ensures all files have the BSD-3-Clause header
4. Tags are pushed; GitHub Actions handles the release PR and publishes

## Code style

- **Formatting:** `gofumpt` (stricter than `gofmt`)
- **Linting:** `golangci-lint` — config in `.golangci.yml` (if present) or defaults
- **Error wrapping:** use `misc.Wrap`, `misc.Wrapf`, `misc.Wrapfl` (never `fmt.Errorf` with `%w` directly)
- **Logging:** always pass `*logging.Logger` via context; retrieve with `logging.FromContext(ctx)`
- **Context pattern:** `clock`, `identifier`, and `logging` all use the same store-in-context / retrieve-from-context pattern
- **Generics:** preferred over `any`-typed helpers wherever the type is constrained
- **No wiki:** documentation lives in this `doc/` folder and inline code comments

## Environment variables

Create a `.env` file at the repo root (gitignored) to set defaults:

```bash
T0_COLOR=true
T0_LOGLEVEL=1   # Debug
```

All commands automatically load `.env` via `misc.DotEnv` on startup.

## sh/ submodule

`sh/lib` is a git submodule containing shared bash utilities (`lib.sh`, `ci.sh`, helper scripts). It is sourced via `BASH_ENV=./sh/lib/lib.sh` in most runner tasks. After a fresh clone:

```bash
git submodule update --init
```
