package models

import (
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

func (w *Weather) Validate() bool {
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
