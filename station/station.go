package station

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

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

func (s *Station) NewWeather(r *http.Request) error {
	d, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return err
	}
	if d.Get("ID") != s.config.ID {
		return errors.New("Station ID does not match")
	}
	if d.Get("PASSWORD") != s.config.Password {
		return errors.New("Station password does not match")
	}
	updatedAt, err := time.Parse("2006-01-02 15:04:05", d.Get("dateutc"))
	if err != nil {
		return err
	}

	w := &models.Weather{
		StationID:   d.Get("ID"),
		Temperature: helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(d.Get("tempf"))),
		DewPoint:    helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(d.Get("dewptf"))),
		Humidity:    helpers.StrToInt(d.Get("humidity")),
		Pressure:    helpers.ConvertHGToKPA(helpers.StrToFloat(d.Get("baromin"))),
		Wind: models.Wind{
			Chill:     helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(d.Get("windchillf"))),
			Direction: helpers.StrToInt(d.Get("winddir")),
			Speed:     helpers.ConvertMileToKilometer(helpers.StrToFloat(d.Get("windspeedmph"))),
			Gust:      helpers.ConvertMileToKilometer(helpers.StrToFloat(d.Get("windgustmph"))),
		},
		Rain: models.Rain{
			In:        helpers.ConvertInchToMillimeter(helpers.StrToFloat(d.Get("rainin"))),
			InDaily:   helpers.ConvertInchToMillimeter(helpers.StrToFloat(d.Get("dailyrainin"))),
			InWeekly:  helpers.ConvertInchToMillimeter(helpers.StrToFloat(d.Get("weeklyrainin"))),
			InMonthly: helpers.ConvertInchToMillimeter(helpers.StrToFloat(d.Get("monthlyrainin"))),
			InYearly:  helpers.ConvertInchToMillimeter(helpers.StrToFloat(d.Get("yearlyrainin"))),
		},
		Solar: models.Solar{
			Radiation: helpers.StrToDecimal(d.Get("solarradiation")),
			UV:        helpers.StrToInt(d.Get("UV")),
		},
		Indoor: models.Indoor{
			Temperature: helpers.ConvertFahrenheitToCelsius(helpers.StrToFloat(d.Get("indoortempf"))),
			Humidity:    helpers.StrToInt(d.Get("indoorhumidity")),
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
