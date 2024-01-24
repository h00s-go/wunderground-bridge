package helpers

import (
	"math"
	"strconv"
)

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func StrToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}

func roundFloat(number float64) float64 {
	return math.Round(number*100) / 100
}

func ConvertFahrenheitToCelsius(f float64) float64 {
	return roundFloat((f - 32) * 5 / 9)
}

func ConvertMileToKilometer(mph float64) float64 {
	return roundFloat(mph * 1.609344)
}

func ConvertHGToKPA(hg float64) float64 {
	return roundFloat(hg * 33.8638866667)
}

func ConvertInchToMillimeter(inch float64) float64 {
	return roundFloat(inch * 25.4)
}
