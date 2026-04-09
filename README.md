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
  http://www.patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=rthomazel%20go
</details>

---

Personal Go library monorepo. Small, independently versioned modules for common infrastructure concerns, plus a set of CLI tools.

**Go 1.26.1** &nbsp;|&nbsp; BSD-3-Clause

## Modules

| Module       | Import path                          | Description                                          |
| ------------ | ------------------------------------ | ---------------------------------------------------- |
| `hue`        | `github.com/rthomazel/go/hue`        | ANSI terminal colors                                 |
| `clock`      | `github.com/rthomazel/go/clock`      | Testable time (`Nower` interface)                    |
| `logging`    | `github.com/rthomazel/go/logging`    | Leveled, colored, context-aware logger               |
| `misc`       | `github.com/rthomazel/go/misc`       | General utilities (errors, env, generics, slices, …) |
| `identifier` | `github.com/rthomazel/go/identifier` | UUID ID generation interface                         |
| `jsonutil`   | `github.com/rthomazel/go/jsonutil`   | JSON marshal/unmarshal helpers                       |
| `httpmisc`   | `github.com/rthomazel/go/httpmisc`   | HTTP client, middleware, and transport helpers       |

## CLI tools

| Command                 | Description                                    |
| ----------------------- | ---------------------------------------------- |
| `cmd/copyright`         | Check/add copyright headers (Go, stdlib only)  |
| `sh/generate_gowork.sh` | Regenerate `go.work` from module layout (bash) |

## Quick start

```bash
git clone https://github.com/rthomazel/go
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
