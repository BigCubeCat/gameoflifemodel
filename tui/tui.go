package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// RunTui run TUI
func RunTui() {
	SWITCH = false
	if err := tea.NewProgram(NewWindow()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
