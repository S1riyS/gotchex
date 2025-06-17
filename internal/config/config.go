package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Watch *WatchConfig `yaml:"watch"`
	Run   *RunConfig   `yaml:"run"`
}

type WatchConfig struct {
	Delay        int      `yaml:"delay"`
	IncludeDir   []string `yaml:"include_dir"`
	IncludeRegex []string `yaml:"include_regex"`
	ExcludeDir   []string `yaml:"exclude_dir"`
	ExcludeRegex []string `yaml:"exclude_regex"`
}

type RunConfig struct {
	Build *string `yaml:"build"`
	Run   string  `yaml:"run"`
}

const DEFAULT_CONFIG_PATH = "gotchex.yaml"

func MustLoad(configPath string) *Config {
	if configPath == "" {
		configPath = DEFAULT_CONFIG_PATH
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config
	content, err := os.ReadFile(configPath)
	if err != nil {
		panic("Error reading file: " + err.Error())
	}

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		panic("Error unmarshalling config")
	}

	return &cfg
}
