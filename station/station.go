package station

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/h00s-go/wunderground-bridge/api/helpers"
	"github.com/h00s-go/wunderground-bridge/api/models"
	"github.com/h00s-go/wunderground-bridge/config"
	"github.com/h00s-go/wunderground-bridge/mqtt"
)

type Station struct {
	config   *config.Station
	logger   *log.Logger
	Watchdog Watchdog
	Weather  *models.Weather
}

type Watchdog struct {
	SuccessfulLastUpdate bool
	FailedAttempts       int
}

func NewStation(config *config.Station, logger *log.Logger) *Station {
	return &Station{
		config:   config,
		logger:   logger,
		Watchdog: Watchdog{},
		Weather:  &models.Weather{},
	}
}

func (s *Station) NewWeather(ctx *fiber.Ctx) error {
	if ctx.Query("ID") != s.config.ID {
		return errors.New("Station ID does not match")
	}
	if ctx.Query("PASSWORD") != s.config.Password {
		return errors.New("Station password does not match")
	}
	updatedAt, err := time.Parse("2006-01-02 15:04:05", ctx.Query("dateutc"))
	if err != nil {
		return err
	}

	w := &models.Weather{
		StationID:   ctx.Query("ID"),
		Temperature: helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(ctx.Query("tempf"))),
		DewPoint:    helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(ctx.Query("dewptf"))),
		Humidity:    helpers.StrToInt(ctx.Query("humidity")),
		Pressure:    helpers.ConvertHGToKPA(helpers.StrToFloat(ctx.Query("baromin"))),
		Wind: models.Wind{
			Chill:     helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(ctx.Query("windchillf"))),
			Direction: helpers.StrToInt(ctx.Query("winddir")),
			Speed:     helpers.ConvertMileToKilometer(helpers.StrToFloat(ctx.Query("windspeedmph"))),
			Gust:      helpers.ConvertMileToKilometer(helpers.StrToFloat(ctx.Query("windgustmph"))),
		},
		Rain: models.Rain{
			In:        helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("rainin"))),
			InDaily:   helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("dailyrainin"))),
			InWeekly:  helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("weeklyrainin"))),
			InMonthly: helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("monthlyrainin"))),
			InYearly:  helpers.ConvertInchToMillimeter(helpers.StrToFloat(ctx.Query("yearlyrainin"))),
		},
		Solar: models.Solar{
			Radiation: helpers.StrToDecimal(ctx.Query("solarradiation")),
			UV:        helpers.StrToInt(ctx.Query("UV")),
		},
		Indoor: models.Indoor{
			Temperature: helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(ctx.Query("indoortempf"))),
			Humidity:    helpers.StrToInt(ctx.Query("indoorhumidity")),
		},
		UpdatedAt: updatedAt,
	}

	validWeather := w.Validate()
	s.UpdateWatchDog(validWeather)
	if !validWeather {
		return errors.New("invalid weather data")
	}

	s.Weather = w
	return nil
}

func (s *Station) UpdateWatchDog(success bool) {
	if s.config.WatchdogEnabled {
		if success {
			s.Watchdog.SuccessfulLastUpdate = true
			s.Watchdog.FailedAttempts = 0
			return
		}
		if s.Watchdog.FailedAttempts >= s.config.RebootOnFailedAttempts {
			s.logger.Println("Attempting reboot")
			go s.attemptReboot()
			s.Watchdog.FailedAttempts = 0
		}
		s.Watchdog.SuccessfulLastUpdate = false
		s.Watchdog.FailedAttempts++
		s.logger.Printf("Failed %v time(s) to update weather data", s.Watchdog.FailedAttempts)
	}
}

func (s *Station) attemptReboot() {
	_, err := http.Get(fmt.Sprintf("%s/msgreboot.htm", s.config.URL))
	if err != nil {
		s.logger.Printf("Error while reboot attempt: %v\n", err)
	}
}

func (s *Station) PublishWeatherToMQTT(mqtt *mqtt.MQTT) {
	if mqtt.Config.Enabled {
		w, err := json.Marshal(s.Weather)
		if err == nil {
			mqtt.Publish(s.config.MQTTUpdateTopic, string(w))
		}
	}
}

func (s *Station) UpdateWunderground(w *Wunderground, query string) error {
	if w.Config.Enabled {
		url := fmt.Sprintf("%v?%v", w.Config.UpdateURL, query)
		_, err := http.Get(url)
		return err
	}
	return nil
}
