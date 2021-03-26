package main

import "fmt"
import "strings"
import "encoding/json"
import "io/ioutil"

func findRules(model MODEL, G int, det bool, t int) {
	fmt.Println("Starting Evolution")
	for g := 0; g < G; g++ {
		fmt.Println("Generation ->", g)
		model.NextGeneration()
		if det {
			saveState(model, fmt.Sprintf("t%vg%v.json", t, g))
		}
	}
	fmt.Println("End evolution")
	if !det {
		saveState(model, fmt.Sprintf("t%v.json", t))
	}
}

func saveState(mod MODEL, fileName string) error {
	s_string := mod.GetB()
	b_string := mod.GetS()
	fmt.Println(mod.GetData())
	out := RLECode(dataToString(mod.GetData()))
	output := dataModel{
		D:    mod.getN(),
		SIZE: mod.getSIZE(),
		DATA: out,
		B:    strings.Join(b_string, ","),
		S:    strings.Join(s_string, ","),
	}
	file, _ := json.MarshalIndent(output, "", "")
	saveErr := ioutil.WriteFile(fileName, file, 0644)
	return saveErr
}
