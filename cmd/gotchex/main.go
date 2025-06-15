package main

import (
	"log"

	"github.com/S1riyS/gotchex/internal/config"
	"github.com/S1riyS/gotchex/internal/runner"
	"github.com/S1riyS/gotchex/internal/watcher"
)

func main() {
	cfg := &config.Config{
		Build: config.BuildConfig{
			Command:      "echo 'File changed!'",
			Delay:        1000,
			IncludeDir:   []string{"."},
			IncludeRegex: []string{".go"},
			ExcludeDir:   []string{".git", "tmp"},
			ExcludeRegex: []string{},
		},
	}

	commandRunner := runner.New(cfg.Build.Command)
	fw, err := watcher.New(cfg, commandRunner)
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}

	fw.PrintWatchedFiles()

	log.Println("Starting file watcher...")
	if err := fw.Start(); err != nil {
		log.Fatalf("Watcher failed: %v", err)
	}
}
