package application

import (
	"errors"
	"net/url"
	"time"

	"github.com/shopspring/decimal"
)

type Weather struct {
	StationID   string          `json:"station_id"`
	Temperature decimal.Decimal `json:"temperature"`
	DewPoint    decimal.Decimal `json:"dew_point"`
	Humidity    int             `json:"humidity"`
	Pressure    decimal.Decimal `json:"pressure"`
	Wind        Wind            `json:"wind"`
	Rain        Rain            `json:"rain"`
	Solar       Solar           `json:"solar"`
	Indoor      Indoor          `json:"indoor"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type Wind struct {
	Chill     decimal.Decimal `json:"chill"`
	Direction int             `json:"direction"`
	Speed     decimal.Decimal `json:"speed"`
	Gust      decimal.Decimal `json:"gust"`
}

type Rain struct {
	In        decimal.Decimal `json:"in"`
	InDaily   decimal.Decimal `json:"in_daily"`
	InWeekly  decimal.Decimal `json:"in_weekly"`
	InMonthly decimal.Decimal `json:"in_monthly"`
	InYearly  decimal.Decimal `json:"in_yearly"`
}

type Solar struct {
	Radiation decimal.Decimal `json:"radiation"`
	UV        int             `json:"uv"`
}

type Indoor struct {
	Temperature decimal.Decimal `json:"temperature"`
	Humidity    int             `json:"humidity"`
}

func NewWeatherFromStation(data string) (*Weather, error) {
	d, err := url.ParseQuery(data)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse("2006-01-02 15:04:05", d.Get("dateutc"))
	if err != nil {
		return nil, err
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
		return nil, errors.New("invalid weather data")
	}
	return w, nil
}

func (w *Weather) validate() bool {
	if w.Temperature.IntPart() < -50 || w.Temperature.IntPart() > 50 {
		return false
	}
	if w.Humidity < 0 || w.Humidity > 100 {
		return false
	}
	if w.DewPoint.IntPart() < -50 || w.DewPoint.IntPart() > 50 {
		return false
	}
	if w.Pressure.IntPart() < 800 || w.Pressure.IntPart() > 1200 {
		return false
	}
	return true
}
