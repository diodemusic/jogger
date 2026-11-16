package main

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

func cloneRepo(ctx context.Context, out io.Writer, repo string, workspace string) error {
	cmd := exec.CommandContext(ctx, "git", "clone", "--quiet", "--no-hardlinks", "--", repo, workspace)
	cmd.Stdout = out
	cmd.Stderr = out

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("clone %s: %w", repo, err)
	}

	return nil
}
