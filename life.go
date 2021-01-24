package main

import (
	"math/rand"
)

type Life struct {
	Data     []bool
	newData  []bool
	dataSize int
	SIZE     int
	N        int
	B        map[int]bool
	S        map[int]bool
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
func (life *Life) Setup(b []int, s []int, data []int) {
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
		for i := 0; i < life.dataSize; i++ {
			life.Data[i] = false
			life.newData[i] = false
		}
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
		return life.B[life.countNeighbours(index)]
	}
	return life.S[life.countNeighbours(index)]
}

func (life *Life) countNeighbours(index int) int {
	coords := []int{index}
	countN := 0
	for i := 0; i < life.N; i++ {
		step := intPow(life.SIZE, i)
		newCoords := []int{index}
		for _, a := range coords {
			left := a - step
			right := a + step
			if life.inWorld(left) {
				if life.getCell(left) {
					countN++
				}
				newCoords = append(newCoords, left)
			}
			if life.inWorld(right) {
				if life.getCell(right) {
					countN++
				}
				newCoords = append(newCoords, right)
			}
		}
		coords = append(coords, newCoords...)
	}
	return countN
}

func (life *Life) NextGeneration() {
	for i := range life.Data {
		life.newData[i] = life.applyRules(i)
	}
	life.Data, life.newData = life.newData, life.Data
}
