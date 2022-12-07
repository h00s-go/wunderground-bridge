package application

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/h00s-go/wunderground-bridge/config"
)

type Station struct {
	config   *config.Station
	logger   *log.Logger
	Watchdog Watchdog
	Weather  *Weather
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
		Weather:  &Weather{},
	}
}

func (s *Station) NewWeather(r *http.Request) error {
	d, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return err
	}
	updatedAt, err := time.Parse("2006-01-02 15:04:05", d.Get("dateutc"))
	if err != nil {
		return err
	}

	w := &Weather{
		StationID:   d.Get("ID"),
		Temperature: convertFahrenheitToCelsius(strToFloat(d.Get("tempf"))),
		DewPoint:    convertFahrenheitToCelsius(strToFloat(d.Get("dewptf"))),
		Humidity:    strToInt(d.Get("humidity")),
		Pressure:    convertHGToKPA(strToFloat(d.Get("baromin"))),
		Wind: Wind{
			Chill:     convertFahrenheitToCelsius(strToFloat(d.Get("windchillf"))),
			Direction: strToInt(d.Get("winddir")),
			Speed:     convertMileToKilometer(strToFloat(d.Get("windspeedmph"))),
			Gust:      convertMileToKilometer(strToFloat(d.Get("windgustmph"))),
		},
		Rain: Rain{
			In:        convertInchToMillimeter(strToFloat(d.Get("rainin"))),
			InDaily:   convertInchToMillimeter(strToFloat(d.Get("dailyrainin"))),
			InWeekly:  convertInchToMillimeter(strToFloat(d.Get("weeklyrainin"))),
			InMonthly: convertInchToMillimeter(strToFloat(d.Get("monthlyrainin"))),
			InYearly:  convertInchToMillimeter(strToFloat(d.Get("yearlyrainin"))),
		},
		Solar: Solar{
			Radiation: strToDecimal(d.Get("solarradiation")),
			UV:        strToInt(d.Get("UV")),
		},
		Indoor: Indoor{
			Temperature: convertFahrenheitToCelsius(strToFloat(d.Get("indoortempf"))),
			Humidity:    strToInt(d.Get("indoorhumidity")),
		},
		UpdatedAt: updatedAt,
	}

	if !w.validate() {
		s.UpdateWatchDog(false)
		return errors.New("invalid weather data")
	}
	s.Weather = w
	s.UpdateWatchDog(true)
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
