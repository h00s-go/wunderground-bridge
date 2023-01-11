package services

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/h00s-go/wunderground-bridge/mqtt"
	"github.com/h00s-go/wunderground-bridge/station"
)

type Services struct {
	Logger            *log.Logger
	Station           *station.Station
	Wunderground      *station.Wunderground
	MQTT              *mqtt.MQTT
	WebsocketUpgrader *websocket.Upgrader
}
