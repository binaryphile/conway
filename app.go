package main

import (
	"github.com/binaryphile/conway/userinterface"
	"time"
)

type App struct {
	newTickerChan TickerChanFactory
	ui            userinterface.TermboxUI
}

func NewApp() App {
	return App{
		newTickerChan: NewTickerChan,
		ui:            userinterface.NewTermboxUI(NewTermboxAdapter()),
	}
}

func (a App) Close() {
	a.ui.Close()
}

func (a App) NewTickerChan(duration time.Duration) (<-chan time.Time, func()) {
	return a.newTickerChan(duration)
}

func (a App) Run(initialState State) {
	tickerChan, stopTicker := a.NewTickerChan(1 * time.Second)
	defer stopTicker()

	done := make(chan struct{})
	go func() {
		a.ui.WaitForInput()
		close(done)
	}()

	var state State
	nextState := NewStateIterator(initialState, Evolve)
	for {
		select {
		case <-tickerChan:
			state, nextState = nextState()
			a.ui.Show(state.Grid())
		case <-done:
			return
		}
	}
}
