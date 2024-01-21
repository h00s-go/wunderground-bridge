package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/wunderground-bridge/api/services"
)

type WeatherController struct {
	services *services.Services
}

func NewWeatherController(services *services.Services) *WeatherController {
	return &WeatherController{
		services: services,
	}
}

func (wc *WeatherController) GetWeatherHandler(ctx *fiber.Ctx) error {
	weather, err := json.MarshalIndent(wc.services.Station.Weather, "", "\t")
	if err == nil {
		ctx.Status(http.StatusOK)
		return ctx.Send(weather)
	}
	ctx.Status(http.StatusInternalServerError)
	return err
}

func (wc *WeatherController) GetWeatherUpdateHandler(ctx *fiber.Ctx) error {
	/*c, err := wc.services.WebsocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		wc.services.Logger.Print("Error while upgrading:", err)
		return
	}
	defer c.Close()
	for {
		weather, err := json.Marshal(wc.services.Station.Weather)
		if err == nil {
			err = c.WriteMessage(websocket.TextMessage, weather)
			if err != nil {
				wc.services.Logger.Print("Error while writing to websocket:", err)
				return
			}
		}
	}*/
	return nil
}
