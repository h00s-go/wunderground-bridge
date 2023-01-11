package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/h00s-go/wunderground-bridge/api/services"
)

type WeatherController struct {
	services *services.Services
}

func NewWeatherController(services *services.Services) *WeatherController {
	return &WeatherController{
		services: services,
	}
}

func (wc *WeatherController) WeatherHandler(w http.ResponseWriter, r *http.Request) {
	weather, err := json.MarshalIndent(wc.services.Station.Weather, "", "\t")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(weather)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (wc *WeatherController) WeatherUpdateHandler(w http.ResponseWriter, r *http.Request) {
	c, err := wc.services.WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		wc.services.Logger.Print("Error while upgrading:", err)
		return
	}
	defer c.Close()
	for {
		weather, err := json.Marshal(wc.services.Station.Weather)
		if err == nil {
			err = c.WriteMessage(websocket.TextMessage, weather)
			if err != nil {
				wc.services.Logger.Print("Error while writing to websocket:", err)
				return
			}
		}
	}
}
