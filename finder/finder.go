package finder

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/TwinProduction/go-color"
	model "github.com/bigcubecat/gameoflifemodel/model"
	utils "github.com/bigcubecat/gameoflifemodel/utils"
)

func getDen(data string) (uint, bool) {
	var (
		countA float64
	)
	countA = 0
	alive := false
	for _, c := range data {
		if string(c) == "A" {
			countA += 1.0
			alive = true
		}
	}
	return uint((countA / float64(len(data))) * 100), alive
}

func random(probability int) bool {
	value := rand.Intn(100)
	return value <= probability
}

func generateData(probability int, dataSize int) []bool {
	var a []bool
	count := dataSize * probability / 100
	for i := 0; i < count; i++ {
		a = append(a, true)
	}
	for i := 0; i < dataSize-count; i++ {
		a = append(a, false)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func FindRules(mod model.MODEL, G int, T int, fileName string, probability int, b []int, s []int, dataSize int) {
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
		fmt.Println(color.Ize(color.Green, "Starting test "+strconv.Itoa(t)))
		test := Test{
			AttemptID: a.ID,
			Count:     uint(G),
		}
		res := DB.Create(&test)
		if res.Error != nil {
			fmt.Println(color.Ize(color.Red, "ERROR. Create test fail"))
			return
		}
		mod.Setup(b, s, generateData(probability, dataSize))
		outputData := utils.DataToString(mod.GetData())

		str, _ := getDen(outputData)

		gen := Generation{
			TestID:        test.ID,
			Generation:    uint(0),
			Data:          utils.RLECode(outputData),
			StartDensity:  uint(probability),
			FinishDensity: str,
		}
		e := DB.Create(&gen)

		if e.Error != nil {
			fmt.Println(color.Ize(color.Red, "ERROR. Create generation fail"))
			return
		}

		fmt.Println(color.Ize(color.Green, "Starting test"))
		for g := uint(1); g <= uint(G); g++ {
			fmt.Println(color.Ize(color.Cyan, "Generation ->"), g)
			mod.NextGeneration()
			outputData := utils.DataToString(mod.GetData())
			str, alive := getDen(outputData)
			gen := Generation{
				TestID:        test.ID,
				Generation:    g,
				Data:          utils.RLECode(outputData),
				StartDensity:  uint(probability),
				FinishDensity: str,
			}
			e := DB.Create(&gen)

			if e.Error != nil {
				fmt.Println(color.Ize(color.Red, "ERROR. Create generation fail"))
				return
			}
			if !alive {
				break
			}
		}
		fmt.Println(color.Ize(color.Green, "End evolution"))
	}
}
