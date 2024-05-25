package main

import (
	"github.com/binaryphile/conway/userinterface"
	"time"
)

type AppConfig struct {
	termbox userinterface.Termbox
}

type App struct {
	termbox           userinterface.Termbox
	tickerChanFactory func(time.Duration) (<-chan time.Time, func())
}

func NewApp(c AppConfig) App {
	return App{
		termbox:           c.termbox,
		tickerChanFactory: NewTickerChan,
	}
}

func (a App) Close() {
	a.termbox.Close()
}

func (a App) NewTickerChan(duration time.Duration) (<-chan time.Time, func()) {
	return a.tickerChanFactory(duration)
}
