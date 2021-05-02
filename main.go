package main

import (
	"fmt"
	"os"

	"github.com/TwinProduction/go-color"
	finder "github.com/bigcubecat/gameoflifemodel/finder"
	lifeModel "github.com/bigcubecat/gameoflifemodel/model"
	"github.com/bigcubecat/gameoflifemodel/tui"
	"github.com/bigcubecat/gameoflifemodel/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/pflag"
)

// RunProgram run finder
func RunProgram(m lifeModel.MODEL) {
	var (
		showHelp        bool
		dimension       int
		size            int
		attempts        int
		B               string
		S               string
		countGeneration int
		b               []int
		s               []int
		model3d         bool
		model4d         bool
		model5d         bool
		probability     int
		fileName        string
		f               bool
		inputFile       string
	)
	pflag.IntVarP(&dimension, "dimension", "d", 3, "dimension of world")
	pflag.IntVarP(&size, "size", "S", 128, "side size")
	pflag.StringVarP(&B, "b-rule", "b", "5", "Rules for birth")
	pflag.StringVarP(&S, "s-rule", "s", "4,5", "Rules for save")
	pflag.IntVarP(&countGeneration, "count", "g", 100, "count generations.")
	pflag.BoolVarP(&showHelp, "help", "h", false, "Show help message")
	pflag.BoolVarP(&model3d, "model3d", "3", false, "Use 3D model")
	pflag.BoolVarP(&model4d, "model4d", "4", false, "Use 4D model")
	pflag.BoolVarP(&model5d, "model5d", "5", false, "Use 5D model")
	pflag.IntVarP(&attempts, "attempt", "a", 100, "Count attempts")
	pflag.IntVarP(&probability, "probability", "p", 50, "probability in %")
	pflag.StringVarP(&fileName, "out", "o", "output.db", "Database name")
	pflag.StringVarP(&inputFile, "input", "i", "", "Input file")
	pflag.BoolVarP(&f, "use-file", "u", false, "Use file")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		fmt.Println("Use \",\" to split different numbers on rule.")
		fmt.Println("Use \"{start}.{end}\" to set range [start, end] (end and start includes)")
		return
	}

	var d []bool

	if f {
		dataModel, err := finder.ReadJSON(attempts, inputFile)
		if err != nil {
			fmt.Println(color.Ize(color.Red, "Error"))
			fmt.Println(err)
			return
		}
		B = dataModel.B
		S = dataModel.S
		dimension = dataModel.D
		size = dataModel.SIZE
		d = utils.StringToData(utils.RLEDecode(dataModel.DATA))
	}
	var model lifeModel.MODEL

	if m == nil {
		if model3d {
			model = &lifeModel.Life3d{
				SIZE: size,
				N:    3,
			}
			dimension = 3
		} else if model4d {
			model = &lifeModel.Life4d{
				SIZE: size,
				N:    4,
			}
			dimension = 4
		} else if model5d {
			model = &lifeModel.Life5d{
				SIZE: size,
				N:    5,
			}
			dimension = 5
		} else {
			model = &lifeModel.Life{
				SIZE: size,
				N:    dimension,
			}
		}
	} else {
		model = m
	}
	dataSize := utils.IntPow(size, dimension)
	b = utils.ReadRule(B)
	s = utils.ReadRule(S)

	fmt.Println(color.Ize(color.Green, "Start game of life"))
	if f {
		model.Setup(b, s, d)
		finder.RunFromFile(model, attempts, fileName)
		return
	}
	chanel := make(chan tui.ChangeModel, 1)
	go func() {
		defer close(chanel)
		finder.Run(model, countGeneration, attempts, fileName, probability, b, s, dataSize, chanel)
	}()
	p := tea.NewProgram(tui.NewModel(chanel, B, S, dimension, size))
	p.EnterAltScreen()
	if err := p.Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(color.Ize(color.Green, "Finish. No Errors"))
}

func main() {
	RunProgram(nil)
}
