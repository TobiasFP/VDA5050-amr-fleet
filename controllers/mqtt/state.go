package mqttstate

import (
	"TobiasFP/BotNana/config"
	"TobiasFP/BotNana/models"
	"encoding/json"
	"errors"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

var Client mqtt.Client

type Publisher interface {
	Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token
}

func StartMqtt() {
	conf := config.GetConfig()
	broker := conf.GetString("mqttBroker")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("Error connecting to MQTT broker:", token.Error())
	}
	Client = client
}

func OnStateReceived(_ mqtt.Client, message mqtt.Message) {
	msg := message.Payload()
	topic := message.Topic()
	log.Printf("Received message: %s from topic: %s\n", msg, topic)

	var amrFromMqtt models.State
	unmarshallErr := json.Unmarshal([]byte(msg), &amrFromMqtt)
	if unmarshallErr != nil {
		log.Fatal(unmarshallErr.Error())
	}
	if manufacturer, serial, leaf, ok := parseTopic(topic); ok && leaf == topicState {
		if amrFromMqtt.Manufacturer == "" {
			amrFromMqtt.Manufacturer = manufacturer
		} else if amrFromMqtt.Manufacturer != manufacturer {
			log.Printf("State topic manufacturer mismatch: topic=%s payload=%s", manufacturer, amrFromMqtt.Manufacturer)
		}
		if amrFromMqtt.SerialNumber == "" {
			amrFromMqtt.SerialNumber = serial
		} else if amrFromMqtt.SerialNumber != serial {
			log.Printf("State topic serial mismatch: topic=%s payload=%s", serial, amrFromMqtt.SerialNumber)
		}
	}

	// Check if already in DB:
	var amrInDB models.State
	amrInDBResult := models.SqlDB.Preload("AgvPosition").Where("serial_number = ?", amrFromMqtt.SerialNumber).First(&amrInDB)
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
		models.UpdateAmrStateInDb(models.SqlDB, amrInDB, amrFromMqtt)
		return
	}

	unmarshallErr = models.CreateAmrStateInDb(models.SqlDB, amrFromMqtt)
	if unmarshallErr != nil {
		log.Print("Error creating the mqtt amr in db")
		log.Print(unmarshallErr.Error())
	}
}

func AssignOrder(client Publisher, order models.Order) error {
	message, err := json.Marshal(order)
	if err != nil {
		return err
	}
	topic, topicErr := OrderTopic(order.Manufacturer, order.SerialNumber)
	if topicErr != nil {
		return topicErr
	}
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	return token.Error()

}

func PublishInstantAction(client Publisher, action models.InstantAction) error {
	message, err := json.Marshal(action)
	if err != nil {
		return err
	}
	topic, topicErr := InstantActionsTopic(action.Manufacturer, action.SerialNumber)
	if topicErr != nil {
		return topicErr
	}
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	return token.Error()
}
