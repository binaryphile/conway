package main

import (
	e "github.com/hajimehoshi/ebiten/v2"
	u "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const slowdownFactor = 5

type Game struct {
	board           [][]bool
	slowdownCounter int

	height int
	width  int

	paused     bool
	stalePress bool
	started    bool
}

func NewGame(width, height int) *Game {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}

	return &Game{
		board: board,

		paused:     false,
		stalePress: false,
		started:    false,

		height: height,
		width:  width,
	}
}

func (g *Game) Draw(screen *e.Image) {
	screen.Clear()
	for y := range g.board {
		for x := range g.board[y] {
			if g.board[y][x] {
				vector.DrawFilledRect(screen, float32(x*cellSize), float32(y*cellSize), cellSize, cellSize, color.White, false)
			}
		}
	}

	if g.paused {
		u.DebugPrint(screen, "Paused")
	} else if !g.started {
		u.DebugPrint(screen, "Press space to start")
	}
}

func (g *Game) Initialize(leftOffset, topOffset int, state [][]bool) {
	for i := range state {
		i2 := i + leftOffset
		for j := range state[i] {
			j2 := j + topOffset
			g.board[j2][i2] = state[i][j]
		}
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return g.width * cellSize, g.height * cellSize
}

func (g *Game) Update() error {
	if g.slowdownCounter == slowdownFactor {
		g.slowdownCounter = 0
	} else {
		g.slowdownCounter += 1
		return nil
	}

	if e.IsKeyPressed(e.KeySpace) {
		if !g.stalePress {
			if !g.started {
				g.started = true
			} else {
				g.paused = !g.paused
			}
			g.stalePress = true
		}
	} else {
		g.stalePress = false
	}

	if g.paused || !g.started {
		return nil
	}

	next := make([][]bool, g.height)
	for y := range g.board {
		next[y] = make([]bool, g.width)
		for x := range g.board[y] {
			aliveNeighbors := g.countAliveNeighbors(x, y)
			if g.board[y][x] {
				next[y][x] = aliveNeighbors == 2 || aliveNeighbors == 3
			} else {
				next[y][x] = aliveNeighbors == 3
			}
		}
	}
	g.board = next

	return nil
}

func (g *Game) countAliveNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			x2, y2 := x+dx, y+dy
			if x2 >= 0 && x2 < g.width && y2 >= 0 && y2 < g.height {
				if g.board[y2][x2] {
					count++
				}
			}
		}
	}
	return count
}
