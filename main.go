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

func saveToFile(model Life, countGeneration int, fileName string) error {
	newData := "D: " + strconv.Itoa(model.N) + ";\nSize: " +
		strconv.Itoa(model.SIZE) + ";\n" + RLECode(dataToString(model.GetData())) +
		";\nGeneration: " + strconv.Itoa(countGeneration) + ";\n"
	f, err := os.Create(fileName)
	defer f.Close()
	f.WriteString(newData)
	return err
}

func main() {
	var (
		showHelp bool

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
		// TODO
		content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			panic(err)
		}
		str := RLEDecode(strings.Split(string(content), ";")[2])
		for _, c := range str {
			if string(c) == "A" {
				d = append(d, true)
			} else {
				d = append(d, false)
			}
		}

	}
	model.Setup(b, s, d)
	fmt.Println("Model is created")
	for i := 0; i < countGeneration; i++ {
		model.NextGeneration()
	}
	fmt.Println("Starting write to file")
	saveErr := saveToFile(model, countGeneration, outputFile)
	if saveErr != nil {
		panic(saveErr)
	} else {
		fmt.Println("Finish write to file. No errors")
	}
}
