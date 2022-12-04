package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/h00s-go/wunderground-bridge/config"
	"github.com/h00s-go/wunderground-bridge/mqtt"
)

type Application struct {
	config  *config.Config
	logger  *log.Logger
	mqtt    *mqtt.MQTT
	weather *Weather
}

func NewApplication(config *config.Config, logger *log.Logger, mqtt *mqtt.MQTT) *Application {
	return &Application{
		config:  config,
		logger:  logger,
		mqtt:    mqtt,
		weather: &Weather{},
	}
}

func (app *Application) publishWeatherToMQTT() {
	w, err := json.Marshal(app.weather)
	if err == nil {
		app.mqtt.Publish(app.config.MQTT.UpdateTopic, string(w))
	}
}

func (app *Application) updateWunderground(query string) {
	url := fmt.Sprintf("%v?%v", app.config.Wunderground.UpdateURL, query)
	_, err := http.Get(url)
	if err != nil {
		app.logger.Printf("Error updating wunderground: %v\n", err)
	}
}
