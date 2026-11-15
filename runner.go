package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runStep(repoPath string, step Step) error {
	script := strings.Join(step.Commands, "\n")

	cmd := exec.Command(
		"docker", "run",
		"--rm",
		"--volume", repoPath+":/workspace",
		"--workdir", "/workspace",
		step.Image,
		"sh", "-e", "-x", "-c", script,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("step %s: %w", step.Name, err)
	}

	return nil
}
