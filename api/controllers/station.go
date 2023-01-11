package controllers

import (
	"net/http"

	"github.com/h00s-go/wunderground-bridge/api/services"
)

type StationController struct {
	services *services.Services
}

func NewStationController(services *services.Services) *StationController {
	return &StationController{
		services: services,
	}
}

func (sc *StationController) StationUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if err := sc.services.Station.NewWeather(r); err == nil {
		go sc.services.Station.PublishWeatherToMQTT(sc.services.MQTT)
		go sc.services.Station.UpdateWunderground(sc.services.Wunderground, r.URL.RawQuery)
	} else {
		sc.services.Logger.Println("Error parsing weather: ", err, r.URL.RawQuery)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
