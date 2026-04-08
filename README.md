<pre>

                                       ██████╗  ██████╗
                                      ██╔════╝ ██╔═══██╗
                                      ██║  ███╗██║   ██║
                                      ██║   ██║██║   ██║
                                      ╚██████╔╝╚██████╔╝
                                       ╚═════╝  ╚═════╝

      :::::::::     :::      ::::::::  :::    :::     :::      ::::::::  :::::::::: ::::::::
     :+:    :+:  :+: :+:   :+:    :+: :+:   :+:    :+: :+:   :+:    :+: :+:       :+:    :+:
    +:+    +:+ +:+   +:+  +:+        +:+  +:+    +:+   +:+  +:+        +:+       +:+
   +#++:++#+ +#++:++#++: +#+        +#++:++    +#++:++#++: :#:        +#++:++#  +#++:++#++
  +#+       +#+     +#+ +#+        +#+  +#+   +#+     +#+ +#+   +#+# +#+              +#+
 #+#       #+#     #+# #+#    #+# #+#   #+#  #+#     #+# #+#    #+# #+#       #+#    #+#
###       ###     ###  ########  ###    ### ###     ###  ########  ########## ########
</pre>
<details>
  <summary>art info </summary>
  http://www.patorjk.com/software/taag/#p=display&f=Alligator&t=packages
  http://www.patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=tcodes0%20go
</details>

---

Personal Go library monorepo. Small, independently versioned modules for common infrastructure concerns, plus a set of CLI tools.

**Go 1.26.1** &nbsp;|&nbsp; BSD-3-Clause

## Modules

| Module       | Import path                        | Description                                          |
| ------------ | ---------------------------------- | ---------------------------------------------------- |
| `hue`        | `github.com/tcodes0/go/hue`        | ANSI terminal colors                                 |
| `clock`      | `github.com/tcodes0/go/clock`      | Testable time (`Nower` interface)                    |
| `logging`    | `github.com/tcodes0/go/logging`    | Leveled, colored, context-aware logger               |
| `misc`       | `github.com/tcodes0/go/misc`       | General utilities (errors, env, generics, slices, …) |
| `identifier` | `github.com/tcodes0/go/identifier` | UUID ID generation interface                         |
| `jsonutil`   | `github.com/tcodes0/go/jsonutil`   | JSON marshal/unmarshal helpers                       |
| `httpmisc`   | `github.com/tcodes0/go/httpmisc`   | HTTP client, middleware, and transport helpers       |

## CLI tools

| Command                 | Description                                    |
| ----------------------- | ---------------------------------------------- |
| `cmd/copyright`         | Check/add copyright headers (Go, stdlib only)  |
| `sh/generate-gowork.sh` | Regenerate `go.work` from module layout (bash) |

## Quick start

```bash
git clone https://github.com/tcodes0/go
cd go
git submodule update --init
./run setup        # verify tooling
./run build misc   # build a module
./run test logging # test a module
```

## Documentation

- [Overview & architecture](doc/overview.md)
- [Module reference](doc/modules.md)
- [CLI commands](doc/commands.md)
- [Development guide](doc/development.md)

## License

BSD-3-Clause. See [LICENSE](LICENSE).
