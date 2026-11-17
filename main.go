package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "fmt", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "jogger: %v\n", err)
		os.Exit(1)
	}
}
