package flags

import "flag"

var (
	ConfigPath string
)

func Load() {
	// Flags
	flag.StringVar(&ConfigPath, "config-path", "", "path to config file")

	flag.Parse()
}
