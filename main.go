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

func main() {
	cmd := exec.Command("go", "fmt", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "jogger: %v\n", err)
		os.Exit(1)
	}

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
}
