package model

import (
	utils "github.com/bigcubecat/gameoflifemodel/utils"
)

// Life Game Of Life in N+1 dimansion torus
type Life struct {
	Data      []bool
	dataSize  int
	SIZE      int
	N         int
	B         map[int]bool
	S         map[int]bool
	steps     []int
	neighbors map[string][]int
	coords    []int
}

// Setup setup model
func (life *Life) Setup(b []int, s []int, data []bool) {
	// Generate rules maps
	life.B = make(map[int]bool)
	life.S = make(map[int]bool)
	for _, i := range b {
		life.B[i] = true
	}
	for _, i := range s {
		life.S[i] = true
	}
	life.dataSize = utils.IntPow(life.SIZE, life.N)
	life.Data = data
	for i := 0; i <= life.N; i++ {
		step := utils.IntPow(life.SIZE, i)
		life.steps = append(life.steps, step)
	}
	life.neighbors = make(map[string][]int, 0)
	ch := make(chan string)
	go func() {
		defer close(ch)
		utils.Permutation("MLR", "", life.N, ch) // generate all angles
	}()
	for i := range ch {
		life.neighbors[i] = []int{} // TODO
	}

	life.coords = append(life.coords, 0)
	for _, s := range life.steps[0:life.N] {
		var newCoords []int
		for _, a := range life.coords {
			newCoords = append(newCoords, a-s)
			newCoords = append(newCoords, a+s)
		}
		life.coords = append(life.coords, newCoords...)
		newCoords = nil
	}
}

func (life *Life) checkBoreders(index int) (string, bool) {
	var answer string
	border := false
	for i := 0; i < life.N; i++ {
		if t := index % life.steps[i+1]; 0 <= t && t < life.steps[i] {
			// "LEFT" boreder by i degree
			answer += "L"
			border = true
		} else if t := (index + life.steps[i]) % life.steps[i+1]; 0 <= t && t < life.steps[i] {
			// "RIGHT" border by i degree
			answer += "R"
			border = true
		} else {
			answer += "M"
		}
	}
	return answer, border // +1
}

func (life *Life) countNeighbors(index int) int {
	countN := 0
	bs, border := life.checkBoreders(index)
	coords := make([]int, len(life.coords))
	copy(coords, life.coords[0:len(life.coords)])
	if border {
		for i, b := range bs {
			if b != 0 {
				for _, elem := range life.coords {
					if elem == 0 {
						continue
					}
					if life.neighbors[elem][i] != 0 {
						coords[i] += -int(b) * life.steps[i+1]
					}
				}
			}

		}
	}
	for _, c := range coords {
		if life.Data[index+c] {
			countN++
		}
	}
	return countN
}

func (life *Life) applyRules(index int) bool {
	if life.Data[index] {
		return life.S[life.countNeighbors(index)]
	}
	return life.B[life.countNeighbors(index)]
}

func (life *Life) NextGeneration() {
	newData := make([]bool, life.dataSize)
	for i := 0; i < life.dataSize; i++ {
		newData[i] = life.applyRules(i)
	}
	life.Data, newData = newData, life.Data
}

// GetData return data
func (life *Life) GetData() []bool {
	return life.Data
}

// GetB return rule for birth in string
func (life *Life) GetB() string {
	return utils.ListKeys(life.B)
}

// GetS return rule for save in string
func (life *Life) GetS() string {
	return utils.ListKeys(life.S)
}

// GetN return dimension
func (life *Life) GetN() int {
	return life.N
}

// GetSIZE returns side size
func (life *Life) GetSIZE() int {
	return life.SIZE
}

// SetN : set dimension (need re-setup)
func (life *Life) SetN(n int) {
	life.N = n
}

// SetSIZE : set size (need re-setup)
func (life *Life) SetSIZE(n int) {
	life.SIZE = n
}
