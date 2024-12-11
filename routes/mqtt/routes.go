package mqttroutes

import (
	"TobiasFP/BotNana/config"
	mqttstate "TobiasFP/BotNana/controllers/mqtt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartMqtt() {
	conf := config.GetConfig()
	broker := conf.GetString("mqttBroker")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("Error connecting to MQTT broker:", token.Error())
	}

	stateTopic := "state"
	if token := client.Subscribe(stateTopic, 0, mqttstate.OnStateReceived); token.Wait() && token.Error() != nil {
		log.Panic("Error subscribing to topic:", token.Error())
	}

	log.Println("Subscribed to topic:", stateTopic)

	// Wait for a signal to exit the program gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	client.Unsubscribe(stateTopic)
	client.Disconnect(250)
}
