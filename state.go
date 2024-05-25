package main

import (
	"github.com/binaryphile/conway/iterator"
	"strings"
)

type State [][]bool
type StateIterator = iterator.Iterator[State]

func StateFromString(s string) State {
	// Split the input string by newlines
	lines := strings.Split(s, "\n")
	numRows := len(lines)
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

// Create an iterator that yields the initial state and subsequent states using newState
func NewStateIterator(initialState string, newState func(State) State) StateIterator {
	state := StateFromString(initialState)

	var next StateIterator
	next = func() (State, StateIterator) {
		state = newState(state)
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
			if s[i][j] {
				grid[i][j] = '#'
			} else {
				grid[i][j] = '_'
			}
		}
	}

	return grid
}
