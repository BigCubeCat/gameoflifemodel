package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	index        int
	p            []int
	dimension    int
	size         int
	alivePersent int
	bRule        string
	sRule        string
	GProgress    *progress.Model
	APrgogres    *progress.Model
	G            float64
	A            float64
}

func (m MainModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *MainModel) UpdateData(G float64, A float64) {
	m.A = A
	m.G = G
}

func (m MainModel) View() string {
	header := fmt.Sprintf("Dimension: %d; ", m.dimension)
	header += fmt.Sprintf("Side size: %d;\n", m.size)
	header += fmt.Sprintf("B-rule = %s\n", m.bRule)
	header += fmt.Sprintf("S-rule = %s\n\n\n", m.sRule)
	canvas := []string{}
	for i := 0; i < 10; i++ {
		canvas = append(canvas, "")
		for _, e := range m.p {
			if int(e) > i {
				canvas[i] += " "
			} else {
				canvas[i] += "*"
			}
		}
	}
	top := strings.Join(canvas, "\n") + "\n\n\n"
	footer := ""
	footer += "Attempt:\n" + m.APrgogres.View(m.A) + "\n"
	footer += "Generation:\n" + m.GProgress.View(m.G) + "\n\n"
	return header + top + footer
}

func NewModel() MainModel {
	p, _ := progress.NewModel(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	a, _ := progress.NewModel(progress.WithScaledGradient("#FDFF8C", "#FF7CCB"))
	return MainModel{GProgress: p, APrgogres: a}
}
