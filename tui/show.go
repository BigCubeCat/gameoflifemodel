package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ShowModel struct {
	index int
}

func (m ShowModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ShowModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m ShowModel) View() string {
	return ""
}
