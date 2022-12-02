package application

import (
	"encoding/json"
	"net/http"
)

func (app *Application) weatherHandler(w http.ResponseWriter, r *http.Request) {
	weather, err := json.Marshal(app.weather)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(weather)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) weatherUpdateHandler(w http.ResponseWriter, r *http.Request) {
	weather, err := NewWeatherFromStation(r.URL.RawQuery)
	if err == nil {
		go app.publishWeatherToMQTT(weather)
	} else {
		app.logger.Println("Error parsing weather: ", err)
	}
	w.WriteHeader(http.StatusOK)
	go app.updateWunderground(r.URL.RawQuery)
}

func (app *Application) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
