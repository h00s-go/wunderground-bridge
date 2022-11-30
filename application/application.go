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
	config *config.Config
	logger *log.Logger
	mqtt   *mqtt.MQTT
}

func NewApplication(config *config.Config, logger *log.Logger, mqtt *mqtt.MQTT) *Application {
	return &Application{
		config: config,
		logger: logger,
		mqtt:   mqtt,
	}
}

func (app *Application) publishWeatherToMQTT(weather *Weather) {
	w, err := json.Marshal(weather)
	if err == nil {
		app.mqtt.Publish(app.config.MQTT.WeatherTopic, fmt.Sprint(w))
	}
}

func (app *Application) updateWunderground(query string) {
	if app.config.Wunderground.PassUpdate {
		url := fmt.Sprintf("%v?%v", app.config.Wunderground.UpdateURL, query)
		_, err := http.Get(url)
		if err != nil {
			app.logger.Printf("Error: %v\n", err)
		}
	}
}
