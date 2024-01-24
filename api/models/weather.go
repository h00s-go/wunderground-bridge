package models

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/wunderground-bridge/api/helpers"
)

type Weather struct {
	StationID   string    `json:"station_id"`
	Temperature float64   `json:"temperature"`
	DewPoint    float64   `json:"dew_point"`
	Humidity    int       `json:"humidity"`
	Pressure    float64   `json:"pressure"`
	Wind        Wind      `json:"wind"`
	Rain        Rain      `json:"rain"`
	Solar       Solar     `json:"solar"`
	Indoor      Indoor    `json:"indoor"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Wind struct {
	Chill     float64 `json:"chill"`
	Direction int     `json:"direction"`
	Speed     float64 `json:"speed"`
	Gust      float64 `json:"gust"`
}

type Rain struct {
	In        float64 `json:"in"`
	InDaily   float64 `json:"in_daily"`
	InWeekly  float64 `json:"in_weekly"`
	InMonthly float64 `json:"in_monthly"`
	InYearly  float64 `json:"in_yearly"`
}

type Solar struct {
	Radiation float64 `json:"radiation"`
	UV        int     `json:"uv"`
}

type Indoor struct {
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
}

func NewWeather(ctx *fiber.Ctx) (*Weather, error) {
	updatedAt, err := time.Parse("2006-01-02 15:04:05", ctx.Query("dateutc"))
	if err != nil {
		return nil, err
	}

	w := &Weather{
		StationID:   ctx.Query("ID"),
		Temperature: helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(ctx.Query("tempf"))),
		DewPoint:    helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(ctx.Query("dewptf"))),
		Humidity:    helpers.StrToInt(ctx.Query("humidity")),
		Pressure:    helpers.ConvertHGToKPA(helpers.StrToFloat(ctx.Query("baromin"))),
		Wind: Wind{
			Chill:     helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(ctx.Query("windchillf"))),
			Direction: helpers.StrToInt(ctx.Query("winddir")),
			Speed:     helpers.ConvertMileToKilometer(helpers.StrToFloat(ctx.Query("windspeedmph"))),
			Gust:      helpers.ConvertMileToKilometer(helpers.StrToFloat(ctx.Query("windgustmph"))),
		},
		Rain: Rain{
			In:        helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("rainin"))),
			InDaily:   helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("dailyrainin"))),
			InWeekly:  helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("weeklyrainin"))),
			InMonthly: helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("monthlyrainin"))),
			InYearly:  helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("yearlyrainin"))),
		},
		Solar: Solar{
			Radiation: helpers.StrToFloat(ctx.Query("solarradiation")),
			UV:        helpers.StrToInt(ctx.Query("UV")),
		},
		Indoor: Indoor{
			Temperature: helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(ctx.Query("indoortempf"))),
			Humidity:    helpers.StrToInt(ctx.Query("indoorhumidity")),
		},
		UpdatedAt: updatedAt,
	}

	if !w.Validate() {
		return nil, errors.New("invalid weather data")
	}
	return w, nil
}

func (w *Weather) Validate() bool {
	if w.Temperature < -50 || w.Temperature > 50 {
		return false
	}
	if w.Humidity < 0 || w.Humidity > 100 {
		return false
	}
	if w.DewPoint < -50 || w.DewPoint > 50 {
		return false
	}
	if w.Pressure < 800 || w.Pressure > 1200 {
		return false
	}
	return true
}
