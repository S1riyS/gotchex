package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/S1riyS/gotchex/internal/config"
)

type Runner struct {
	config *config.RunConfig
}

func New(cfg *config.RunConfig) *Runner {
	return &Runner{config: cfg}
}

func (r *Runner) Run() error {
	// Excute build command
	if r.config.Build != nil {
		fmt.Println("Building...")
		err := setUpCmd(*r.config.Build).Run()
		if err != nil {
			return err
		}
	}

	// Execute run command
	fmt.Println("Running...")
	err := setUpCmd(r.config.Run).Run()
	if err != nil {
		return err
	}

	return nil
}

func setUpCmd(command string) *exec.Cmd {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
