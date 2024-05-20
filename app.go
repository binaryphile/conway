package main

import (
	"github.com/binaryphile/conway/userinterface"
	"time"
)

type AppConfig struct {
	termbox userinterface.Termbox
}

type App struct {
	termbox       userinterface.Termbox
	tickerFactory TickerIfaceFactory
}

func NewApp(c AppConfig) App {
	return App{
		termbox:       c.termbox,
		tickerFactory: NewTickerAdapterFactory(time.NewTicker),
	}
}

func (a App) Close() {
	a.termbox.Close()
}

func (a App) NewTicker(duration time.Duration) TickerAdapter {
	return a.tickerFactory(duration)
}
