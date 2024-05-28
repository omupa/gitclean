package interative

import "fmt"

func (m model) View() string {
	if m.screen == "select-branches" {
		return selectBranchesScreen(m)
	}

	return branchesDeleted(m)
}

func selectBranchesScreen(m model) string {
	if len(m.branches) == 0 {
		return "No branches can be deleted!"
	}

	s := "Which branch you want to delete?\n\n"

	for i, branch := range m.branches {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selectedBranches[branch]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, branch)
	}

	s += "\nPress SPACE to select/unselect"
	s += "\nPress ENTER to delete selected branches"
	s += "\nPress Q to quit.\n"

	return s
}

func branchesDeleted(m model) string {
	s := "The following branches was deleted:\n\n"

	for key, value := range m.selectedBranches {
		deleted := "[Yes]"
		if !value {
			deleted = "[Not]"
		}

		s += deleted + " " + key + "\n"
	}

	s += "\n"

	return s
}
