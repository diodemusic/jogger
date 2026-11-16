package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func cloneRepo(ctx context.Context, repo string, workspace string) error {
	cmd := exec.CommandContext(ctx, "git", "clone", "--quiet", repo, workspace)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("clone %s: %w", repo, err)
	}

	return nil
}
