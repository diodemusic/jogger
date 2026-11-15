package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: jogger <repo-path>")
		os.Exit(2)
	}

	err := run(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "jogger:", err)
		os.Exit(1)
	}

	fmt.Println("pipeline passed")
}

func run(repoPath string) error {
	pipeline, err := loadPipeline(filepath.Join(repoPath, ".jogger.yml"))
	if err != nil {
		return err
	}

	workspace, err := os.MkdirTemp("", "jogger-")
	if err != nil {
		return fmt.Errorf("create workspace: %w", err)
	}
	defer os.RemoveAll(workspace)

	err = os.CopyFS(workspace, os.DirFS(repoPath))
	if err != nil {
		return fmt.Errorf("copy repo to workspace: %w", err)
	}

	for i, step := range pipeline.Steps {
		fmt.Printf("step %d: %s (image %s)\n", i+1, step.Name, step.Image)

		err := runStep(workspace, step)
		if err != nil {
			return err
		}
	}

	return nil
}
