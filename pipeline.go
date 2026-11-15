package main

import (
	"fmt"
	"os"

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

	err = yaml.Load(data, &pipeline)
	if err != nil {
		return pipeline, fmt.Errorf("parse pipeline: %w", err)
	}

	return pipeline, nil
}
