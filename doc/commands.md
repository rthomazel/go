# CLI Commands

The `run` script at the repo root is the main entrypoint for local development.

```bash
./run help   # list all commands
```

---

## cmd/copyright

**Purpose:** Check and optionally fix copyright headers in source files. Stdlib only.

**Usage:**

```
copyright -c TOKEN [-f] [-s] [-i PATTERN] GLOB[,GLOB...]
```

**Flags:**
| Flag | Description |
|------|-------------|
| `-c TOKEN` | Comment token prepended to each header line (required) |
| `-f` | Fix: write header to files missing it |
| `-s` | Shebang: preserve first line, insert header after it |
| `-i PATTERN` | Extended regexp to ignore paths (default: `mock_|/.local/`) |
| `GLOB,...` | Comma-separated filename globs (required) |

Reads the header text from the `COPYRIGHT_HEADER` env var, or a file named
`copyright-header` next to the binary or in the working directory.

```bash
# check only
go run cmd/copyright/main.go -c '// ' '*.go'

# fix Go files
go run cmd/copyright/main.go -f -c '// ' '*.go'

# fix shell scripts (shebang-aware)
go run cmd/copyright/main.go -f -s -c '# ' '*.sh'

# via run script
./run copyright:go
./run copyright:sh
```

---

## sh/generate_gowork.sh

**Purpose:** Regenerate `go.work` from the current module layout. Reads the Go
version from `go.mod`, finds all top-level sub-module directories, and writes
a fresh `go.work`. Called via `./run generate:gowork`.
