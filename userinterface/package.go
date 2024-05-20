package userinterface

import (
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	ColorDefault = termbox.ColorDefault
	EventKey     = termbox.EventKey
	InputEsc     = termbox.InputEsc
)

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
