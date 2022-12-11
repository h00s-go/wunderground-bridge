package main

import (
	"log"
	"net/http"
	"os"

	"github.com/h00s-go/wunderground-bridge/application"
	"github.com/h00s-go/wunderground-bridge/config"
	"github.com/h00s-go/wunderground-bridge/mqtt"
)

func main() {
	config := config.NewConfig()

	var MQTT *mqtt.MQTT
	if config.MQTT.Enabled {
		MQTT = mqtt.NewMQTT(&config.MQTT)
		defer MQTT.Close()
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	station := application.NewStation(&config.Station, logger)
	app := application.NewApplication(config, logger, station, MQTT)

	logger.Println("Listening on :8080")
	http.ListenAndServe(":8080", app.NewRouter())
}
