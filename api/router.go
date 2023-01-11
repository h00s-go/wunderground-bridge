package api

import (
	"net/http"

	"github.com/h00s-go/wunderground-bridge/api/controllers"
)

func (api *API) NewRouter() *http.ServeMux {
	wc := controllers.NewWeatherController(api.services)
	sc := controllers.NewStationController(api.services)

	m := http.NewServeMux()
	m.HandleFunc("/weatherstation/updateweatherstation.php", sc.StationUpdateHandler)

	m.HandleFunc("/api/weather/update", wc.WeatherUpdateHandler)
	m.HandleFunc("/api/weather", wc.WeatherHandler)

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return m
}
