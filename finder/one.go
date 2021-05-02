package finder

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bigcubecat/gameoflifemodel/model"
	"github.com/bigcubecat/gameoflifemodel/utils"
)

// DataModel data about world from json
type DataModel struct {
	D    int
	SIZE int
	DATA string
	B    string
	S    string
}

// ReadJSON : read json file and return data
func ReadJSON(attempts int, fileName string) (DataModel, error) {
	jsonFile, err := os.Open(fileName)
	defer jsonFile.Close()
	if err != nil {
		return DataModel{}, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return DataModel{}, err
	}
	var dat DataModel
	if err := json.Unmarshal(byteValue, &dat); err != nil {
		return DataModel{}, err
	}
	return dat, nil
}

// RunFromFile run
func RunFromFile(m model.MODEL, attempt int, out string) error {
	for i := 0; i < attempt; i++ {
		m.NextGeneration()
	}
	return writeToFile(m, out)
}

func writeToFile(m model.MODEL, output string) error {
	jsonData, err := json.Marshal(toDataModel(m))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(output, jsonData, 0644)
}

func toDataModel(m model.MODEL) DataModel {
	var dat DataModel
	dat.DATA = utils.RLECode(utils.DataToString(m.GetData()))
	dat.B = m.GetB()
	dat.S = m.GetS()
	dat.D = m.GetN()
	dat.SIZE = m.GetSIZE()
	return dat
}
