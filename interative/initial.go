package interative

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/omupa/gitclean/operator"
)

func (m model) Init() tea.Cmd {
	return nil
}

func StartInterative(force bool) {
	p := tea.NewProgram(initialModel(force))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel(force bool) model {
	allBranches := operator.FetchBranches()
	allSelectedBranches := make(map[string]bool, len(allBranches))

	for _, branch := range allBranches {
		allSelectedBranches[branch] = false
	}

	m := model{
		screen:           SelectBranches,
		force:            force,
		cursor:           0,
		branches:         allBranches,
		selectedBranches: allSelectedBranches,
	}

	return m
}
