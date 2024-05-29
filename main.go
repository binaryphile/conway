package main

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"strings"
)

const (
	screenWidth  = 800
	screenHeight = 600
	cellSize     = 10
)

func main() {
	game := NewGame(screenWidth/cellSize, screenHeight/cellSize)
	game.Initialize(10, 10, StateFromString(heredoc.Doc(`
		_##_
		_#_#
		_#_
	`)))

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")
	ebiten.SetTPS(2)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func StateFromString(s string) [][]bool {
	lines := strings.Split(s, "\n")
	numRows := len(lines) - 1
	numCols := len(lines[0])

	state := make([][]bool, numCols)
	for i := range state {
		state[i] = make([]bool, numRows)
	}

	for i := 0; i < numRows; i++ {
		for j, char := range lines[i] {
			state[j][i] = char == '#'
		}
	}

	return state
}
