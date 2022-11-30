package application

import (
	"net/http"
)

func (app *Application) weatherUpdateHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Println(r.URL.RawQuery)
	w.WriteHeader(http.StatusOK)
	go app.updateWunderground(r.URL.RawQuery)
}

func (app *Application) defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
