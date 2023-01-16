package services

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/h00s-go/wunderground-bridge/api/models"
	"github.com/h00s-go/wunderground-bridge/mqtt"
)

type Services struct {
	Logger            *log.Logger
	Station           *models.Station
	Wunderground      *models.Wunderground
	MQTT              *mqtt.MQTT
	WebsocketUpgrader *websocket.Upgrader
}

func NewServices(logger *log.Logger, station *models.Station, wunderground *models.Wunderground, mqtt *mqtt.MQTT) *Services {
	return &Services{
		Logger:       logger,
		Station:      station,
		Wunderground: wunderground,
		MQTT:         mqtt,
		WebsocketUpgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
