package userinterface

import "github.com/nsf/termbox-go"

type Adapter struct{}

func (a Adapter) Clear(fg, bg termbox.Attribute) error {
	return termbox.Clear(fg, bg)
}

func (a Adapter) Close() {
	termbox.Close()
}

func (a Adapter) Flush() error {
	return termbox.Flush()
}

func (a Adapter) Init() error {
	return termbox.Init()
}

func (a Adapter) PollEvent() termbox.Event {
	return termbox.PollEvent()
}

func (a Adapter) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
}

func (a Adapter) SetInputMode(mode termbox.InputMode) termbox.InputMode {
	return termbox.SetInputMode(mode)
}
