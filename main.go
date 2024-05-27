package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

// getCurrentBranch runs a git command to get the current branch name
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

// commitWithBranchName combines the branch name and commit message and commits it
func commitWithBranchName(branchName string, commitMessage string) (string, error) {
	// Combine the branch name and commit message
	message := branchName + " " + commitMessage

	// Run the git commit command with the combined message
	cmd := exec.Command("git", "commit", "-m", message)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func main() {
	// Define a new command line flag for the commit message
	wordPtr := flag.String("m", "Default commit message", "A message for the commit")
	flag.Parse()

	// Get the current branch name
	branch, err := getCurrentBranch()
	if err != nil {
		fmt.Println("Error on getCurrentBranch():", err.Error())
		return
	}

	// Commit with the branch name and the provided message
	output, err := commitWithBranchName(branch, *wordPtr)
	if err != nil {
		fmt.Println("Error on commitWithBranchName():", err.Error())
		return
	}

	// Print the output of the commit command
	fmt.Println(output)
}
