package main

import (
	"fmt"
	"github.com/TwinProduction/go-color"
	finder "github.com/bigcubecat/gameoflifemodel/finder"
	lifeModel "github.com/bigcubecat/gameoflifemodel/model"
	"github.com/spf13/pflag"
	"strconv"
	"strings"
)

func readRule(rule string) []int {
	var answer []int
	r := strings.Split(rule, ",")
	for _, e := range r {
		elem, err := strconv.Atoi(e)
		if err != nil {
			if strings.Contains(e, ".") {
				ran := strings.Split(e, ".")
				var subrange []int
				start, err1 := strconv.Atoi(ran[0])
				fin, err2 := strconv.Atoi(ran[1])
				if err1 == nil && err2 == nil {
					for i := start; i <= fin; i++ {
						subrange = append(subrange, i)
					}
				}
				answer = append(answer, subrange...)
			} else {
				break
			}
		} else {
			answer = append(answer, elem)
		}
	}
	return answer
}

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
	)
	pflag.IntVarP(&dimension, "dimension", "d", 3, "dimension of world")
	pflag.IntVarP(&size, "size", "S", 128, "side size")
	pflag.StringVarP(&B, "b-rule", "b", "5", "Rules for birth")
	pflag.StringVarP(&S, "s-rule", "s", "4,5", "Rules for save")
	pflag.IntVarP(&countGeneration, "count", "g", 100, "count generations.")
	pflag.BoolVarP(&showHelp, "help", "h", false, "Show help message")
	pflag.BoolVarP(&model3d, "model3d", "3", false, "Use 3D model")
	pflag.BoolVarP(&model3d, "model4d", "4", false, "Use 4D model")
	pflag.BoolVarP(&model3d, "model5d", "5", false, "Use 5D model")
	pflag.IntVarP(&attempts, "attempt", "a", 100, "Count attempts")
	pflag.IntVarP(&probability, "probability", "p", 50, "probability in %")
	pflag.StringVarP(&fileName, "out", "o", "output.db", "Database name")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		fmt.Println("Use \",\" to split different numbers on rule.")
		fmt.Println("Use \"{start}.{end}\" to set range [start, end] (end and start includes)")
		return
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
		} else if model5d {
			model = &lifeModel.Life5d{
				SIZE: size,
				N:    5,
			}
		} else {
			model = &lifeModel.Life{
				SIZE: size,
				N:    dimension,
			}
		}
	} else {
		model = m
	}
	dataSize := lifeModel.IntPow(size, dimension)
	b = readRule(B)
	s = readRule(S)

	fmt.Println(color.Ize(color.Green, "Start game of life"))
	finder.FindRules(model, countGeneration, attempts, fileName, probability, b, s, dataSize)
	fmt.Println(color.Ize(color.Green, "Finish. No Errors"))
}

func main() {
	RunProgram(nil)
}
