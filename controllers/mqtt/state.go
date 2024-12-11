package mqttstate

import (
	"TobiasFP/BotNana/models"
	"encoding/json"
	"errors"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func OnStateReceived(_ mqtt.Client, message mqtt.Message) {
	msg := message.Payload()
	topic := message.Topic()
	log.Printf("Received message: %s from topic: %s\n", msg, topic)

	var amrFromMqtt models.State
	unmarshallErr := json.Unmarshal([]byte(msg), &amrFromMqtt)
	if unmarshallErr != nil {
		log.Fatal(unmarshallErr.Error())
	}

	// Check if already in DB:
	var amrInDB models.State
	amrInDBResult := models.DB.Preload("AgvPosition").First(&amrInDB).Where("SerialNumber = ?", amrFromMqtt.SerialNumber)
	if amrInDBResult.Error != nil {
		if errors.Is(amrInDBResult.Error, gorm.ErrRecordNotFound) {
			// If the amr is simply not in the db, create it.
		} else {
			log.Print("Error fetching from db")
			log.Print(amrInDBResult.Error.Error())
			return
		}
	}

	if amrInDBResult.RowsAffected == 1 {
		// if the amr is already in the db, just update it.
		models.UpdateAmrStateInDb(models.DB, amrInDB, amrFromMqtt)
		return
	}

	unmarshallErr = models.CreateAmrStateInDb(models.DB, amrFromMqtt)
	if unmarshallErr != nil {
		log.Print("Error creating the mqtt amr in db")
		log.Print(unmarshallErr.Error())
	}
}
