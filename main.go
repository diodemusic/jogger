package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: jogger <repo>")
		os.Exit(2)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	err := run(ctx, os.Args[1])
	interrupted := ctx.Err() != nil
	stop()

	if interrupted {
		fmt.Fprintln(os.Stderr, "jogger: interrupted")
		os.Exit(130)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "jogger:", err)
		os.Exit(1)
	}

	fmt.Println("pipeline passed")
}

func run(ctx context.Context, repo string) error {
	workspace, err := os.MkdirTemp("", "jogger-")
	if err != nil {
		return fmt.Errorf("create workspace: %w", err)
	}

	defer os.RemoveAll(workspace)

	err = cloneRepo(ctx, repo, workspace)
	if err != nil {
		return err
	}

	pipeline, err := loadPipeline(filepath.Join(workspace, ".jogger.yml"))
	if err != nil {
		return err
	}

	for i, step := range pipeline.Steps {
		fmt.Printf("step %d: %s (image %s)\n", i+1, step.Name, step.Image)

		err := runStep(ctx, workspace, step)
		if err != nil {
			return err
		}
	}

	return nil
}
