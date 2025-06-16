package main

import (
	"log"

	"github.com/S1riyS/gotchex/internal/config"
	"github.com/S1riyS/gotchex/internal/flags"
	"github.com/S1riyS/gotchex/internal/runner"
	"github.com/S1riyS/gotchex/internal/watcher"
)

func main() {
	// Load flags
	flags.Load()

	// Load config
	cfg := config.MustLoad(flags.ConfigPath)

	commandRunner := runner.New(cfg.Run)
	fw, err := watcher.New(cfg.Watch, commandRunner)
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}

	if err = fw.PrintWatchedFiles(); err != nil {
		log.Printf("Failed to print watched files: %v", err)
	}

	log.Println("Starting file watcher...")
	if err := fw.Start(); err != nil {
		log.Fatalf("Watcher failed: %v", err)
	}
}
