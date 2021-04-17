package tui

import "strconv"

var SELECTED int
var DM DataModel

type DataModel struct {
	D    int
	SIZE int
	B    string
	S    string
	P    int
	G    int
	O    string
}

func SetSelected(i int) {
	SELECTED = i
}
func ReadInput(d string, S string, b string, s string, p string, a string, g string, o string) {
	D, e := strconv.Atoi(d)
	if e != nil {
		return
	}
	size, e := strconv.Atoi(S)
	if e != nil {
		return
	}
	P, e := strconv.Atoi(p)
	if e != nil {
		return
	}
	G, e := strconv.Atoi(g)
	DM = DataModel{D: D, SIZE: size, B: b, S: s, P: P, G: G, O: o}
}
