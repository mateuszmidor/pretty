# Go module for pretty-printing

[![Go Reference](https://pkg.go.dev/badge/github.com/mateuszmidor/pretty.svg)](https://pkg.go.dev/github.com/mateuszmidor/pretty)

Print `os.Stdout` file:

```go
package main

import (
	"os"

	"github.com/mateuszmidor/pretty"
)

func main() {
	pretty.Print(os.Stdout)
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

## Revisions

- v1.0.0: `func Print(resource interface{})`