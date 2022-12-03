package application

import (
	"net/url"
)

type Weather struct {
	StationID         string  `json:"station_id"`
	Temperature       float64 `json:"temperature"`
	Humidity          int     `json:"humidity"`
	DewPoint          float64 `json:"dewpoint"`
	WindChill         float64 `json:"wind_chill"`
	WindDirection     int     `json:"wind_direction"`
	WindSpeed         float64 `json:"wind_speed"`
	WindGust          float64 `json:"wind_gust"`
	RainIn            float64 `json:"rain_in"`
	RainInDaily       float64 `json:"rain_in_daily"`
	RainInWeekly      float64 `json:"rain_in_weekly"`
	RainInMonthly     float64 `json:"rain_in_monthly"`
	RainInYearly      float64 `json:"rain_in_yearly"`
	SolarRadiation    float64 `json:"solar_radiation"`
	UV                int     `json:"uv"`
	IndoorTemperature float64 `json:"indoor_temperature"`
	IndoorHumidity    int     `json:"indoor_humidity"`
	Pressure          float64 `json:"pressure"`
}

func NewWeatherFromStation(data string) (*Weather, error) {
	d, err := url.ParseQuery(data)
	if err != nil {
		return nil, err
	}

	return &Weather{
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
		SolarRadiation:    strToFloat(d.Get("solarradiation")),
		UV:                strToInt(d.Get("UV")),
		IndoorTemperature: convertFahrenheitToCelsius(strToFloat(d.Get("indoortempf"))),
		IndoorHumidity:    strToInt(d.Get("indoorhumidity")),
		Pressure:          convertHGToKPA(strToFloat(d.Get("baromin"))),
	}, nil
}
