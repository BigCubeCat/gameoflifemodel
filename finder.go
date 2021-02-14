package main

import "fmt"
import "strings"
import "strconv"
import "encoding/json"
import "io/ioutil"

func findRules(size int, D int, G int) {
	var b []int
	var s []int
	c := intPow(3, D)
	end := (4 * c) / 9
	for i := c / 3; i < end; i++ {
		b = append(b, i)
	}
	for i := (2 * c) / 9; i < end; i++ {
		s = append(s, i)
	}
	for t := 0; t < 100; t++ {
		life := Life{N: D, SIZE: size}
		var d []bool

		d = make([]bool, intPow(size, D))
		fmt.Print(len(d))
		for x := range d {
			d[x] = random() && random()
		}
		life.Setup(b, s, d)
		fmt.Println("Starting Evolution")
		for g := 0; g < G; g++ {
			life.NextGeneration()
			fmt.Println("Start coding")
			var s_string []string
			var b_string []string
			for _, value := range s {
				s_string = append(s_string, strconv.Itoa(value))
			}
			for _, value := range b {
				b_string = append(b_string, strconv.Itoa(value))
			}
			out := RLECode(dataToString(life.GetData()))
			output := dataModel{
				D:    life.N,
				SIZE: life.SIZE,
				DATA: out,
				B:    strings.Join(b_string, ","),
				S:    strings.Join(s_string, ","),
			}
			file, _ := json.MarshalIndent(output, "", "")
			outputFile := "data" + strconv.Itoa(D) + "try" + strconv.Itoa(t) + ".life"
			fmt.Println("Starting write to file")
			saveErr := ioutil.WriteFile(outputFile, file, 0644)
			if saveErr != nil {
				panic(saveErr)
			} else {
				fmt.Println("Finish write to file. No errors")
			}
			fmt.Println("End evolution")
		}
	}
}
