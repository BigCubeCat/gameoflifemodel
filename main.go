package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func dataToString(data []bool) string {
	answer := ""
	for _, e := range data {
		if e {
			answer += "A"
		} else {
			answer += "D"
		}
	}
	return answer
}

func saveToFile(model Life, data string, countGeneration int, fileName string) error {
	newData := "D: " + strconv.Itoa(model.N) + ";\nSize: " +
		strconv.Itoa(model.SIZE) + ";\n" + data +
		"Generation: " + strconv.Itoa(countGeneration) + ";\n"
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, e := f.WriteString(newData)
	return e
}

func main() {
	var (
		showHelp        bool
		last            int
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
	pflag.StringVarP(&outputFile, "out", "o", "data.life", "output file")
	pflag.StringVarP(&inputFile, "in", "i", "", "output file")
	pflag.IntVarP(&dimension, "dimension", "d", 3, "dimension of world")
	pflag.IntVarP(&size, "size", "S", 10, "side size")
	pflag.StringVarP(&B, "b-rule", "b", "5", "Rules for birth")
	pflag.StringVarP(&S, "s-rule", "s", "4,5", "Rules for save")
	pflag.IntVarP(&countGeneration, "count", "g", 100, "count generations.")
	pflag.IntVarP(&last, "last", "l", 1, "write to file last {value} generations")
	pflag.BoolVarP(&showHelp, "help", "h", false,
		"Show help message")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		return
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
	model := Life{
		SIZE: size,
		N:    dimension,
	}
	var d []bool
	if len(inputFile) != 0 {
		content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			panic(err)
		}
		contentStrings := strings.Split(string(content), ";")
		str := RLEDecode(contentStrings[2])
		newN, _ := strconv.Atoi(strings.Split(contentStrings[0], ": ")[1])
		newSize, _ := strconv.Atoi(strings.Split(contentStrings[1], ": ")[1])

		model.N = newN
		model.SIZE = newSize
		for _, c := range str {
			if string(c) == "A" {
				d = append(d, true)
			} else {
				d = append(d, false)
			}
		}

	}
	lastGens := ""       // last generations
	model.Setup(b, s, d) // Set rules and data, if data exists
	fmt.Println("Model is created")
	for i := countGeneration; i > 0; i-- {
		model.NextGeneration()
		if i <= last {
			lastGens += RLECode(dataToString(model.GetData())) + ";\n"
		}
	}
	fmt.Println("Starting write to file")
	saveErr := saveToFile(model, lastGens, countGeneration, outputFile)
	if saveErr != nil {
		panic(saveErr)
	} else {
		fmt.Println("Finish write to file. No errors")
	}
}
