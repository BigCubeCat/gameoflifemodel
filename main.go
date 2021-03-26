package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
)

type dataModel struct {
	D    int
	SIZE int
	DATA string
	B    string
	S    string
}

func random() bool {
	return rand.Uint64()&(1<<63) == 0
}

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

func main() {
	var (
		showHelp        bool
		finderMod       bool
		outputFile      string
		inputFile       string
		dimension       int
		size            int
		B               string
		S               string
		countGeneration int
		b               []int
		s               []int
		model3d         bool
		details         bool
	)
	pflag.BoolVarP(&finderMod, "find", "f", false, "find mod")
	pflag.StringVarP(&outputFile, "out", "o", "data.life", "output file")
	pflag.StringVarP(&inputFile, "in", "i", "", "input file")
	pflag.IntVarP(&dimension, "dimension", "d", 3, "dimension of world")
	pflag.IntVarP(&size, "size", "S", 128, "side size")
	pflag.StringVarP(&B, "b-rule", "b", "5", "Rules for birth")
	pflag.StringVarP(&S, "s-rule", "s", "4,5", "Rules for save")
	pflag.IntVarP(&countGeneration, "count", "g", 100, "count generations.")
	pflag.BoolVarP(&showHelp, "help", "h", false, "Show help message")
	pflag.BoolVarP(&model3d, "model3d", "3", false, "Use 3D model")
	pflag.BoolVarP(&details, "details", "e", false, "save all generations")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		fmt.Println("Use \",\" to split different numbers on rule.")
		fmt.Println("Use \"{start}.{end}\" to set range [start, end] (end and start includes)")
		return
	}

	var model MODEL

	if model3d {
		model = &Life3d{
			SIZE: size,
			N:    3,
		}
		dimension = 3
	} else {
		model = &Life{
			SIZE: size,
			N:    dimension,
		}
	}
	var d []bool
	if inputFile == "" {
		d = make([]bool, intPow(size, dimension))
		for i := range d {
			d[i] = random() && random()
		}
	} else {
		byteData, err := ioutil.ReadFile(inputFile)
		if err != nil {
			panic(err)
		}
		var md dataModel
		json.Unmarshal(byteData, &md)
		B = md.B
		S = md.S
		model.setN(md.D)
		model.setSIZE(md.SIZE)
		str_data := RLEDecode(md.DATA)
		for _, c := range str_data {
			if string(c) == "A" {
				d = append(d, true)
			} else {
				d = append(d, false)
			}
		}
	}
	b = readRule(B)
	s = readRule(S)

	model.Setup(b, s, d) // Set rules and data, if data exists
	fmt.Println("HERE")
	fmt.Println(model.GetData())
	fmt.Println("Model is created")
	if finderMod {
		for t := 0; t < 100; t++ {
			findRules(model, countGeneration, details, t)
		}
		return
	}
	for i := countGeneration; i > 0; i-- {
		model.NextGeneration()
	}
	out := RLECode(dataToString(model.GetData()))
	fmt.Println("Starting write to file")
	output := dataModel{
		D:    model.getN(),
		SIZE: model.getSIZE(),
		DATA: out,
		B:    B,
		S:    S,
	}
	file, _ := json.MarshalIndent(output, "", "")
	saveErr := ioutil.WriteFile(outputFile, file, 0644)
	if saveErr != nil {
		panic(saveErr)
	} else {
		fmt.Println("Finish write to file. No errors")
	}
}
