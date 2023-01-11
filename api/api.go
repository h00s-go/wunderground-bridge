package api

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/h00s-go/wunderground-bridge/api/services"
	"github.com/h00s-go/wunderground-bridge/config"
	"github.com/h00s-go/wunderground-bridge/mqtt"
	"github.com/h00s-go/wunderground-bridge/station"
)

type API struct {
	config   *config.Config
	services *services.Services
}

func NewAPI(config *config.Config, logger *log.Logger, station *station.Station, wunderground *station.Wunderground, mqtt *mqtt.MQTT) *API {
	return &API{
		config: config,
		services: &services.Services{
			Logger:       logger,
			Station:      station,
			Wunderground: wunderground,
			MQTT:         mqtt,
			WebsocketUpgrader: &websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	}
}
