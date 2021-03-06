package finder

import (
	"fmt"

	"github.com/TwinProduction/go-color"
	model "github.com/bigcubecat/gameoflifemodel/model"
	TUI "github.com/bigcubecat/gameoflifemodel/tui"
	utils "github.com/bigcubecat/gameoflifemodel/utils"
)

// Run run model
func Run(mod model.MODEL, G int, T int, fileName string, probability int, b []int, s []int, dataSize int, change chan TUI.ChangeModel) {
	InitDatabase(fileName)
	a := Attempt{
		Size:      uint(mod.GetSIZE()),
		Dimension: uint(mod.GetN()),
		Count:     uint(T),
	}
	rulesWritten := false
	attempt := DB.Create(&a)
	if attempt.Error != nil {
		fmt.Println(color.Ize(color.Red, "ERROR. Create Attempt fail"))
		return
	}
	for t := 0; t < T; t++ {
		//fmt.Println(color.Ize(color.Green, "Starting test "+strconv.Itoa(t)))
		test := Test{
			AttemptID:     a.ID,
			Count:         uint(G), // Need update, if finish early
			StartDensity:  uint(probability),
			FinishDensity: uint(probability),
		}
		res := DB.Create(&test)
		if res.Error != nil {
			fmt.Println(color.Ize(color.Red, "ERROR. Create test fail"))
			return
		}
		mod.Setup(b, s, utils.GenerateData(probability, dataSize))
		if !rulesWritten {
			rulesWritten = true
			DB.Model(&a).Update("B", mod.GetB())
			DB.Model(&a).Update("S", mod.GetS())
		}
		early := false
		count := uint(G)
		outputData := ""
		var alive bool
		//fmt.Println(color.Ize(color.Green, "Starting test"))
		_t := float64(t) / float64(T)
		change <- TUI.ChangeModel{A: _t, G: float64(0)}
		for g := uint(0); g <= uint(G); g++ {
			//fmt.Println(color.Ize(color.Cyan, "Generation ->"), g)
			outputData = utils.DataToString(mod.GetData())
			alive = utils.IsAlive(outputData)
			gen := Generation{
				TestID:     test.ID,
				Generation: g,
				Data:       utils.RLECode(outputData),
			}
			e := DB.Create(&gen)

			if e.Error != nil {
				fmt.Println(color.Ize(color.Red, "ERROR. Create generation fail"))
				return
			}
			if !alive {
				early = true
				count = g
				break
			}
			mod.NextGeneration()
			change <- TUI.ChangeModel{A: _t, G: float64(g) / float64(G)}
		}
		//fmt.Println(color.Ize(color.Green, "End evolution"))
		DB.Model(&test).Update("FinishDensity", uint(utils.GetDensity(outputData)))
		DB.Model(&test).Update("Alive", alive)
		if early {
			DB.Model(&test).Update("Count", count)
		}
	}
	change <- TUI.ChangeModel{Finished: true}
}
