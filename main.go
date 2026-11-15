package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: jogger <repo-path>")
		os.Exit(2)
	}

	repoPath := os.Args[1]
	fmt.Println("repo:", repoPath)
}
