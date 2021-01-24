package main

import (
	"github.com/spf13/pflag"
	"strconv"
	"strings"
)

func main() {
	var (
		showHelp        bool
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
	pflag.StringVarP(&inputFile, "in", "i", "data.life", "output file")
	pflag.IntVarP(&dimension, "dimension", "d", 3, "dimension of world")
	pflag.IntVarP(&size, "size", "S", 10, "side size")
	pflag.StringVarP(&B, "b-rule", "b", "5", "Rules for birth")
	pflag.StringVarP(&S, "s-rule", "s", "4;5", "Rules for save")
	pflag.IntVarP(&countGeneration, "count", "g", 100, "count generations.")
	pflag.BoolVarP(&showHelp, "help", "h", false,
		"Show help message")
	pflag.Parse()
	if showHelp {
		pflag.Usage()
		return
	}
	stringB := strings.Split(B, ";")
	for _, e := range stringB {
		elem, err := strconv.Atoi(e)
		if err != nil {
			break
		}
		b = append(b, elem)
	}
	stringS := strings.Split(S, ";")
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
	var d []int
	model.Setup(b, s, d)
	for i := 0; i < countGeneration; i++ {
		model.NextGeneration()
	}
}
