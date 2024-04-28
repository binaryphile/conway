package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	err = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	termbox.SetCell(0, 0, 'a', termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()

	time.Sleep(2 * time.Second)
}
