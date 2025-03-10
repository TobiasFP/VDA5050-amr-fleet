package models

import (
	"TobiasFP/BotNana/conn"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&NodeState{},
		&EdgeState{},
		&ActionState{},
		&BatteryState{},
		&State{},
		&StateError{},
		&AgvPosition{},
		&AmrMapData{},
		&AmrMap{},
		&SafetyState{},
		&Info{},
		&InfoReference{},
		&Load{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.AutoMigrate(
		&ActionParameter{},
		&Action{},
		&InstantAction{},
		&Corridor{},
		&NodePosition{},
		&Trajectory{},
		&Edge{},
		&Node{},
		&NodeMeta{},
		&Order{},
		&OrderTemplate{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.AutoMigrate(
		&OrderTemplateDetails{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ConnectTestingDatabase() {
	db, err := gorm.Open(sqlite.Open("gormtest.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = db
}

func ConnectDatabase() {
	db, err := conn.GetMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
}

func AddTestData() {
	addTestMap()
}

func addTestMap() {
	// Only add test data if there is no data in the database
	var maps []AmrMap
	DB.Find(&maps)
	if len(maps) > 0 {
		return
	}

	var amrMap AmrMap
	mapFile, err := os.Open("assets/maps/99187cd1-8b4b-4f5a-ac11-e455928409de.pgm")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	defer func() {
		if err := mapFile.Close(); err != nil {
			panic(err)
		}
	}()

	amrMap.MapDescription = "Test Map"

	byteContainer, err := io.ReadAll(mapFile)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	id := uuid.MustParse("99187cd1-8b4b-4f5a-ac11-e455928409de")
	amrMap.MapID = id
	amrMap.MapData.Data = byteContainer
	createErr := DB.Create(&amrMap)
	if createErr.Error != nil {
		log.Fatal(createErr.Error.Error())
		return
	}
}
