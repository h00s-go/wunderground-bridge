package models

import (
	"fmt"
	"io"
	"net/http"

	"github.com/h00s-go/wunderground-bridge/config"
)

type Wunderground struct {
	Config *config.Wunderground
}

func NewWunderground(config *config.Wunderground) *Wunderground {
	return &Wunderground{
		Config: config,
	}
}

func (w *Wunderground) Update(query string) error {
	if w.Config.Enabled {
		url := fmt.Sprintf("%v?%v", w.Config.UpdateURL, query)
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		_, err = io.ReadAll(response.Body)
		return err
	}
	return nil
}
