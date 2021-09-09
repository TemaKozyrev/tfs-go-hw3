package domain

import (
	"errors"
	"time"
)

type Price struct {
	Ticker string
	Value  float64
	TS     time.Time
}

var ErrUnknownPeriod = errors.New("unknown period")

type CandlePeriod string

const (
	CandlePeriod1m  CandlePeriod = "1m"
	CandlePeriod2m  CandlePeriod = "2m"
	CandlePeriod10m CandlePeriod = "10m"
)

func PeriodTS(period CandlePeriod, ts time.Time) (time.Time, error) {
	var m int

	switch period {
	case CandlePeriod1m:
		m = ts.Minute()
	case CandlePeriod2m:
		m = ((ts.Minute() - 1) / 2) * 2
	case CandlePeriod10m:
		m = ((ts.Minute() - 1) / 10) * 10
	default:
		return time.Time{}, ErrUnknownPeriod
	}

	return time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), m, 0, 0, ts.Location()), nil
}

type Candle struct {
	Ticker string
	Period CandlePeriod // Интервал
	Open   float64      // Цена открытия
	High   float64      // Максимальная цена
	Low    float64      // Минимальная цена
	Close  float64      // Цена закрытие
	TS     time.Time    // Время начала интервала
}
