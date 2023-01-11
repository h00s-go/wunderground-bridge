package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/wunderground-bridge/api/services"
)

type StationController struct {
	services *services.Services
}

func NewStationController(services *services.Services) *StationController {
	return &StationController{
		services: services,
	}
}

func (sc *StationController) GetStationUpdateHandler(ctx *fiber.Ctx) error {
	if err := sc.services.Station.NewWeather(ctx); err == nil {
		go sc.services.Station.PublishWeatherToMQTT(sc.services.MQTT)
		go sc.services.Station.UpdateWunderground(sc.services.Wunderground, string(ctx.Request().URI().QueryString()))
	} else {
		sc.services.Logger.Println("Error parsing weather: ", err, string(ctx.Request().URI().QueryString()))
	}
	ctx.Status(http.StatusOK)
	return ctx.SendString("success")
}
