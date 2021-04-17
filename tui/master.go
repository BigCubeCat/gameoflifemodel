package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredButtonStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle            = lipgloss.NewStyle()

	focusedSubmitButton = "[ " + focusedStyle.Render("Submit") + " ]"
	blurredSubmitButton = "[ " + blurredButtonStyle.Render("Submit") + " ]"
)

type MainModel struct {
	start tea.Model
	show  tea.Model
}

func (m MainModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if SELECTED == 0 {
		m.start, cmd = m.start.Update(msg)
	} else {
		m.show, cmd = m.show.Update(msg)
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m MainModel) View() string {
	if SELECTED == 0 {
		return m.start.View()
	}
	return m.show.View()
}

func NewWindow() MainModel {
	return MainModel{NewStartModel(), nil}
}
