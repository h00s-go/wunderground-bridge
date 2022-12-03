package application

import (
	"encoding/json"
	"net/http"
)

func (app *Application) weatherHandler(w http.ResponseWriter, r *http.Request) {
	weather, err := json.MarshalIndent(app.weather, "", "\t")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(weather)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) weatherUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	app.weather, err = NewWeatherFromStation(r.URL.RawQuery)
	if err == nil {
		if app.config.MQTT.Enabled {
			go app.publishWeatherToMQTT()
		}
		if app.config.Wunderground.Enabled {
			go app.updateWunderground(r.URL.RawQuery)
		}
	} else {
		app.logger.Println("Error parsing weather: ", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func (app *Application) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
