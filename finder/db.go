package finder

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Attempt struct {
	ID        uint `gorm:"primaryKey"`
	B         string
	S         string
	Size      uint
	Dimension uint
	Count     uint
	Tests     []Test `gorm:"foreignKey:AttemptID"`
}

type Test struct {
	ID            uint `gorm:"primaryKey"`
	AttemptID     uint
	Count         uint
	Alive         bool
	StartDensity  uint
	FinishDensity uint
	Generations   []Generation `gorm:"foreignKey:TestID"`
}

type Generation struct {
	ID         uint `gorm:"primaryKey"`
	TestID     uint
	Generation uint `gorm:"default:0"`
	Data       string
}

// InitDatabase init database
func InitDatabase(databaseName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(databaseName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error)},
	)
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Attempt{}, &Test{}, &Generation{})
}
