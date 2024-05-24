package interative

import tea "github.com/charmbracelet/bubbletea"

func InitialModel() model {
	allBranches := fetchBranches()
	allSelectedBranches := make(map[string]bool, len(allBranches))

	for _, branch := range allBranches {
		allSelectedBranches[branch] = false
	}

	m := model{
		branches:         allBranches,
		cursor:           0,
		selectedBranches: allSelectedBranches,
		screen:           "select-branches",
	}

	return m
}

func (m model) Init() tea.Cmd {
	return fetchBranchesCmd()
}
