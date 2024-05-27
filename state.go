package main

import (
	iter "github.com/binaryphile/iterator"
	"strings"
)

type State [][]int
type StateIterator = iter.Iterator[State]

func StateFromString(borderSize int, s string) State {
	// Split the input string by newlines
	lines := strings.Split(s, "\n")
	numRows := len(lines) - 1
	numCols := len(lines[0])

	state := make(State, numCols+2*borderSize)
	for i := range state {
		state[i] = make([]int, numRows+2*borderSize)
	}

	// Iterate over each line and assign values to the columns of the State matrix
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				state[j+borderSize][i+borderSize] = 1
			} else if char == '_' {
				state[j+borderSize][i+borderSize] = 0
			}
		}
	}

	return state
}

// NewStateIterator returns an iterator that yields the initial state and subsequent states using newState.
func NewStateIterator(state State, evolve func(State) State) StateIterator {
	var next StateIterator
	next = func() (State, StateIterator) {
		state = evolve(state)
		return state, next
	}

	return func() (State, StateIterator) {
		return state, next
	}
}

func (s State) Grid() [][]rune {
	grid := make([][]rune, len(s))
	for i := range grid {
		grid[i] = make([]rune, len(s[0]))
	}

	for i := range s {
		for j := range s[i] {
			if s[i][j] == 1 {
				grid[i][j] = '#'
			} else {
				grid[i][j] = ' '
			}
		}
	}

	return grid
}

func (s State) LiveNeighborCount(x, y int) int {
	count := 0

	for i := -1; i < 2; i++ {
		x2 := x + i
		if x2 < 0 || x2 == len(s) {
			continue
		}

		for j := -1; j < 2; j++ {
			y2 := y + j
			if y2 < 0 || y2 == len(s[x2]) {
				continue
			}

			if !(i == 0 && j == 0) {
				count += s[x2][y2]
			}
		}
	}

	return count
}
