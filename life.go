package main

type Life struct {
	Data     []bool
	dataSize int
	SIZE     int
	N        int
	B        map[int]bool
	S        map[int]bool
	steps    []int
}

func intPow(a int, b int) int {
	answer := 1
	for i := 1; i <= b; i++ {
		answer *= a
	}
	return answer
}

func (life *Life) getN() int {
	return life.N
}

func (life *Life) getSIZE() int {
	return life.SIZE
}

func (life *Life) setN(n int) {
	life.N = n
}

func (life *Life) setSIZE(n int) {
	life.SIZE = n
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
	life.Data = data
	for i := 0; i < life.N; i++ {
		step := intPow(life.SIZE, i)
		life.steps = append(life.steps, step)
	}
}

func (life *Life) inWorld(index int) bool {
	return index < life.dataSize && index >= 0
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
	coords := []int{0}
	for _, s := range life.steps {
		var newCoords []int
		for _, a := range coords {
			left := a - s
			if life.inWorld(index + left) {
				newCoords = append(newCoords, left)
				if life.getCell(index + left) {
					countN += 1
				}
			}
			if (index+s)%life.SIZE == 0 { // if it's corner on current dimension, don't calculate right boreder
				continue
			}
			right := a + s
			if life.inWorld(index + right) {
				newCoords = append(newCoords, right)
				if life.getCell(index + right) {
					countN += 1
				}
			}
		}
		coords = append(coords, newCoords...)
		newCoords = nil
	}
	return countN
}

func (life *Life) NextGeneration() {
	newData := make([]bool, life.dataSize)
	for i := 0; i < life.dataSize; i++ {
		newData[i] = life.applyRules(i)
	}
	life.Data, newData = newData, life.Data
}

func (life *Life) GetData() []bool {
	return life.Data
}

func (life *Life) GetB() []string {
	return ListKeys(life.B)
}

func (life *Life) GetS() []string {
	return ListKeys(life.S)
}
