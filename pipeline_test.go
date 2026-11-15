package main

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		pipeline Pipeline
		wantErr  string
	}{
		{
			name: "valid",
			pipeline: Pipeline{Steps: []Step{
				{Name: "test", Image: "golang:1.27", Commands: []string{"go test ./..."}},
			}},
			wantErr: "",
		},
		{
			name:     "no steps",
			pipeline: Pipeline{},
			wantErr:  "no steps",
		},
		{
			name: "missing name",
			pipeline: Pipeline{Steps: []Step{
				{Image: "golang:1.27", Commands: []string{"go test ./..."}},
			}},
			wantErr: "step 1: missing name",
		},
		{
			name: "missing image",
			pipeline: Pipeline{Steps: []Step{
				{Name: "test", Commands: []string{"go test ./..."}},
			}},
			wantErr: "step 1 (test): missing image",
		},
		{
			name: "no commands",
			pipeline: Pipeline{Steps: []Step{
				{Name: "test", Image: "golang:1.27"},
			}},
			wantErr: "step 1 (test): no commands",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.pipeline)

			got := ""
			if err != nil {
				got = err.Error()
			}

			if got != tt.wantErr {
				t.Errorf("got error %q, want %q", got, tt.wantErr)
			}
		})
	}
}
