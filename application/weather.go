package application

import (
	"net/url"
	"strconv"
)

type Weather struct {
	Temperature       float64 `json:"temperature"`
	Humidity          int     `json:"humidity"`
	DewPoint          float64 `json:"dewpoint"`
	WindChill         float64 `json:"windchill"`
	WindDirection     int     `json:"winddirection"`
	WindSpeed         float64 `json:"windspeed"`
	WindGust          float64 `json:"windgust"`
	RainIn            float64 `json:"rainin"`
	RainInDaily       float64 `json:"dailyrainin"`
	RainInWeekly      float64 `json:"weeklyrainin"`
	RainInMonthly     float64 `json:"monthlyrainin"`
	RainInYearly      float64 `json:"yearlyrainin"`
	SolarRadiation    float64 `json:"solarradiastion"`
	UV                int     `json:"uv"`
	IndoorTemperature float64 `json:"indoortemp"`
	IndoorHumidity    int     `json:"indoorhumidity"`
	Pressure          float64 `json:"pressure"`
}

func NewWeather(data string) (*Weather, error) {
	d, err := url.ParseQuery(data)
	if err != nil {
		return nil, err
	}

	return &Weather{
		Temperature:       fahrenheitToCelsius(toFloat(d.Get("tempf"))),
		Humidity:          toInt(d.Get("humidity")),
		DewPoint:          fahrenheitToCelsius(toFloat(d.Get("dewptf"))),
		WindChill:         fahrenheitToCelsius(toFloat(d.Get("windchillf"))),
		WindDirection:     toInt(d.Get("winddir")),
		WindSpeed:         mphToKph(toFloat(d.Get("windspeedmph"))),
		WindGust:          mphToKph(toFloat(d.Get("windgustmph"))),
		RainIn:            toFloat(d.Get("rainin")),
		RainInDaily:       toFloat(d.Get("dailyrainin")),
		RainInWeekly:      toFloat(d.Get("weeklyrainin")),
		RainInMonthly:     toFloat(d.Get("monthlyrainin")),
		RainInYearly:      toFloat(d.Get("yearlyrainin")),
		UV:                toInt(d.Get("UV")),
		IndoorTemperature: fahrenheitToCelsius(toFloat(d.Get("indoortempf"))),
		IndoorHumidity:    toInt(d.Get("indoorhumidity")),
		Pressure:          hgToKpa(toFloat(d.Get("baromin"))),
	}, nil
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func toFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func mphToKph(mph float64) float64 {
	return mph * 1.609344
}

func hgToKpa(hg float64) float64 {
	return hg * 33.8638866667
}
