# CLI Commands

All commands live under `cmd/` and are built from `github.com/tcodes0/go/cmd`. They share a common pattern: `T0_COLOR` / `T0_LOGLEVEL` env vars for output control, a `-v`/`-version` flag, and a `passAway` recover-on-main pattern.

The `run` wrapper at the repo root is the main entrypoint for local development tasks:

```bash
./run <task> [package] [args...]   # delegates to .build/t0runner
./run ci [ci-task]                  # delegates to sh/lib/ci.sh
```

---

## t0runner

**Purpose:** YAML-configured task runner. Replaces make/task for this repo.

**Usage:**
```
t0runner [-v] [-config file] <task> [package] [extra args...]
```

**Flags:**
| Flag | Description |
|------|-------------|
| `-config` | Path to config file (default: `.t0runnerrc.yml` or `.t0runnerrc.yaml`) |
| `-v`, `-version` | Print version and exit |

**Config format** (`.t0runnerrc.yml`):
```yaml
version: "0.2.0"
tasks:
  - name: build
    package: true          # requires a package name as second argument
    env:
      - HOME=<inherit>     # copy value from current environment
      - PATH=<inherit>
    exec:
      - ./sh/workflows/module_pr/build.sh <package>   # <package> = the package arg
```

**Template variables:**
| Variable | Replaced with |
|----------|---------------|
| `<package>` | The package name argument passed on the command line |
| `<inherit>` | The current value of that env var from the parent process |
| `<space>` | A single space (workaround for YAML parsing) |

**Tasks defined in this repo** (from `.t0runnerrc.yml`):

| Task | Package? | Description |
|------|----------|-------------|
| `build` | yes | Compile a module |
| `install` | yes | Build and install a module binary |
| `lint` | yes | Run `golangci-lint` on a module |
| `go-lint-fix` | yes | Auto-fix lint issues |
| `format` | yes | Run `gofumpt` + prettier on a module |
| `test` | yes | Run tests with coverage |
| `format-configs` | no | Prettier all YAML/JSON configs |
| `generate-mocks` | no | Generate mock files via `sh/lib/go/generate_mocks.sh` |
| `spellcheck` | no | Run `cspell` over the repo |
| `setup` | no | Run `sh/setup.sh` to verify/install tooling |
| `generate-go-work` | no | Regenerate `go.work` via `cmd/gengowork` |
| `new-module` | no | Scaffold a new library module |
| `new-command` | no | Scaffold a new CLI command |
| `copyright-fix-go` | no | Add/fix copyright headers in `.go` files |
| `copyright-fix-sh` | no | Add/fix copyright headers in `.sh` files |

**Example:**
```bash
./run build logging
./run test misc
./run lint cmd
./run format-configs
```

Unknown task names trigger a fuzzy "did you mean" suggestion.

---

## t0changelog

**Purpose:** Generate a `CHANGELOG.md` section from `git log` output and GitHub PR data. Parses [Conventional Commits](https://www.conventionalcommits.org/).

**Usage:**
```
t0changelog -url <github-url> [-title <title>] [-tagprefixes <prefixes>] [-config <file>] [-tagsfile <file>]
```

**Flags:**
| Flag | Default | Description |
|------|---------|-------------|
| `-url` | required | GitHub repo URL (must begin with `https://github.com/`) |
| `-title` | `""` | Release title; version and date are appended automatically |
| `-tagprefixes` | `""` | Comma-separated tag prefixes to find releases (e.g. `logging/,misc/`) |
| `-config` | `.commitlintrc.yml` | Path to commitlint config |
| `-tagsfile` | `""` | If set, writes the new tags found to this file |
| `-v`, `-version` | | Print version and exit |

**Commit type → section name mapping** (from `config.yml`):

| Commit type | Section |
|-------------|----------|
| `feat` | Features |
| `fix` | Bug Fixes |
| `perf` | Performance |
| `refactor` | Improvements |
| `docs` | Documentation |
| `build` | Build |
| `ci` | CI |
| `revert` | Revert |
| `style` | Styling |
| `test` | Tests |
| `misc`, `chore` | Other |

Breaking changes (commits with `!` or a `BREAKING CHANGE:` footer) are collected into a separate **Breaking Changes** section at the top.

GitHub PR numbers are fetched in parallel from the GitHub API (reads `GITHUB_TOKEN` if set).

---

## t0copyright

**Purpose:** Check or add copyright header comments to source files.

**Usage:**
```
t0copyright -check <glob> [-fix] [-comment <token>] [-shebang] [-ignore <regexp>]
```

**Flags:**
| Flag | Description |
|------|-------------|
| `-check` | Glob pattern of files to check (required, e.g. `*.go`) |
| `-fix` | Write header to files that are missing it; requires `-comment` |
| `-comment` | Token prepended to each header line (e.g. `//` for Go, `#` for shell) |
| `-shebang` | Preserve the first line of the file (shebang), insert header after it |
| `-ignore` | Regexp to ignore matching paths (default: `/?mock_.*|.local/.*`) |
| `-v`, `-version` | Print version and exit |

The header text is embedded from `config.yml` at build time.

**Example:**
```bash
# Check all Go files
t0copyright -check *.go

# Fix Go files
t0copyright -check *.go -fix -comment "// "

# Fix shell scripts (with shebang preservation)
t0copyright -check *.sh -fix -comment "# " -shebang
```

---

## t0filer

**Purpose:** File management — batch symlink creation, deletion, and backup via a config file.

**Usage:**
```
t0filer -config <file> [-commit]
```

**Flags:**
| Flag | Description |
|------|-------------|
| `-config` | Path to filer config file (required) |
| `-commit`, `-c` | Apply changes (dry-run by default) |
| `-v`, `-version` | Print version and exit |

**Actions:**
| Action | Description | Config format |
|--------|-------------|---------------|
| `link` | Create symlinks | Pairs of lines: source then link path |
| `remove` | Delete files | One file per line |
| `backup` | Copy with `.bak` extension | One file per line |

The config file starts with an action name on the first line, followed by file paths.

**Example config:**
```
link
/home/user/.config/real-config
/home/user/.config/symlink-config
```

Without `-commit`, the tool prints what it would do. With `-commit`, changes are applied.

---

## gengowork (internal tool)

**Purpose:** Regenerate `go.work` and `go.work.sum` from the current module layout. Used via the `generate-go-work` runner task.

**Usage:**
```bash
./run generate-go-work
# or:
go run cmd/gengowork/main.go
```

Scans for all `go.mod` files, determines versions, and writes a fresh `go.work` with the correct `use` directives. The output file header warns `// generated do not edit.`
