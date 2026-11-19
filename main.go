package main

import (
	"fmt"
	"os"
	"os/exec"

	"go.yaml.in/yaml/v4"
)

type Steps struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name     string   `yaml:"name"`
	Image    string   `yaml:"image"`
	Commands []string `yaml:"commands"`
}

func run(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func main() {
	file, err := os.ReadFile(".jogger.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "jogger: %v\n", err)
		os.Exit(1)
	}

	var steps Steps

	if err := yaml.Load(file, &steps, yaml.WithKnownFields()); err != nil {
		fmt.Fprintf(os.Stderr, "jogger: %v\n", err)
		os.Exit(1)
	}

	for _, step := range steps.Steps {
		for _, command := range step.Commands {
			err := run(command)
			if err != nil {
				fmt.Fprintf(os.Stderr, "jogger: %s: %q: %v\n", step.Name, command, err)
				os.Exit(1)
			}
		}
	}
}
