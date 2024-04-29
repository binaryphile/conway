package main

import (
	"github.com/nsf/termbox-go"
	"strings"
)

func main() {
	cleanup, err := initialize()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	show(generateOutput())

	waitForInput()
}

func generateOutput() string {
	return `
#_#
___
_#_`
}

func gridFromString(s string) [][]rune {
	// Split the input into individual lines
	lines := strings.Split(s, "\n")

	lines = lines[1:]

	// Pre-allocate a two-dimensional slice of runes
	grid := make([][]rune, len(lines))
	for i := range grid {
		grid[i] = make([]rune, len(lines[0])) // Pre-allocate each row
	}

	// Populate the two-dimensional slice of runes
	for i, line := range lines {
		row := []rune(line) // Convert each line to a slice of runes
		copy(grid[i], row)  // Copy the slice of runes to the pre-allocated row
	}

	return grid
}

func initialize() (_ func(), err error) {
	err = termbox.Init()
	if err != nil {
		return
	}

	termbox.SetInputMode(termbox.InputEsc)
	err = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		return
	}

	return termbox.Close, nil
}

func show(output string) {
	showGrid(gridFromString(output))
}

func showGrid(grid [][]rune) {
	fg := termbox.ColorDefault
	bg := termbox.ColorDefault

	for rowNum := range grid {
		for colNum := range grid[rowNum] {
			termbox.SetCell(colNum, rowNum, grid[rowNum][colNum], fg, bg) // SetCell is col, row
		}
	}

	termbox.Flush()
}

func waitForInput() {
	for {
		if event := termbox.PollEvent(); event.Type == termbox.EventKey {
			break
		}
	}
}
