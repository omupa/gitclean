package interative

type model struct {
	screen           string
	force            bool
	cursor           int
	branches         []string
	selectedBranches map[string]bool
}
