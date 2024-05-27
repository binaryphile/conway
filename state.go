package main

import (
	iter "github.com/binaryphile/iterator"
	"strings"
)

type State [][]bool
type StateIterator = iter.Iterator[State]

func StateFromString(s string) State {
	// Split the input string by newlines
	lines := strings.Split(s, "\n")
	numRows := len(lines) - 1
	numCols := len(lines[0])

	// Create the State matrix with dimensions [numCols][numRows]
	state := make(State, numCols)
	for i := range state {
		state[i] = make([]bool, numRows)
	}

	// Iterate over each line and assign values to the columns of the State matrix
	for i, line := range lines {
		for j, char := range line {
			if char == '#' {
				state[j][i] = true
			} else if char == '_' {
				state[j][i] = false
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

	return next
}

func (s State) Grid() [][]rune {
	grid := make([][]rune, len(s))
	for i := range grid {
		grid[i] = make([]rune, len(s[0]))
	}

	for i := range s {
		for j := range s[i] {
			if s[i][j] {
				grid[i][j] = '#'
			} else {
				grid[i][j] = '_'
			}
		}
	}

	return grid
}
