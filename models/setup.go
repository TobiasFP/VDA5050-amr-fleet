package models

import (
	"TobiasFP/BotNana/conn"
	"log"

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
