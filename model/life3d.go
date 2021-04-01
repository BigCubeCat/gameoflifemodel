package model

type Life3d struct {
	Data    [][][]bool
	newData [][][]bool
	SIZE    int
	N       int
	B       map[int]bool
	S       map[int]bool
}

func (life *Life3d) GetN() int {
	return life.N
}

func (life *Life3d) GetSIZE() int {
	return life.SIZE
}

func (life *Life3d) SetN(n int) {
	life.N = n
}

func (life *Life3d) SetSIZE(n int) {
	life.SIZE = n
}

func (life *Life3d) Setup(b []int, s []int, data []bool) {
	life.B = make(map[int]bool)
	life.S = make(map[int]bool)
	for _, i := range b {
		life.B[i] = true
	}
	for _, i := range s {
		life.S[i] = true
	}
	index := 0
	life.Data = make([][][]bool, life.SIZE)
	for i := range life.Data {
		life.Data[i] = make([][]bool, life.SIZE)
		for j := range life.Data[i] {
			life.Data[i][j] = make([]bool, life.SIZE)
		}
	}
	life.newData = make([][][]bool, life.SIZE)
	for i := range life.newData {
		life.newData[i] = make([][]bool, life.SIZE)
		for j := range life.newData[i] {
			life.newData[i][j] = make([]bool, life.SIZE)
		}
	}
	f := len(data) == 0
	data = append(data, false)
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				life.Data[i][j][k] = data[index]
				life.newData[i][j][k] = false
				if !f {
					index++
				}
			}
		}
	}
}

func (life *Life3d) GetData() []bool {
	var answer []bool
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				answer = append(answer, life.Data[i][j][k])
			}
		}
	}
	return answer
}

func (life *Life3d) getIndex(x int) int {
	return (life.SIZE + x) % life.SIZE
}

func (life *Life3d) NextGeneration() {
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				for x := -1; x <= 1; x++ {
					count := 0
					for y := -1; y <= 1; y++ {
						for z := -1; z <= 1; z++ {
							if x == 0 && y == 0 && z == 0 {
								continue
							}
							if life.Data[life.getIndex(x)][life.getIndex(y)][life.getIndex(z)] {
								count += 1
							}
						}
					}
					if life.Data[i][j][k] {
						life.newData[i][j][k] = life.S[count]
					} else {
						life.newData[i][j][k] = life.B[count]
					}
				}
			}
		}
	}
	life.Data, life.newData = life.newData, life.Data
}

func (life *Life3d) GetB() string {
	return ListKeys(life.B)
}

func (life *Life3d) GetS() string {
	return ListKeys(life.S)
}
