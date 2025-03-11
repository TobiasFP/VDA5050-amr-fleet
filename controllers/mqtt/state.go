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

func AssignOrder(client mqtt.Client, order models.Order) {
	message, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}
	token := client.Publish("state", 0, false, message)
	token.Wait()

}
