package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"strings"
)

const (
	width    = 140
	height   = 90
	cellSize = 10
)

func main() {
	ebiten.SetWindowSize(width*cellSize, height*cellSize)

	game := NewGame(width, height)

	initFigure := rPentomino
	figureWidth, figureHeight := SizeFromString(initFigure)

	leftOffset := int(float64(width-figureWidth) / 2)
	topOffset := int(float64(height-figureHeight) / 2)
	game.Initialize(leftOffset, topOffset, StateFromString(initFigure))

	ebiten.SetWindowTitle("Conway's Game of Life")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func SizeFromString(s string) (width, height int) {
	lines := strings.Split(s, "\n")
	return len(strings.TrimSpace(lines[1])), len(lines) - 2
}

func StateFromString(s string) [][]bool {
	lines := strings.Split(s, "\n")
	lines = lines[1 : len(lines)-1]
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	state := make([][]bool, len(lines[0]))
	for i := range lines[0] {
		state[i] = make([]bool, len(lines))
	}

	for i := range lines {
		for j := range lines[i] {
			state[j][i] = lines[i][j] == '#'
		}
	}

	return state
}
