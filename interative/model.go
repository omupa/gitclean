package interative

type model struct {
	screen   string
	force    bool
	cursor   int
	branches []string
	// The key will be the branch name and the value indicates whether that branch has been deleted
	selectedBranches map[string]bool
}

func hasNotBranchesToDelete(m model) bool {
	return m.branches == nil || len(m.branches) == 0
}

func emptyModel(force bool) model {
	allBranches := make([]string, 0)
	return loadBranches(allBranches, force)
}

func loadBranches(allBranches []string, force bool) model {
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
