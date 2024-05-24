package interative

import (
	"fmt"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func fetchBranches() []string {
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

func fetchBranchesCmd() tea.Cmd {
	return func() tea.Msg {
		return fetchBranches()
	}
}

func deleteBranchesCmd(m model) tea.Cmd {
	for branchToBeDeleted := range m.selectedBranches {
		deleteBranch(m, branchToBeDeleted)
	}

	return func() tea.Msg { return nil }
}

func deleteBranch(m model, branchToBeDeleted string) {
	if branchToBeDeleted == "* " || branchToBeDeleted == "" {
		return
	}

	cmd := exec.Command("git", "branch", "-d", strings.TrimSpace(branchToBeDeleted))
	_, err := cmd.Output()

	if err != nil {
		m.selectedBranches[branchToBeDeleted] = false
	} else {
		m.selectedBranches[branchToBeDeleted] = true
	}
}
