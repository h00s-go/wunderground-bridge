package application

import (
	"errors"
	"net/url"
	"time"
)

type Station struct {
	Watchdog Watchdog
	Weather  *Weather
}

type Watchdog struct {
	FailedAttempts int
}

func NewStation() *Station {
	return &Station{
		Watchdog: Watchdog{},
		Weather:  &Weather{},
	}
}

func (s *Station) NewWeather(data string) error {
	d, err := url.ParseQuery(data)
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
		return errors.New("invalid weather data")
	}
	s.Weather = w
	return nil
}
