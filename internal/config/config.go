package config

type Config struct {
	Build BuildConfig `yaml:"build"`
}

type BuildConfig struct {
	Command      string   `yaml:"command"`
	Delay        int      `yaml:"delay"`
	IncludeDir   []string `yaml:"include_dir"`
	IncludeRegex []string `yaml:"include_regex"`
	ExcludeDir   []string `yaml:"exclude_dir"`
	ExcludeRegex []string `yaml:"exclude_regex"`
}
