package main

import "time"

type TickerChanFactory = func(time.Duration) (<-chan time.Time, func())

func NewTickerChan(duration time.Duration) (<-chan time.Time, func()) {
	ticker := time.NewTicker(duration)
	return ticker.C, ticker.Stop
}
