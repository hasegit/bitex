package domain

import "time"

type MarketData struct {
	Date  time.Time
	Price float64
}
