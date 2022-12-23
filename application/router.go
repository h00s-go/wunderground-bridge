package application

import "net/http"

func (app *Application) NewRouter() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/weatherstation/updateweatherstation.php", app.stationUpdateHandler)

	m.HandleFunc("/api/weather/update", app.weatherUpdateHandler)
	m.HandleFunc("/api/weather", app.weatherHandler)

	m.HandleFunc("/", app.defaultHandler)
	return m
}
