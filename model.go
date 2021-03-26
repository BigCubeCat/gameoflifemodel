package main

import "strconv"

type MODEL interface {
	NextGeneration()
	Setup(a []int, b []int, c []bool)
	GetData() []bool
	getN() int
	getSIZE() int
	setN(n int)
	setSIZE(n int)
	GetB() []string
	GetS() []string
}

func ListKeys(m map[int]bool) []string {
	var answer []string
	for k := range m {
		answer = append(answer, strconv.Itoa(k))
	}
	return answer
}
