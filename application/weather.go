package application

import (
	"net/url"
	"time"

	"github.com/shopspring/decimal"
)

type Weather struct {
	StationID         string          `json:"station_id"`
	Temperature       decimal.Decimal `json:"temperature"`
	Humidity          int             `json:"humidity"`
	DewPoint          decimal.Decimal `json:"dew_point"`
	WindChill         decimal.Decimal `json:"wind_chill"`
	WindDirection     int             `json:"wind_direction"`
	WindSpeed         decimal.Decimal `json:"wind_speed"`
	WindGust          decimal.Decimal `json:"wind_gust"`
	RainIn            decimal.Decimal `json:"rain_in"`
	RainInDaily       decimal.Decimal `json:"rain_in_daily"`
	RainInWeekly      decimal.Decimal `json:"rain_in_weekly"`
	RainInMonthly     decimal.Decimal `json:"rain_in_monthly"`
	RainInYearly      decimal.Decimal `json:"rain_in_yearly"`
	SolarRadiation    decimal.Decimal `json:"solar_radiation"`
	UV                int             `json:"uv"`
	IndoorTemperature decimal.Decimal `json:"indoor_temperature"`
	IndoorHumidity    int             `json:"indoor_humidity"`
	Pressure          decimal.Decimal `json:"pressure"`
	UpdatedAt         time.Time       `json:"updated_at"`
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
		StationID:         d.Get("ID"),
		Temperature:       convertFahrenheitToCelsius(strToFloat(d.Get("tempf"))),
		Humidity:          strToInt(d.Get("humidity")),
		DewPoint:          convertFahrenheitToCelsius(strToFloat(d.Get("dewptf"))),
		WindChill:         convertFahrenheitToCelsius(strToFloat(d.Get("windchillf"))),
		WindDirection:     strToInt(d.Get("winddir")),
		WindSpeed:         convertMileToKilometer(strToFloat(d.Get("windspeedmph"))),
		WindGust:          convertMileToKilometer(strToFloat(d.Get("windgustmph"))),
		RainIn:            convertInchToMillimeter(strToFloat(d.Get("rainin"))),
		RainInDaily:       convertInchToMillimeter(strToFloat(d.Get("dailyrainin"))),
		RainInWeekly:      convertInchToMillimeter(strToFloat(d.Get("weeklyrainin"))),
		RainInMonthly:     convertInchToMillimeter(strToFloat(d.Get("monthlyrainin"))),
		RainInYearly:      convertInchToMillimeter(strToFloat(d.Get("yearlyrainin"))),
		SolarRadiation:    strToDecimal(d.Get("solarradiation")),
		UV:                strToInt(d.Get("UV")),
		IndoorTemperature: convertFahrenheitToCelsius(strToFloat(d.Get("indoortempf"))),
		IndoorHumidity:    strToInt(d.Get("indoorhumidity")),
		Pressure:          convertHGToKPA(strToFloat(d.Get("baromin"))),
		UpdatedAt:         updatedAt,
	}

	/*if w.validate() {
		return nil, errors.New("invalid weather data")
	}*/
	return w, nil
}

/*
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
*/
