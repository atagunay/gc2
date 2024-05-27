package main

import (
	"flag"
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

func commitWithBranchName(branchName string, commitMessage string) (string, error) {
	message := branchName + " " + commitMessage

	cmd := exec.Command("git", "commit", "-m", message)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func main() {

	wordPtr := flag.String("m", "Default commit message", "A message for the commit")
	flag.Parse()

	branch, err := getCurrentBranch()
	if err != nil {
		fmt.Println("Error on getCurrentBranch():", err.Error())
		return
	}

	output, err := commitWithBranchName(branch, *wordPtr)
	if err != nil {
		fmt.Println("Error on commitWithBranchName():", err.Error())
		return
	}

	fmt.Println(output)
}
