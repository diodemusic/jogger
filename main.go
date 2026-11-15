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

	repoPath := os.Args[1]

	pipelinePath := filepath.Join(repoPath, ".jogger.yml")

	pipeline, err := loadPipeline(pipelinePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "jogger:", err)
		os.Exit(1)
	}

	for i, step := range pipeline.Steps {
		fmt.Printf("step %d: %s (image %s)\n", i+1, step.Name, step.Image)

		for _, command := range step.Commands {
			fmt.Println("  $", command)
		}
	}
}
