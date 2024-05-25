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
	newTickerChan TickerChanFactory
}

func NewApp(c AppConfig) App {
	return App{
		termbox:       c.termbox,
		newTickerChan: NewTickerChan,
	}
}

func (a App) Close() {
	a.termbox.Close()
}

func (a App) NewTickerChan(duration time.Duration) (<-chan time.Time, func()) {
	return a.newTickerChan(duration)
}

func (a App) Run(initialState string) {
	ui := userinterface.NewTermboxUI(a.termbox)

	tickerChan, tickerStop := a.newTickerChan(1 * time.Second)
	defer tickerStop()

	done := make(chan struct{})
	go func() {
		ui.WaitForInput()
		close(done)
	}()

	var state State
	nextState := NewStateIterator(initialState, newState)
	for {
		select {
		case <-tickerChan:
			state, nextState = nextState()
			ui.Show(state.Grid())
		case <-done:
			return
		}
	}
}
