package interative

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/omupa/gitclean/operator"
)

func fetchBranchesCmd() tea.Cmd {
	return func() tea.Msg {
		return operator.FetchBranches()
	}
}

func deleteBranchesCmd(m model) tea.Cmd {
	for branchToBeDeleted := range m.selectedBranches {
		deleteBranch(m, branchToBeDeleted)
	}

	return func() tea.Msg { return nil }
}

func deleteBranch(m model, branchToBeDeleted string) {
	err := operator.DeleteBranch(branchToBeDeleted, false)

	if err != nil {
		m.selectedBranches[branchToBeDeleted] = false
	} else {
		m.selectedBranches[branchToBeDeleted] = true
	}
}
