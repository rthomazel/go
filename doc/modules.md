# Modules

Each directory is an independently versioned Go module published at `github.com/rthomazel/go/<name>`. Modules follow semver. The latest published versions are listed below as of the last changelog entry (September 2024); run `go get github.com/rthomazel/go/<name>@latest` to pull the most recent release.

---

## `hue` — Terminal colors

**Module:** `github.com/rthomazel/go/hue`
**No external dependencies.**

Thin wrapper around ANSI 256-color escape codes. Provides named color constants and helpers for colorizing terminal output.

```go
import "github.com/rthomazel/go/hue"

// Colorize a string and reset:
fmt.Println(hue.Printc(hue.Red, "error:") + hue.End + " something went wrong")

// Get the raw escape code:
code := hue.TermColor(hue.Yellow) // "\033[38;05;215m"
```

Exported color names: `Gray`, `Brown`, `BrightRed`, `Red`, `Yellow`, `Blue`.
`End` terminates any active color/format sequence.

---

## `clock` — Testable time

**Module:** `github.com/rthomazel/go/clock`
**No external dependencies.**

Defines the `Nower` interface so callers can swap real time for a fixed value in tests.

```go
import "github.com/rthomazel/go/clock"

// Production: real time
nower := &clock.Time{Location: *time.UTC}
ctx = nower.WithContext(ctx)

// Tests: fixed time
nower := clock.Static{Source: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}

// Retrieve from context:
nower, err := clock.FromContext(ctx)
```

**Types:** `Nower` (interface), `Time` (real time), `Static` (fixed time).

---

## `logging` — Structured leveled logger

**Module:** `github.com/rthomazel/go/logging`
**Depends on:** `hue`

Wraps `log.Logger` with levels (Debug, Info, Warn, Error, Fatal, None), optional ANSI color output, key/value data maps, and context propagation.

```go
import "github.com/rthomazel/go/logging"

logger := logging.Create(logging.OptLevel(logging.LDebug), logging.OptColor())
ctx = logger.WithContext(ctx)

logger.Infof("starting server on %s", addr)
logger.ErrorData(map[string]any{"err": err, "path": path}, "handler failed")

// Retrieve from context:
logger := logging.FromContext(ctx)
```

**Levels (lowest to highest):** `LDebug=1`, `LInfo=2`, `LWarn=3`, `LError=4`, `LFatal=5`, `LNone=6`.
Messages below the configured level are silently dropped.

Controlled at runtime via the `T0_LOGLEVEL` environment variable.

---

## `misc` — General utilities

**Module:** `github.com/rthomazel/go/misc`
**No internal deps; depends on `testify` for tests.**

A grab-bag of small, composable helpers. All generic where possible.

| File        | Exported API                            | Description                                               |
| ----------- | --------------------------------------- | --------------------------------------------------------- |
| `err.go`    | `Wrap`, `Wrapf`, `Wrapfl`               | Error wrapping with message, format, or file:line         |
| `env.go`    | `LookupEnv[T]`, `DotEnv`                | Generic env var lookup with fallback; `.env` file loading |
| `envtag.go` | `ApplyToFields` / `FieldUpdater`        | Struct field iteration via reflection                     |
| `slice.go`  | `Find[T]`, `Uniq[T]`                    | Generic slice find and dedup                              |
| `pick.go`   | `PickValid[T]`, `Default[T]`            | First non-nil/non-zero value; or default                  |
| `copy.go`   | `ToPtr[T]`, `Copy[T]`, `CopyPointed[T]` | Pointer/value copy helpers                                |
| `nil.go`    | `IsNil`, `IsZero`                       | Panic-safe nil/zero checks via reflection                 |
| `merge.go`  | `Merge[T]`                              | Struct partial update (base + partial, with ignore list)  |
| `field.go`  | `ApplyToFields[T]`                      | Apply a `FieldUpdater` to all struct fields               |
| `signal.go` | `RoutineHandleStopSignal`               | Block until SIGINT/SIGTERM/SIGHUP then call handler       |
| `time.go`   | `Seconds`, `Minutes`, `Hours`, `Days`   | `time.Duration` constructors from integers                |

---

## `identifier` — ID generation

**Module:** `github.com/rthomazel/go/identifier`
**Depends on:** `google/uuid`

Defines a `Generator` interface and a UUID-backed implementation. Context-propagation pattern mirrors `clock` and `logging`.

```go
import "github.com/rthomazel/go/identifier"

gen := &identifier.UUIDGenerator{}
ctx = gen.WithContext(ctx)

gen, err := identifier.FromContext(ctx)
id := gen.Generate() // e.g. "550e8400-e29b-41d4-a716-446655440000"
```

The interface makes it easy to inject deterministic IDs in tests.

---

## `jsonutil` — JSON helpers

**Module:** `github.com/rthomazel/go/jsonutil`
**Depends on:** `misc`

Small set of wrappers that remove boilerplate from common JSON I/O patterns.

```go
import "github.com/rthomazel/go/jsonutil"

// Marshal to an io.ReadCloser (handy for HTTP request bodies):
rc, err := jsonutil.MarshalReader(myStruct)

// Marshal and build an *http.Request in one step:
req, err := jsonutil.MarshalRequest(ctx, http.MethodPost, url, myStruct)

// Unmarshal from a response body:
result, err := jsonutil.UnmarshalReader[MyType](resp.Body)

// Unmarshal from raw bytes:
result, err := jsonutil.UnmarshalBytes[MyType](data)
```

---

## `httpmisc` — HTTP client & middleware

**Module:** `github.com/rthomazel/go/httpmisc`
**Depends on:** `logging`, `misc`

Higher-level HTTP helpers for both client and server sides.

### Client

`Client` wraps `*http.Client` with API key auth, base URL, user agent, and request timeout. All standard HTTP methods are available.

```go
c := &httpmisc.Client{}
err := c.Init(&httpmisc.SetClientOptions{
    BaseURL:   "https://api.example.com",
    APIKey:    token,
    UserAgent: "myapp/1.0",
})
resp, data, err := c.Get(ctx, "/v1/resource", nil, nil)
```

### Server middleware / utilities

| Symbol            | Description                                                             |
| ----------------- | ----------------------------------------------------------------------- |
| `Recoverer(next)` | `http.Handler` middleware — catches panics, logs them, returns 500      |
| `MaxSize`         | `http.ResponseWriter` wrapper — auto-flushes after `Max` bytes written  |
| `Roundtrip`       | `http.RoundTripper` with debug logging of request and response metadata |
