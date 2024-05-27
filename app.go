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
	tickerChan, stopTicker := a.NewTickerChan(500 * time.Millisecond)
	defer stopTicker()

	done := make(chan struct{})
	go func() {
		a.ui.WaitForInput()
		close(done)
	}()

	var state State
	nextState := NewStateIterator(initialState, Evolve)
	a.ui.Show(initialState.Grid())
	state, nextState = nextState()
	for {
		select {
		case <-tickerChan:
			a.ui.Show(state.Grid())
			state, nextState = nextState()
		case <-done:
			return
		}
	}
}
