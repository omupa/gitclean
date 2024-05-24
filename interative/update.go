package interative

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.branches)-1 {
				m.cursor++
			}

		case " ":
			_, ok := m.selectedBranches[m.branches[m.cursor]]
			if ok {
				delete(m.selectedBranches, m.branches[m.cursor])
			} else {
				m.selectedBranches[m.branches[m.cursor]] = false
			}

		case "enter":
			deleteBranchesCmd(m)
			m.screen = "branches-deleted"
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
