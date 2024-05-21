package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "branch")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	branches := strings.Split(string(stdout), "\n")
	deleteBranches(branches)
}

func deleteBranches(branches []string) {
	for _, branch := range branches {

		if strings.Contains(branch, "* ") || branch == "" {
			continue
		}

		cmd := exec.Command("git", "branch", "-d", strings.TrimSpace(branch))
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Print(string(stdout))
	}
}
