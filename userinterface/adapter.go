package userinterface

import "github.com/nsf/termbox-go"

type TermboxAdapter struct{}

func (a TermboxAdapter) Clear(fg, bg termbox.Attribute) error {
	return termbox.Clear(fg, bg)
}

func (a TermboxAdapter) Close() {
	termbox.Close()
}

func (a TermboxAdapter) Flush() error {
	return termbox.Flush()
}

func (a TermboxAdapter) Init() error {
	return termbox.Init()
}

func (a TermboxAdapter) PollEvent() termbox.Event {
	return termbox.PollEvent()
}

func (a TermboxAdapter) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
}

func (a TermboxAdapter) SetInputMode(mode termbox.InputMode) termbox.InputMode {
	return termbox.SetInputMode(mode)
}
