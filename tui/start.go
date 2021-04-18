package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SubmitSignal tea.Cmd

type StartWindow struct {
	index  int
	Fields []textinput.Model
	submit string
}

func NewStartModel() StartWindow {
	Fields := []textinput.Model{}
	for i := 0; i < 8; i++ {
		Fields = append(Fields, textinput.NewModel())
	}
	placeholders := []string{
		"dimension", "side size", "rule for birth", "rule for save", "probability", "count attempts", "count generations", "output database",
	}
	for i := range Fields {
		Fields[i].Placeholder = placeholders[i]
	}
	Fields[0].Focus()
	Fields[0].PromptStyle = focusedStyle
	Fields[0].TextStyle = focusedStyle
	return StartWindow{0, Fields, blurredSubmitButton}

}
func (m StartWindow) Init() tea.Cmd {
	return textinput.Blink
}

func (m StartWindow) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// Cycle between m.Fields
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" && m.index == len(m.Fields) {
				SWITCH = true
				SetSelected(1)
				ReadInput(m.Fields[0].Value(), m.Fields[1].Value(), m.Fields[2].Value(), m.Fields[3].Value(), m.Fields[4].Value(), m.Fields[5].Value(), m.Fields[6].Value(), m.Fields[7].Value())
				return m, nil
			}
			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.index--
			} else {
				m.index++
			}
			if m.index > len(m.Fields) {
				m.index = 0
			} else if m.index < 0 {
				m.index = len(m.Fields)
			}
			for i := 0; i <= len(m.Fields)-1; i++ {
				if i == m.index {
					// Set focused state
					m.Fields[i].Focus()
					m.Fields[i].PromptStyle = focusedStyle
					m.Fields[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.Fields[i].Blur()
				m.Fields[i].PromptStyle = noStyle
				m.Fields[i].TextStyle = noStyle
			}

			if m.index == len(m.Fields) {
				m.submit = focusedSubmitButton
			} else {
				m.submit = blurredSubmitButton
			}

			return m, nil
		}
	}

	// Handle character input and blinks
	m, cmd = updateFields(msg, m)
	return m, cmd
}

func updateFields(msg tea.Msg, m StartWindow) (StartWindow, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	for i := range m.Fields {
		m.Fields[i], cmd = m.Fields[i].Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m StartWindow) View() string {
	s := "\n"
	for i := 0; i < len(m.Fields); i++ {
		s += m.Fields[i].View()
		if i < len(m.Fields)-1 {
			s += "\n"
		}
	}

	s += "\n\n" + m.submit + "\n"
	return s
}
