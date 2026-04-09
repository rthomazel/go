# Development Guide

## Prerequisites

- **Go 1.26.1+** (managed via [mise](https://mise.jdx.dev/) + `.tool-versions`)
- **Node.js** — required for prettier (YAML/JSON formatting)
- **prettier** — `npm install -g prettier`

All Go tools (golangci-lint, gofumpt, godotenv) are declared in `go.mod` and installed via:

```bash
go install tool
```

Run setup to handle all of the above:

```bash
./bin/setup      # or: ./run setup
```

This is idempotent and safe to re-run. In cloud/CI environments it auto-detects via `GITHUB_TOKEN` and installs the full toolchain.

## Common tasks

```bash
./run help            # list all commands

./run build logging   # build a module
./run install cmd     # build and install a module binary
./run test misc       # run tests for a module
./run test:all        # run tests for all modules

./run lint cmd        # lint a module
./run format httpmisc # format a module
./run format          # format all go modules

./run copyright:go    # fix copyright headers in .go files
./run copyright:sh    # fix copyright headers in .sh files

./run generate:gowork # regenerate go.work (after adding/removing modules)
./run new:module      # scaffold a new library module
./run new:command     # scaffold a new CLI command
```

## Adding a new library module

```bash
./run new:module
```

The script (`sh/new_module.sh`) scaffolds the module directory, `go.mod`, and initial source files, then regenerates `go.work`.

## Adding a new CLI command (gengowork-style)

```bash
./run new:command
```

The script (`sh/new_command.sh`) copies the `cmd/template` directory and wires it into the workspace.

## Module versioning & releases

Each sub-module is versioned independently using Git tags in the form `<module>/v<semver>`, e.g. `logging/v0.3.0`.

The release workflow:

1. Commits follow [Conventional Commits](https://www.conventionalcommits.org/) format
2. Tags are pushed; GitHub Actions handles the release PR and publishes

## Code style

- **Formatting:** `gofumpt` (stricter than `gofmt`) — run via `./run format <module>`
- **Linting:** `golangci-lint` — run via `./run lint <module>`
- **Error wrapping:** use `misc.Wrap`, `misc.Wrapf`, `misc.Wrapfl`
- **Logging:** pass `*logging.Logger` via context; retrieve with `logging.FromContext(ctx)`
- **Context pattern:** `clock`, `identifier`, and `logging` all use store-in-context / retrieve-from-context
- **No wiki:** documentation lives in this `doc/` folder and inline code comments

## Environment variables

Create a `.env` file at the repo root (gitignored) to set local defaults:

```bash
T0_COLOR=true
T0_LOGLEVEL=1   # Debug
```

All commands load `.env` automatically via `godotenv` or `misc.DotEnv`.

## sh/ submodule

`sh/lib` is a git submodule containing shared bash utilities (`lib.sh`, `ci.sh`, helper scripts). After a fresh clone:

```bash
git submodule update --init
```
