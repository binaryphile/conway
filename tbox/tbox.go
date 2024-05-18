package tbox

import (
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

type Tbox struct {
	termbox Termbox
}

func NewTbox(termbox Termbox) Tbox {
	return Tbox{
		termbox: termbox,
	}
}

func (t Tbox) Show(output string) {
	t.showGrid(gridFromString(output))
}

func (t Tbox) WaitForInput() {
	for {
		if event := t.termbox.PollEvent(); event.Type == EventKey {
			break
		}
	}
}

func AssertNil(err error) {
	if err != nil {
		panic(err)
	}
}

func (t Tbox) showGrid(grid [][]rune) {
	fg := ColorDefault
	bg := ColorDefault

	for rowNum := range grid {
		for colNum := range grid[rowNum] {
			termbox.SetCell(colNum, rowNum, grid[rowNum][colNum], fg, bg) // SetCell is col, row
		}
	}

	err := termbox.Flush()
	AssertNil(err)
}
