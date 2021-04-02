package finder

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Attempt struct {
	gorm.Model
	B         string
	S         string
	Size      uint
	Dimension uint
	Count     uint
	Tests     []Test `gorm:"foreignKey:AttemptID"`
}

type Test struct {
	gorm.Model
	AttemptID   uint
	Count       uint
	Generations []Generation `gorm:"foreignKey:TestID"`
}

type Generation struct {
	gorm.Model
	TestID        uint
	Generation    uint `gorm:"default:0"`
	StartDensity  uint
	FinishDensity uint
	Data          string
}

func InitDatabase(databaseName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Attempt{}, &Test{}, &Generation{})
}
