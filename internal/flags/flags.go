package flags

import (
	"flag"
	"fmt"
	"sync"
)

const usage = `Usage: gotchex [flags...]

Starts a file watcher that automatically executes shell commands when files change.

Examples:
1. 'gotchex' will use the default config file (./gotchex.yaml)
2. 'gotchex -config-path=./gotchex.go.yaml' will use the config file at './gotchex.go.yaml'

'''
watch:
  delay: 1000
  include_dir: ["."]
  include_regex: [".*\\.go$", ".*\\.mod$"]
  exclude_dir: [./vendor, ./tmp]
  exclude_regex: [".*_test\\.go$", ".*\\.tmp$"]

run:
  build: "go build -o app cmd/main.go"
  run: "./app"
'''

Options:
`

// Flags
var (
	ConfigPath string
)

// Sync
var once sync.Once

func Load() {
	once.Do(func() {
		// Flag Definitions
		flag.StringVar(&ConfigPath, "config-path", "", "Path to config file")

		// Usage
		flag.Usage = func() {
			fmt.Print(usage)
			flag.PrintDefaults()
		}

		flag.Parse()
	})
}
