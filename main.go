package main

import (
	"TobiasFP/BotNana/config"
	mqttstate "TobiasFP/BotNana/controllers/mqtt"
	"TobiasFP/BotNana/models"

	mqttroutes "TobiasFP/BotNana/routes/mqtt"
	restroutes "TobiasFP/BotNana/routes/rest"
	"flag"
	"log"

	"os"
)

func main() {
	environment := flag.String("e", "production", "")
	flag.Parse()

	config.Init(*environment)

	if *environment != "production" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)
	}

	log.Println("Starting up")

	models.ConnectDatabase()
	models.MigrateDB(models.DB)
	mqttstate.StartMqtt()
	go mqttroutes.StartSubscribing(mqttstate.Client)
	restroutes.StartGin()
}
