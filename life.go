package main

import (
	"math/rand"
)

type Life struct {
	Data            []bool
	newData         []bool
	dataSize        int
	SIZE            int
	N               int
	B               map[int]bool
	S               map[int]bool
	neighborsCoords []int
}

func randBool() bool {
	return rand.Float32() < 0.5
}

func intPow(a int, b int) int {
	answer := 1
	for i := 1; i <= b; i++ {
		answer *= a
	}
	return answer
}
func (life *Life) Setup(b []int, s []int, data []bool) {
	life.B = make(map[int]bool)
	life.S = make(map[int]bool)
	for _, i := range b {
		life.B[i] = true
	}
	for _, i := range s {
		life.S[i] = true
	}
	life.dataSize = intPow(life.SIZE, life.N)
	life.Data = make([]bool, life.dataSize)
	life.newData = make([]bool, life.dataSize)
	if len(data) == 0 {
		for i := 0; i < life.dataSize; i++ {
			life.Data[i] = randBool()
			life.newData[i] = false
		}
	} else {
		life.Data = data
		life.newData = data
	}
	life.neighborsCoords = []int{0}
	for i := 0; i < life.N; i++ {
		step := intPow(life.SIZE, i)
		var newCoords []int
		for _, a := range life.neighborsCoords {
			left := a - step
			right := a + step
			newCoords = append(newCoords, left)
			newCoords = append(newCoords, right)
		}
		life.neighborsCoords = append(life.neighborsCoords, newCoords...)
	}
}

func (life *Life) inWorld(index int) bool {
	return index < len(life.Data) && index >= 0
}

func (life *Life) getCell(index int) bool {
	if !life.inWorld(index) {
		return false
	}
	return life.Data[index]
}

func (life *Life) applyRules(index int) bool {
	cell := life.getCell(index)
	if cell {
		return life.S[life.countNeighbours(index)]
	}
	return life.B[life.countNeighbours(index)]
}

func (life *Life) countNeighbours(index int) int {
	countN := 0
	for _, indx := range life.neighborsCoords {
		if indx == 0 {
			continue
		}
		coord := index + indx
		if life.getCell(coord) {
			countN++
		}
	}
	return countN
}

func (life *Life) NextGeneration() {
	for i := range life.Data {
		life.newData[i] = life.applyRules(i)
	}
	life.Data, life.newData = life.newData, life.Data
}

func (life *Life) GetData() []bool {
	return life.Data
}
