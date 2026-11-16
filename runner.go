package main

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
)

func runStep(ctx context.Context, out io.Writer, workspace string, step Step) error {
	script := strings.Join(step.Commands, "\n")
	containerName := filepath.Base(workspace)

	cmd := exec.CommandContext(
		ctx,
		"docker", "run",
		"--rm",
		"--name", containerName,
		"--volume", workspace+":/workspace",
		"--workdir", "/workspace",
		step.Image,
		"sh", "-e", "-x", "-c", script,
	)
	cmd.Stdout = out
	cmd.Stderr = out
	cmd.Cancel = func() error {
		return exec.Command("docker", "kill", containerName).Run()
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("step %s: %w", step.Name, err)
	}

	return nil
}
