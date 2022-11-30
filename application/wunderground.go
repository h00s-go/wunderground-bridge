package application

import (
	"fmt"
	"net/http"
)

func (app *Application) updateWunderground(query string) {
	if app.config.UpdateWunderground {
		url := fmt.Sprintf("http://rtupdate.wunderground.com/weatherstation/updateweatherstation.php?%v", query)
		_, err := http.Get(url)
		if err != nil {
			app.logger.Printf("Error: %v\n", err)
		}
	}
}
