package application

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

func (app *Application) weatherHandler(w http.ResponseWriter, r *http.Request) {
	weather, err := json.MarshalIndent(app.station.Weather, "", "\t")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(weather)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) weatherUpdateHandler(w http.ResponseWriter, r *http.Request) {
	c, err := app.websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		app.logger.Print("Error while upgrading:", err)
		return
	}
	defer c.Close()
	for {
		weather, err := json.Marshal(app.station.Weather)
		if err == nil {
			err = c.WriteMessage(websocket.TextMessage, weather)
			if err != nil {
				app.logger.Print("Error while writing to websocket:", err)
				return
			}
		}
	}
}

func (app *Application) stationUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.station.NewWeather(r.URL.RawQuery); err == nil {
		if app.config.MQTT.Enabled {
			go app.publishWeatherToMQTT()
		}
		if app.config.Wunderground.Enabled {
			go app.updateWunderground(r.URL.RawQuery)
		}
	} else {
		app.logger.Println("Error parsing weather: ", err, r.URL.RawQuery)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func (app *Application) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
