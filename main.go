package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: jogger <repo>")
		os.Exit(2)
	}

	err := run(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "jogger:", err)
		os.Exit(1)
	}

	fmt.Println("pipeline passed")
}

func run(repo string) error {
	workspace, err := os.MkdirTemp("", "jogger-")
	if err != nil {
		return fmt.Errorf("create workspace: %w", err)
	}
	defer os.RemoveAll(workspace)

	err = cloneRepo(repo, workspace)
	if err != nil {
		return err
	}

	pipeline, err := loadPipeline(filepath.Join(workspace, ".jogger.yml"))
	if err != nil {
		return err
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
