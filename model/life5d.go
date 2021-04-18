package model

import utils "github.com/bigcubecat/gameoflifemodel/utils"

type Life5d struct {
	Data    [][][][][]bool
	newData [][][][][]bool
	SIZE    int
	N       int
	B       map[int]bool
	S       map[int]bool
}

func (life *Life5d) GetN() int {
	return life.N
}

func (life *Life5d) GetSIZE() int {
	return life.SIZE
}

func (life *Life5d) SetN(n int) {
	life.N = n
}

func (life *Life5d) SetSIZE(n int) {
	life.SIZE = n
}

func (life *Life5d) Setup(b []int, s []int, data []bool) {
	life.B = make(map[int]bool)
	life.S = make(map[int]bool)
	for _, i := range b {
		life.B[i] = true
	}
	for _, i := range s {
		life.S[i] = true
	}
	index := 0
	life.Data = make([][][][][]bool, life.SIZE)
	for i := range life.Data {
		life.Data[i] = make([][][][]bool, life.SIZE)
		for j := range life.Data[i] {
			life.Data[i][j] = make([][][]bool, life.SIZE)
			for w := range life.Data[i][j] {
				life.Data[i][j][w] = make([][]bool, life.SIZE)
				for e := range life.Data[i][j][w] {
					life.Data[i][j][w][e] = make([]bool, life.SIZE)
				}
			}
		}
	}
	life.newData = make([][][][][]bool, life.SIZE)
	for i := range life.newData {
		life.newData[i] = make([][][][]bool, life.SIZE)
		for j := range life.newData[i] {
			life.newData[i][j] = make([][][]bool, life.SIZE)
			for w := range life.newData[i][j] {
				life.newData[i][j][w] = make([][]bool, life.SIZE)
				for e := range life.newData[i][j][w] {
					life.newData[i][j][w][e] = make([]bool, life.SIZE)
					for f := range life.newData[i][j][w][e] {
						life.newData[i][j][w][e][f] = false
					}
				}
			}
		}
	}
	fi := len(data) == 0
	data = append(data, false)
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				for w := 0; w < life.SIZE; w++ {
					for f := 0; f < life.SIZE; f++ {
						life.Data[i][j][k][w][f] = data[index]
						life.newData[i][j][k][w][f] = false
						if !fi {
							index++
						}
					}
				}
			}
		}
	}
}

func (life *Life5d) GetData() []bool {
	var answer []bool
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				for w := 0; w < life.SIZE; w++ {
					for e := 0; e < life.SIZE; e++ {
						answer = append(answer, life.Data[i][j][k][w][e])
					}
				}
			}
		}
	}
	return answer
}

func (life *Life5d) getIndex(x int) int {
	return (life.SIZE + x) % life.SIZE
}

func (life *Life5d) NextGeneration() {
	for i := 0; i < life.SIZE; i++ {
		for j := 0; j < life.SIZE; j++ {
			for k := 0; k < life.SIZE; k++ {
				for s := 0; s < life.SIZE; s++ {
					for f := 0; f < life.SIZE; f++ {
						count := 0
						for x := -1; x <= 1; x++ {
							for y := -1; y <= 1; y++ {
								for z := -1; z <= 1; z++ {
									for w := -1; w <= 1; w++ {
										for e := -1; e <= 1; e++ {
											if x == 0 && y == 0 && z == 0 && w == 0 {
												continue
											}
											if life.Data[life.getIndex(i+x)][life.getIndex(j+y)][life.getIndex(k+z)][life.getIndex(s+w)][life.getIndex(f+e)] {
												count += 1
											}
										}
									}
								}
							}
						}
						if life.Data[i][j][k][s][f] {
							life.newData[i][j][k][s][f] = life.S[count]
						} else {
							life.newData[i][j][k][s][f] = life.B[count]
						}
					}
				}
			}
			life.Data, life.newData = life.newData, life.Data
		}
	}
}

func (life *Life5d) GetB() string {
	return utils.ListKeys(life.B)
}

func (life *Life5d) GetS() string {
	return utils.ListKeys(life.S)
}
