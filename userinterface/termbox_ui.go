package userinterface

import (
	m "github.com/binaryphile/must"
	"github.com/nsf/termbox-go"
)

type Termbox interface {
	Clear(fg, bg termbox.Attribute) error
	Close()
	Flush() error
	Init() error
	PollEvent() termbox.Event
	SetCell(x, y int, ch rune, fg, bg termbox.Attribute)
	SetInputMode(termbox.InputMode) termbox.InputMode
}

type TermboxUI struct {
	termbox Termbox
}

func NewTermboxUI(termbox Termbox) TermboxUI {
	return TermboxUI{
		termbox: termbox,
	}
}

func (t TermboxUI) WaitForInput() {
	for {
		if event := t.termbox.PollEvent(); event.Type == EventKey {
			break
		}
	}
}

func (t TermboxUI) Show(grid [][]rune) {
	fg := ColorDefault
	bg := ColorDefault

	for rowNum := range grid {
		for colNum := range grid[rowNum] {
			termbox.SetCell(colNum, rowNum, grid[rowNum][colNum], fg, bg) // SetCell is col, row
		}
	}

	err := termbox.Flush()
	m.AssertNil(err)
}
