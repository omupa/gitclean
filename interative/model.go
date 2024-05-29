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
