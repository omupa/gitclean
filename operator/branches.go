package operator

import (
	"fmt"
	"os/exec"
	"strings"
)

func FetchBranches() []string {
	cmd := exec.Command("git", "branch")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	allBranches := strings.Split(string(stdout), "\n")
	branchesCanBeDeleted := []string{}

	for _, branch := range allBranches {

		if strings.Contains(branch, "* ") || branch == "" {
			continue
		}

		branchesCanBeDeleted = append(branchesCanBeDeleted, strings.TrimSpace(branch))
	}

	return branchesCanBeDeleted
}

func DeleteBranch(branchToBeDeleted string, force bool) error {
	if branchToBeDeleted == "* " || branchToBeDeleted == "" {
		return nil
	}

	deleteFlag := "-d"
	if force {
		deleteFlag = "-D"
	}

	cmd := exec.Command("git", "branch", deleteFlag, strings.TrimSpace(branchToBeDeleted))
	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

func DeleteAll(force bool) {
	branches := FetchBranches()

	for _, branch := range branches {
		DeleteBranch(branch, force)
	}
}
