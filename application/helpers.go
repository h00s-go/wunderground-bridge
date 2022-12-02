package application

import "strconv"

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func strToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}

func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func convertMileToKilometer(mph float64) float64 {
	return mph * 1.609344
}

func convertHGToKPA(hg float64) float64 {
	return hg * 33.8638866667
}

// function to convert inch to milimeter
func convertInchToMillimeter(inch float64) float64 {
	return inch * 25.4
}
