# fc-consul

fc-consul is a consul configuration provider for [fc](https://github.com/flowchartsman/fc)

## Usage
```go
package consul

import (
	"github.com/flowchartsman/fc"
	"github.com/flowchartsman/fc-consul"
)

func main() {
	fs := flag.NewFlagSet("my-program", flag.ExitOnError)
	var (
		listenAddr = fs.String("listen-addr", "localhost:8080", "listen address")
		refresh    = fs.Duration("refresh", 15*time.Second, "refresh interval")
		debug      = fs.Bool("debug", false, "log debug information")
	)

	fc.Parse(fs,
		fc.WithEnv("MY_PROGRAM"),
		consul.WithNode("localhost:8500", "my/prefix"),
	)
}
```
