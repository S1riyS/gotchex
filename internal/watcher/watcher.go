package watcher

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/S1riyS/gotchex/internal/config"
	"github.com/fsnotify/fsnotify"
)

type Runner interface {
	Run() error
}

type FileWatcher struct {
	watcher              *fsnotify.Watcher
	config               *config.Config
	lastRun              time.Time
	compiledIncludeRegex []*regexp.Regexp
	compiledExcludeRegex []*regexp.Regexp
	runner               Runner
}

func New(config *config.Config, runner Runner) (*FileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("failed to create watcher: %w", err)
	}

	fw := &FileWatcher{
		watcher:              watcher,
		config:               config,
		compiledIncludeRegex: make([]*regexp.Regexp, len(config.Build.IncludeRegex)),
		compiledExcludeRegex: make([]*regexp.Regexp, len(config.Build.ExcludeRegex)),
		runner:               runner,
	}

	// Compile regex patterns once
	for i, pattern := range config.Build.IncludeRegex {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid include regex pattern %q: %w", pattern, err)
		}
		fw.compiledIncludeRegex[i] = re
	}

	for i, pattern := range config.Build.ExcludeRegex {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid exclude regex pattern %q: %w", pattern, err)
		}
		fw.compiledExcludeRegex[i] = re
	}

	// Add directories to watch
	for _, dir := range config.Build.IncludeDir {
		if err := fw.addWatchDir(dir); err != nil {
			log.Printf("Warning: %v", err)
		}
	}

	return fw, nil
}

func (fw *FileWatcher) addWatchDir(dir string) error {
	// Skip excluded directories
	for _, exclude := range fw.config.Build.ExcludeDir {
		if matched, _ := filepath.Match(exclude, dir); matched {
			return fmt.Errorf("skipping excluded directory: %s", dir)
		}
	}

	if err := fw.watcher.Add(dir); err != nil {
		return fmt.Errorf("failed to watch directory %s: %w", dir, err)
	}

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			for _, exclude := range fw.config.Build.ExcludeDir {
				if matched, _ := filepath.Match(exclude, path); matched {
					return filepath.SkipDir
				}
			}
			return fw.watcher.Add(path)
		}
		return nil
	})
}

func (fw *FileWatcher) Start() error {
	defer fw.watcher.Close()

	for {
		select {
		case event, ok := <-fw.watcher.Events:
			if !ok {
				return nil
			}

			if !fw.shouldHandleEvent(event) {
				continue
			}

			// Debounce logic
			if time.Since(fw.lastRun) < time.Duration(fw.config.Build.Delay)*time.Millisecond {
				continue
			}

			fw.lastRun = time.Now()
			go fw.executeCommand()

		case err, ok := <-fw.watcher.Errors:
			if !ok {
				return nil
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}

func (fw *FileWatcher) shouldHandleEvent(event fsnotify.Event) bool {
	// Only handle write events
	if event.Op&fsnotify.Write != fsnotify.Write {
		return false
	}

	filename := filepath.Base(event.Name)

	// Check against exclude patterns first
	for _, re := range fw.compiledExcludeRegex {
		if re.MatchString(filename) {
			return false
		}
	}

	// If no include patterns, allow all non-excluded files
	if len(fw.compiledIncludeRegex) == 0 {
		return true
	}

	// Check against include patterns
	for _, re := range fw.compiledIncludeRegex {
		if re.MatchString(filename) {
			return true
		}
	}

	return false
}

func (fw *FileWatcher) executeCommand() {
	log.Printf("Executing command: %s", fw.config.Build.Command)
	fw.runner.Run()
}

func (fw *FileWatcher) PrintWatchedFiles() error {
	fmt.Println("=== Watched Files ===")

	// Check all included directories
	for _, dir := range fw.config.Build.IncludeDir {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skip directories and excluded files
			if info.IsDir() {
				for _, exclude := range fw.config.Build.ExcludeDir {
					if matched, _ := filepath.Match(exclude, path); matched {
						return filepath.SkipDir
					}
				}
				return nil
			}

			// Check file against exclude patterns
			filename := filepath.Base(path)
			for _, re := range fw.compiledExcludeRegex {
				if re.MatchString(filename) {
					return nil
				}
			}

			// Check against include patterns (if any exist)
			if len(fw.compiledIncludeRegex) > 0 {
				matched := false
				for _, re := range fw.compiledIncludeRegex {
					if re.MatchString(filename) {
						matched = true
						break
					}
				}
				if !matched {
					return nil
				}
			}

			// If we get here, the file is being watched
			fmt.Println(path)
			return nil
		})

		if err != nil {
			return fmt.Errorf("error scanning %s: %w", dir, err)
		}
	}
	return nil
}
