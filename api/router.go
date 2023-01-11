package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/wunderground-bridge/api/controllers"
)

func (api *API) SetRoutes() {
	wc := controllers.NewWeatherController(api.services)
	sc := controllers.NewStationController(api.services)

	api.server.Get("/weatherstation/updateweatherstation.php", sc.GetStationUpdateHandler)

	api.server.Get("/api/weather/update", wc.GetWeatherUpdateHandler)
	api.server.Get("/api/weather", wc.GetWeatherHandler)

	api.server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})
}
