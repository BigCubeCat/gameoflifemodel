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
	)
	pflag.BoolVarP(&finderMod, "find", "f", false, "find mod")
	pflag.StringVarP(&outputFile, "out", "o", "data.life", "output file")
	pflag.StringVarP(&inputFile, "in", "i", "", "input file")
	pflag.IntVarP(&dimension, "dimension", "d", 3, "dimension of world")
	pflag.IntVarP(&size, "size", "S", 128, "side size")
	pflag.StringVarP(&B, "b-rule", "b", "5", "Rules for birth")
	pflag.StringVarP(&S, "s-rule", "s", "4,5", "Rules for save")
	pflag.IntVarP(&countGeneration, "count", "g", 100, "count generations.")
	pflag.BoolVarP(&showHelp, "help", "h", false,
		"Show help message")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		return
	}
	if finderMod {
		findRules(size)
		return
	}
	model := Life{
		SIZE: size,
		N:    dimension,
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
		model.N = md.D
		model.SIZE = md.SIZE
	}
	stringB := strings.Split(B, ",")
	for _, e := range stringB {
		elem, err := strconv.Atoi(e)
		if err != nil {
			break
		}
		b = append(b, elem)
	}
	stringS := strings.Split(S, ",")
	for _, e := range stringS {
		elem, err := strconv.Atoi(e)
		if err != nil {
			break
		}
		s = append(s, elem)
	}

	model.Setup(b, s, d) // Set rules and data, if data exists
	fmt.Println(model.Data)
	fmt.Println("Model is created")
	for i := countGeneration; i > 0; i-- {
		model.NextGeneration()
	}
	out := RLECode(dataToString(model.GetData()))
	fmt.Println("Starting write to file")
	output := dataModel{
		D:    model.N,
		SIZE: model.SIZE,
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
