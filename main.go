package main

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/binaryphile/conway/userinterface"
	m "github.com/binaryphile/must"
)

func main() {
	app := NewApp()
	defer app.Close()

	app.Run(StateFromString(30, heredoc.Doc(`
		_##
		##_
		_#_
	`)))
}

func NewTermboxAdapter() userinterface.TermboxAdapter {
	termbox := userinterface.TermboxAdapter{}
	err := termbox.Init()
	m.AssertNil(err)

	termbox.SetInputMode(userinterface.InputEsc)
	err = termbox.Clear(userinterface.ColorDefault, userinterface.ColorDefault)
	m.AssertNil(err)

	return termbox
}

func Evolve(s State) State {
	state := make([][]int, len(s))

	length := len(s[0])
	for i := range s {
		state[i] = make([]int, length)
	}

	for i := range s {
		for j, cell := range s[i] {
			if cell == 1 {
				neighborCount := s.LiveNeighborCount(i, j)
				if neighborCount == 2 || neighborCount == 3 {
					state[i][j] = 1
				}
			} else {
				if s.LiveNeighborCount(i, j) == 3 {
					state[i][j] = 1
				}
			}
		}
	}

	return state
}
