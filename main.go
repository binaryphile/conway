package main

import (
	"github.com/MakeNowJust/heredoc"
	m "github.com/binaryphile/conway/must"
	"github.com/binaryphile/conway/userinterface"
	"time"
)

func main() {
	app := NewApp(AppConfig{
		termbox: NewTermboxAdapter(),
	})
	defer app.Close()
	Run(app, heredoc.Doc(`
		#_#
		___
		_#_
	`))
}

func NewTermboxAdapter() userinterface.Adapter {
	termbox := userinterface.Adapter{}
	err := termbox.Init()
	m.AssertNil(err)

	termbox.SetInputMode(userinterface.InputEsc)
	err = termbox.Clear(userinterface.ColorDefault, userinterface.ColorDefault)
	m.AssertNil(err)

	return termbox
}

func Run(app App, initialState string) {
	ui := userinterface.NewTermboxUI(app.termbox)

	tickerChan, tickerStop := app.tickerChanFactory(1 * time.Second)
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

func newState(s State) State {
	return s
}
