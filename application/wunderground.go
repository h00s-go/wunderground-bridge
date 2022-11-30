package application

import (
	"fmt"
	"net/http"
)

func (app *Application) updateWunderground(query string) {
	if app.config.UpdateWunderground {
		url := fmt.Sprintf("%v?%v", app.config.WundergroundUpdateURL, query)
		_, err := http.Get(url)
		if err != nil {
			app.logger.Printf("Error: %v\n", err)
		}
	}
}
