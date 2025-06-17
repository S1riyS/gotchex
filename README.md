<div id="readme-top"></div>

[//]: # (Project logo)
<br/>
<h1 align="center">
    <a href="https://github.com/S1riyS/go-whiteboard">
        <img src="docs/assets/logo.png" alt="Logo" width="140" height="140">
    </a>
    <br>
    <a href="https://goreportcard.com/report/github.com/S1riyS/gotchex">
        <img src="https://goreportcard.com/badge/github.com/S1riyS/gotchex">
    </a>
</h1>
<p align="center">
    <em>Gotchex is a lightweight file watcher utility in pure Go that automatically executes shell commands when files change</em>
</p>

---

## ✨ Key features:
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
    - [Golangci-lint](https://golangci-lint.run/) - Go linters runner

## 🚀 Getting Started
