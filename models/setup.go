package models

import (
	"TobiasFP/BotNana/conn"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := conn.GetMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&NodeState{},
		&EdgeState{},
		&ActionState{},
		&BatteryState{},
		&State{},
		&StateError{},
		&AgvPosition{},
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
		&Corridor{},
		&NodePosition{},
		&Trajectory{},
		&Edge{},
		&Order{},
		&Node{},
		&NodeMeta{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	DB = db
}
