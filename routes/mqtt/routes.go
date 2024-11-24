package mqttroutes

import (
	"TobiasFP/BotNana/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}

func StartMqtt() {
	conf := config.GetConfig()
	broker := conf.GetString("mqttBroker")
	topic := "rtr_idle"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("Error connecting to MQTT broker:", token.Error())
	}

	if token := client.Subscribe(topic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
		log.Panic("Error subscribing to topic:", token.Error())
	}

	log.Println("Subscribed to topic:", topic)

	// Wait for a signal to exit the program gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	client.Unsubscribe(topic)
	client.Disconnect(250)

}
