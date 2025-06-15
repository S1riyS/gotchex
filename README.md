# Gotchex

## 📌 About
**Gotchex** is a lightweight file watcher utility in Go that automatically executes shell commands when files change

**Key features**:
- 🔍 Monitors directories/files recursively
- ⚡ Triggers custom commands (e.g. `go run main.go`, `npm start`) on changes
- ⚙️ Configurable via CLI or yaml file (patterns, delay, ignored files, etc...)
- 🚀 Minimal dependencies (only `fsnotify` and `yaml.v3`, pure Go, cross-platform)

## 🛠️ Technologies
- [Go](https://go.dev/) - Language
    - [fsnotify](https://github.com/fsnotify/fsnotify) - Filesystem watcher
    - [yaml.v3](https://github.com/go-yaml/yaml) - YAML parser
- Tools:
    - [Taskfile](https://taskfile.dev/) - Task runner / Build tool