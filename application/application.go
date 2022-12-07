package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/h00s-go/wunderground-bridge/config"
	"github.com/h00s-go/wunderground-bridge/mqtt"
)

type Application struct {
	config            *config.Config
	logger            *log.Logger
	station           *Station
	mqtt              *mqtt.MQTT
	websocketUpgrader *websocket.Upgrader
}

func NewApplication(config *config.Config, logger *log.Logger, station *Station, mqtt *mqtt.MQTT) *Application {
	return &Application{
		config:  config,
		logger:  logger,
		station: station,
		mqtt:    mqtt,
		websocketUpgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (app *Application) publishWeatherToMQTT() {
	w, err := json.Marshal(app.station.Weather)
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
