# Project Overview

`github.com/tcodes0/go` is a personal Go library monorepo by Raphael Thomazella (tcodes0). It contains a collection of small, focused, independently-versioned Go modules that address common infrastructure concerns: logging, HTTP helpers, time abstractions, error wrapping, JSON utilities, terminal colors, and ID generation.

The project also ships four CLI tools (under `cmd/`) and a task runner configuration that ties together linting, building, testing, and release automation.

## Repository layout

```
go/
├── clock/        # Nower interface for testable time
├── cmd/          # CLI tools (t0runner, t0changelog, t0copyright, t0filer, gengowork)
├── httpmisc/     # HTTP client, middleware, and transport helpers
├── hue/          # ANSI terminal color utilities
├── identifier/   # ID/UUID generation interface and implementation
├── jsonutil/     # JSON marshal/unmarshal helpers
├── logging/      # Structured, leveled logger wrapping log.Logger
├── misc/         # General-purpose utilities (errors, env, slices, generics, …)
├── sh/           # Shared bash library (submodule)
├── doc/          # Project documentation (this folder)
├── go.mod        # Root module (github.com/tcodes0/go)
├── go.work       # Go workspace including all sub-modules
└── run           # Entrypoint wrapper: delegates to t0runner or ci.sh
```

## Module graph

Dependencies between the repo's own modules:

```
hue   (no internal deps)
clock (no internal deps)
misc  (no internal deps)
  └── jsonutil
  └── httpmisc → logging → hue
cmd → hue, jsonutil, logging, misc
```

## Tech stack

- **Go 1.26.1** — minimum version for all modules
- **Go workspaces** (`go.work`) — manages all sub-modules locally
- **golangci-lint** — linting
- **gofumpt** — formatting
- **t0runner** — task runner (YAML config, built from `cmd/t0runner`)
- **bash + sh/lib submodule** — CI scripts
- GitHub Actions CI for lint, test, and release workflows

## Environment variables

| Variable       | Default  | Description                       |
|----------------|----------|-----------------------------------|
| `T0_COLOR`     | `false`  | Enable ANSI color in CLI output   |
| `T0_LOGLEVEL`  | `2` (Info) | Log level: 1=Debug 2=Info 3=Warn 4=Error 5=Fatal 6=None |

These can also be set in a `.env` file at the repo root (parsed by `misc.DotEnv`).
