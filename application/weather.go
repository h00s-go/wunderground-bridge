package application

import (
	"net/url"
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
	RainInDaily       float64 `json:"rainindaily"`
	RainInWeekly      float64 `json:"raininweekly"`
	RainInMonthly     float64 `json:"raininmonthly"`
	RainInYearly      float64 `json:"raininyearly"`
	SolarRadiation    float64 `json:"solarradiation"`
	UV                int     `json:"uv"`
	IndoorTemperature float64 `json:"indoortemperature"`
	IndoorHumidity    int     `json:"indoorhumidity"`
	Pressure          float64 `json:"pressure"`
}

func NewWeatherFromStation(data string) (*Weather, error) {
	d, err := url.ParseQuery(data)
	if err != nil {
		return nil, err
	}

	return &Weather{
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
