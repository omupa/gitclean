package interative

type model struct {
	branches         []string
	cursor           int
	selectedBranches map[string]bool
	screen           string
}
