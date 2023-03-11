# Go module for pretty-printing, v2

[![Go Reference](https://pkg.go.dev/badge/github.com/mateuszmidor/pretty/v2.svg)](https://pkg.go.dev/github.com/mateuszmidor/pretty/v2)

Print `os.Stdout` file:

```go
package main

import (
  "os"

  "github.com/mateuszmidor/pretty/v2"
)

func main() {
    const INDENT = 4
	pretty.Print(os.Stdout, INDENT)
}
```

Output:

```text
file:
    pfd:
        fdmu:
            state:
                0
            rsema:
                0
            wsema:
                0
        Sysfd:
            1
        pd:
            runtimeCtx:
                0
        iovecs:
            [empty]
        csema:
            0
        isBlocking:
            1
        IsStream:
            true
        ZeroReadIsEOF:
            true
        isFile:
            true
    name:
        /dev/stdout
    dirinfo:
        [empty]
    nonblock:
        false
    stdoutOrErr:
        true
    appendMode:
        false
```

## v2 comment

Version v2 of a Go module is basically a brand new module with own `go.mod` with `v2` suffix, `v2.0.0` git tag and under `v2` subdirectory:

```go
module github.com/mateuszmidor/pretty/v2

go 1.19
```

https://go.dev/blog/v2-go-modules

## Revisions

- v2.0.0: `func Print(resource interface{}, indentSize uint)`