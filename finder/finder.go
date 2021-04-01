package finder

import (
	"fmt"
	model "gameoflifemodel/model"
	"github.com/TwinProduction/go-color"
)

func FindRules(mod model.MODEL, G int, T int, fileName string) {
	InitDatabase(fileName)
	a := Attempt{
		B:         mod.GetB(),
		S:         mod.GetS(),
		Size:      uint(mod.GetSIZE()),
		Dimension: uint(mod.GetN()),
		Count:     uint(T),
	}
	attempt := DB.Create(&a)

	if attempt.Error != nil {
		fmt.Println(color.Ize(color.Red, "ERROR. Create Attempt fail"))
		return
	}
	for t := 0; t < T; t++ {
		test := Test{
			AttemptID: a.ID,
			Count:     uint(G),
		}
		res := DB.Create(&test)
		if res.Error != nil {
			fmt.Println(color.Ize(color.Red, "ERROR. Create test fail"))
			return
		}
		for g := uint(0); g < uint(G); g++ {
			fmt.Println(color.Ize(color.Green, "Starting Evolution"))
			fmt.Println(color.Ize(color.Cyan, "Generation ->"), g)
			mod.NextGeneration()
			gen := Generation{
				TestID:     test.ID,
				Generation: g,
				Data:       model.RLECode(model.DataToString(mod.GetData())),
			}
			e := DB.Create(&gen)

			if e.Error != nil {
				fmt.Println(color.Ize(color.Red, "ERROR. Create generation fail"))
				return
			}
		}
		fmt.Println(color.Ize(color.Green, "End evolution"))
	}
}
