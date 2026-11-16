package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v4"
)

type Pipeline struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name     string   `yaml:"name"`
	Image    string   `yaml:"image"`
	Commands []string `yaml:"commands"`
}

func loadPipeline(path string) (Pipeline, error) {
	var pipeline Pipeline

	data, err := os.ReadFile(path)
	if err != nil {
		return pipeline, fmt.Errorf("read pipeline: %w", err)
	}

	err = yaml.Load(data, &pipeline, yaml.WithKnownFields())
	if err != nil {
		return pipeline, fmt.Errorf("parse pipeline: %w", err)
	}

	err = validate(pipeline)
	if err != nil {
		return pipeline, fmt.Errorf("invalid pipeline: %w", err)
	}

	return pipeline, nil
}

func validate(pipeline Pipeline) error {
	if len(pipeline.Steps) == 0 {
		return errors.New("no steps")
	}

	for i, step := range pipeline.Steps {
		if step.Name == "" {
			return fmt.Errorf("step %d: missing name", i+1)
		}
		if step.Image == "" {
			return fmt.Errorf("step %d (%s): missing image", i+1, step.Name)
		}
		if strings.HasPrefix(step.Image, "-") {
			return fmt.Errorf("step %d (%s): image must not start with -", i+1, step.Name)
		}
		if len(step.Commands) == 0 {
			return fmt.Errorf("step %d (%s): no commands", i+1, step.Name)
		}
	}

	return nil
}
