package model

import (
	utils "github.com/bigcubecat/gameoflifemodel/utils"
)

// Life Game Of Life in N+1 dimansion torus
type Life struct {
	Data       []bool
	dataSize   int
	SIZE       int
	N          int
	B          map[int]bool
	S          map[int]bool
	steps      []int
	points     map[int][]int
	coords     []int
	Configured bool
}

// Setup setup model
func (life *Life) Setup(b []int, s []int, data []bool) {
	// create Data and memorize SIZE degres
	life.dataSize = utils.IntPow(life.SIZE, life.N)
	life.Data = data
	if life.Configured {
		return
	}
	for i := 0; i <= life.N; i++ {
		step := utils.IntPow(life.SIZE, i)
		life.steps = append(life.steps, step)
	} // Generate rules maps
	life.B = make(map[int]bool)
	life.S = make(map[int]bool)
	for _, i := range b {
		life.B[i] = true
	}
	for _, i := range s {
		life.S[i] = true
	}

	// find neighbors for not border point
	points := make(map[int]string, life.N)
	points[0] = ""
	life.coords = append(life.coords, 0)
	for _, s := range life.steps[0:life.N] {
		var newCoords []int
		for _, a := range life.coords {
			left := a - s
			right := a + s
			points[left] = points[a] + "L"
			points[right] = points[a] + "R"
			newCoords = append(newCoords, left)
			newCoords = append(newCoords, right)
		}
		for _, a := range life.coords {
			points[a] += "M"
		}
		life.coords = append(life.coords, newCoords...)
		newCoords = nil
	}
	// find all possible angles
	neighbors := make(map[string][]int, life.N+1)
	ch := make(chan string)
	go func() {
		defer close(ch)
		utils.Permutation("MLR", "", life.N, ch) // generate all angles
	}()
	for i := range ch {
		coords := make(map[string]int, len(life.coords))
		for k, v := range points {
			coords[v] = k
		}
		for index, char := range i {
			for key := range coords {
				if string(key[index]) == string(char) {
					diff := 0
					if string(char) == "L" {
						diff = 1
					} else if string(char) == "R" {
						diff = -1
					}
					coords[key] += diff * life.steps[index+1]
				}
			}
		}
		var list []int
		for _, value := range coords {
			list = append(list, value)
		}
		neighbors[i] = list
	}
	life.points = make(map[int][]int, life.dataSize)
	for i := range life.Data {
		life.points[i] = neighbors[life.checkBoreders(i)]
	}
	life.Configured = true
}

func (life *Life) checkBoreders(index int) string {
	var answer string
	for i := 0; i < life.N; i++ {
		if t := index % life.steps[i+1]; 0 <= t && t < life.steps[i] {
			// "LEFT" boreder by i degree
			answer += "L"
		} else if t := (index + life.steps[i]) % life.steps[i+1]; 0 <= t && t < life.steps[i] {
			// "RIGHT" border by i degree
			answer += "R"
		} else {
			answer += "M"
		}
	}
	return answer
}

func (life *Life) countNeighbors(index int) int {
	countN := 0
	for _, c := range life.points[index] {
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

// NextGeneration run next generation
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
