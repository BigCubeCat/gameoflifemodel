package tui

import (
	"fmt"
	"time"

	"github.com/TwinProduction/go-color"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	index        int
	dimension    int
	size         int
	alivePersent int
	bRule        string
	sRule        string
	GProgress    *progress.Model
	AProgress    *progress.Model
	G            float64
	A            float64
	chanel       chan ChangeModel
	viewData     string
}

type ChangeModel struct {
	A        float64
	G        float64
	Finished bool
}

var MM MainModel

type tickMsg time.Time

func (m MainModel) Init() tea.Cmd {
	return tickCmd()
}
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/30, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case tickMsg:
		dat := <-m.chanel
		if dat.Finished {
			return m, tea.Quit
		}
		header := fmt.Sprintf(color.Ize(color.Blue, "Dimension: %d; \n"), m.dimension)
		header += fmt.Sprintf(color.Ize(color.Purple, "Side size: %d;\n"), m.size)
		header += fmt.Sprintf(color.Ize(color.Green, "B-rule = %s\n"), m.bRule)
		header += fmt.Sprintf(color.Ize(color.Yellow, "S-rule = %s\n"), m.sRule)
		footer := "\nAttempts: \n"
		footer += m.AProgress.View(dat.A) + "\n\n"
		footer += "Generations: \n"
		footer += m.GProgress.View(dat.G) + "\n\n"
		m.viewData = header + "\n" + footer
		return m, tickCmd()
	}
	return m, nil
}

func (m MainModel) View() string {
	return m.viewData
}

func NewModel(c chan ChangeModel, b string, s string, dim int, size int) MainModel {
	p, _ := progress.NewModel(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
	a, _ := progress.NewModel(progress.WithScaledGradient("#FDFF8C", "#FF7CCB"))
	return MainModel{GProgress: p, AProgress: a, chanel: c, dimension: dim, size: size, bRule: b, sRule: s}
}
