package main

import (
	"github.com/binaryphile/conway/tbox"
)

func main() {
	termbox := newTermbox()
	defer termbox.Close()
	run(termbox)
}

func generateOutput() string {
	return `
#_#
___
_#_`
}

func newTermbox() tbox.Adapter {
	termbox := tbox.Adapter{}
	err := termbox.Init()
	tbox.AssertNil(err)

	termbox.SetInputMode(tbox.InputEsc)
	err = termbox.Clear(tbox.ColorDefault, tbox.ColorDefault)
	tbox.AssertNil(err)

	return termbox
}

func run(termbox tbox.Termbox) {
	termbox.Show(generateOutput())
	termbox.WaitForInput()
}
