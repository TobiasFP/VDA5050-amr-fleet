package main

import (
	"TobiasFP/BotNana/config"
	mqttstate "TobiasFP/BotNana/controllers/mqtt"
	"TobiasFP/BotNana/models"
	"time"

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
	config := config.GetConfig()

	logToLogStash := config.GetBool("logging.logToLogStash")
	if logToLogStash {
		log.SetFlags(0)
		log.SetPrefix(time.Now().Format(time.RFC3339) + " ")
		log.SetFlags(log.Lshortfile)
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)
	}
	log.Println("Starting up")

	models.ConnectDatabase()
	models.MigrateDB(models.SqlDB)

	addTestData := config.GetBool("addTestData")
	if addTestData {
		models.AddTestData()
	}

	models.ConnectElastic()

	mqttstate.StartMqtt()
	go mqttroutes.StartSubscribing(mqttstate.Client)
	restroutes.StartGin()
}
