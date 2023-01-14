package services

import (
	"log"

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
