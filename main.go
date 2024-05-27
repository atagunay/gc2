package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getCurrentBranch() (string, error) {
	// Run the git command to get the current branch name
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Convert the output to a string and trim any whitespace
	branch := strings.TrimSpace(string(output))
	return branch, nil
}

func main() {
	branch, err := getCurrentBranch()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Current Git branch:", branch)
}
