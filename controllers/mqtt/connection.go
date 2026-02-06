package mqttstate

import (
	"TobiasFP/BotNana/models"
	"encoding/json"
	"errors"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func OnConnectionReceived(_ mqtt.Client, message mqtt.Message) {
	msg := message.Payload()
	topic := message.Topic()
	log.Printf("Received connection message: %s from topic: %s\n", msg, topic)

	var connectionFromMqtt models.Connection
	unmarshallErr := json.Unmarshal([]byte(msg), &connectionFromMqtt)
	if unmarshallErr != nil {
		log.Printf("unable to parse connection payload: %v", unmarshallErr)
		return
	}
	if manufacturer, serial, leaf, ok := parseTopic(topic); ok && leaf == topicConnection {
		if connectionFromMqtt.Manufacturer == "" {
			connectionFromMqtt.Manufacturer = manufacturer
		} else if connectionFromMqtt.Manufacturer != manufacturer {
			log.Printf("Connection topic manufacturer mismatch: topic=%s payload=%s", manufacturer, connectionFromMqtt.Manufacturer)
		}
		if connectionFromMqtt.SerialNumber == "" {
			connectionFromMqtt.SerialNumber = serial
		} else if connectionFromMqtt.SerialNumber != serial {
			log.Printf("Connection topic serial mismatch: topic=%s payload=%s", serial, connectionFromMqtt.SerialNumber)
		}
	}

	var connectionInDB models.Connection
	connectionInDBResult := models.SqlDB.Where("serial_number = ?", connectionFromMqtt.SerialNumber).First(&connectionInDB)
	if connectionInDBResult.Error != nil {
		if errors.Is(connectionInDBResult.Error, gorm.ErrRecordNotFound) {
			// If the connection is not in the db, create it.
		} else {
			log.Printf("Error fetching connection from db: %v", connectionInDBResult.Error)
			return
		}
	}

	if connectionInDBResult.RowsAffected == 1 {
		updateErr := models.UpdateConnectionInDb(models.SqlDB, connectionInDB, connectionFromMqtt)
		if updateErr != nil {
			log.Printf("Error updating connection in db: %v", updateErr)
		}
		return
	}

	createErr := models.CreateConnectionInDb(models.SqlDB, connectionFromMqtt)
	if createErr != nil {
		log.Printf("Error creating connection in db: %v", createErr)
	}
}
