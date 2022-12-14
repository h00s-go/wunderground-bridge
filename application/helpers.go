package application

import (
	"strconv"

	"github.com/shopspring/decimal"
)

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

func strToDecimal(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Zero
	}
	return d.Round(2)
}

func convertFahrenheitToCelsius(f float64) decimal.Decimal {
	return decimal.NewFromFloat((f - 32) * 5 / 9).Round(2)
}

func convertMileToKilometer(mph float64) decimal.Decimal {
	return decimal.NewFromFloat(mph * 1.609344).Round(2)
}

func convertHGToKPA(hg float64) decimal.Decimal {
	return decimal.NewFromFloat(hg * 33.8638866667).Round(2)
}

func convertInchToMillimeter(inch float64) decimal.Decimal {
	return decimal.NewFromFloat(inch * 25.4).Round(2)
}
