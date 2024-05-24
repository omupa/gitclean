package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	interative "github.com/omupa/gitclean/interative"
)

func main() {
	p := tea.NewProgram(interative.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
