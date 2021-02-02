package main

import "fmt"
import "strings"
import "strconv"
import "encoding/json"
import "io/ioutil"

func countNeighbours(d int) int {
	return intPow(3, d) - 1
}

func generateS(b int, d int) []int {
	var answer []int
	half := d / 2
	o := 0
	if d%2 == 0 {
		o = 1
	}
	for i := b - half; i <= b+half-o; i++ {
		answer = append(answer, i)
	}
	return answer
}

func findRules(size int) {
	for D := 4; D < 11; D++ {
		var b []int
		var s []int
		v := countNeighbours(D)
		q := v / 4
		for b_variant := q - D; b_variant <= q+D; b_variant++ {
			if b_variant <= 0 {
				continue
			}
			b = append(b, b_variant)
			s = generateS(b_variant, D)
			for t := 0; t < 100; t++ {
				life := Life{N: D, SIZE: size}
				var d []bool

				d = make([]bool, intPow(size, D))
				fmt.Print(len(d))
				for x := range d {
					d[x] = random() && random()
				}
				life.Setup(b, s, d)
				fmt.Println(life.dataSize)

				fmt.Println("Starting Evolution")
				for g := 0; g < 10; g++ {
					life.NextGeneration()
				}
				b_string := strconv.Itoa(b_variant)
				var s_string []string
				for _, value := range s {
					s_string = append(s_string, strconv.Itoa(value))
				}
				out := RLECode(dataToString(life.GetData()))
				fmt.Println("Starting write to file")
				output := dataModel{
					D:    life.N,
					SIZE: life.SIZE,
					DATA: out,
					B:    b_string,
					S:    strings.Join(s_string, ","),
				}
				file, _ := json.MarshalIndent(output, "", "")
				outputFile := "data" + strconv.Itoa(D) + "B" + strconv.Itoa(b_variant) + "try" + strconv.Itoa(t) + ".json"
				saveErr := ioutil.WriteFile(outputFile, file, 0644)
				if saveErr != nil {
					panic(saveErr)
				} else {
					fmt.Println("Finish write to file. No errors")
				}
			}
		}
	}
}
