package model

type Life4d struct {
	Data    [][][][]bool
	newData [][][][]bool
	SIZE    int
	N       int
	B       map[int]bool
	S       map[int]bool
}

func (life *Life4d) GetN() int {
	return life.N
}

func (life *Life4d) GetSIZE() int {
	return life.SIZE
}

func (life *Life4d) SetN(n int) {
	life.N = n
}

func (life *Life4d) SetSIZE(n int) {
	life.SIZE = n
}

func (life *Life4d) Setup(b []int, s []int, data []bool) {
	life.B = make(map[int]bool)
	life.S = make(map[int]bool)
	for _, i := range b {
		life.B[i] = true
	}
	for _, i := range s {
		life.S[i] = true
	}
	index := 0
	life.Data = make([][][][]bool, life.SIZE)
	for i := range life.Data {
		life.Data[i] = make([][][]bool, life.SIZE)
		for j := range life.Data[i] {
			life.Data[i][j] = make([][]bool, life.SIZE)
			for w := range life.Data[i][j] {
				life.Data[i][j][w] = make([]bool, life.SIZE)
			}
		}
	}
	life.newData = make([][][][]bool, life.SIZE)
	for i := range life.newData {
		life.newData[i] = make([][][]bool, life.SIZE)
		for j := range life.newData[i] {
			life.newData[i][j] = make([][]bool, life.SIZE)
			for w := range life.newData[i][j] {
				life.newData[i][j][w] = make([]bool, life.SIZE)
			}
		}
	}
	f := len(data) == 0
	data = append(data, false)
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				for w := 0; w < life.SIZE; w++ {
					life.Data[i][j][k][w] = data[index]
					life.newData[i][j][k][w] = false
					if !f {
						index++
					}
				}
			}
		}
	}
}

func (life *Life4d) GetData() []bool {
	var answer []bool
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				for w := 0; w < life.SIZE; w++ {
					answer = append(answer, life.Data[i][j][k][w])
				}
			}
		}
	}
	return answer
}

func (life *Life4d) getIndex(x int) int {
	return (life.SIZE + x) % life.SIZE
}

func (life *Life4d) NextGeneration() {
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				for s := 0; s < life.SIZE; s++ {
					count := 0
					for x := -1; x <= 1; x++ {
						for y := -1; y <= 1; y++ {
							for z := -1; z <= 1; z++ {
								for w := -1; w <= 1; w++ {
									if x == 0 && y == 0 && z == 0 && w == 0 {
										continue
									}
									if life.Data[life.getIndex(i+x)][life.getIndex(j+y)][life.getIndex(k+z)][life.getIndex(s+w)] {
										count += 1
									}
								}
							}
						}
					}
					if life.Data[i][j][k][s] {
						life.newData[i][j][k][s] = life.S[count]
					} else {
						life.newData[i][j][k][s] = life.B[count]
					}
				}
			}
			life.Data, life.newData = life.newData, life.Data
		}
	}
}

func (life *Life4d) GetB() string {
	return ListKeys(life.B)
}

func (life *Life4d) GetS() string {
	return ListKeys(life.S)
}
