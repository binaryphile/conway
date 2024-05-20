package main

import "time"

type TickerFactory = func(time.Duration) *time.Ticker
type TickerAdapterFactory = func(time.Duration) TickerAdapter
type TickerIfaceFactory = func(time.Duration) TickerIface

type TickerIface interface {
	C() <-chan time.Time
	Stop()
}

type TickerAdapter struct {
	*time.Ticker
}

func NewTickerAdapterFactory(newTicker TickerFactory) TickerAdapterFactory {
	return func(duration time.Duration) TickerAdapter {
		return TickerAdapter{
			Ticker: newTicker(duration),
		}
	}
}

func (t TickerAdapter) C() <-chan time.Time {
	return t.Ticker.C
}
