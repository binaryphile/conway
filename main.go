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

	output := generateOutput()

	show(output)

	waitForInput()
}

func generateOutput() string {
	output := `# #
   
 # `

	return output
}

func gridFromString(s string) [][]rune {
	// Split the input into individual lines
	lines := strings.Split(s, "\n")

	// Calculate the number of rows and columns
	rows := len(lines)
	columns := len(lines[0])

	// Pre-allocate a two-dimensional slice of runes
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, columns) // Pre-allocate each row
	}

	// Populate the two-dimensional slice of runes
	for i, line := range lines {
		runeRow := []rune(line) // Convert each line to a slice of runes
		copy(grid[i], runeRow)  // Copy the slice of runes to the pre-allocated row
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

	for i := range grid {
		for j := range grid[i] {
			termbox.SetCell(i, j, grid[i][j], fg, bg)
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
