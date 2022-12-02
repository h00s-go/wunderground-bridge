package application

import "net/http"

func (app *Application) NewRouter() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/weatherstation/updateweatherstation.php", app.weatherUpdateHandler)
	m.HandleFunc("/weather", app.weatherHandler)
	m.HandleFunc("/", app.defaultHandler)
	return m
}
