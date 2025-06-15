package runner

import (
	"os"
	"os/exec"
)

type Runner struct {
	command string
}

func New(command string) *Runner {
	return &Runner{command: command}
}

func (r *Runner) Run() error {
	cmd := exec.Command("sh", "-c", r.command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
