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
	switch period {
	case CandlePeriod1m:
		return time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), ts.Minute(), 0, 0, ts.Location()), nil
	case CandlePeriod2m:
		return time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), ((ts.Minute()-1)/2)*2, 0, 0, ts.Location()), nil
	case CandlePeriod10m:
		return time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), ((ts.Minute()-1)/10)*10, 0, 0, ts.Location()), nil
	default:
		return time.Time{}, ErrUnknownPeriod
	}
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
