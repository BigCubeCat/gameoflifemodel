package model

import "strings"
import "strconv"

type MODEL interface {
	NextGeneration()
	Setup(a []int, b []int, c []bool)
	GetData() []bool
	GetN() int
	GetSIZE() int
	SetN(n int)
	SetSIZE(n int)
	GetB() string
	GetS() string
}

type dataModel struct {
	D    int
	SIZE int
	DATA string
	B    string
	S    string
}

func ListKeys(m map[int]bool) string {
	var answer []string
	for k := range m {
		answer = append(answer, strconv.Itoa(k))
	}
	return strings.Join(answer, ",")
}
