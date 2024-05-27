package main

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/binaryphile/conway/userinterface"
	m "github.com/binaryphile/must"
)

func main() {
	app := NewApp()
	defer app.Close()

	app.Run(StateFromString(heredoc.Doc(`
		#_#
		___
		_#_
	`)))
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

func Evolve(s State) State {
	return s
}
