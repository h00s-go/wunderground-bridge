package main

import (
	"log"
	"os"

	"github.com/h00s-go/wunderground-bridge/api"
	"github.com/h00s-go/wunderground-bridge/api/models"
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
	wunderground := models.NewWunderground(&config.Wunderground)
	station := models.NewStation(&config.Station, logger)

	api := api.NewAPI(config, logger, station, wunderground, MQTT)
	api.Start()
	api.WaitForShutdown()
}
