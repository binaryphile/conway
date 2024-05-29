package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
	"time"
)

type Game struct {
	board   [][]bool
	paused  bool
	started bool

	debounceInterval time.Duration
	lastSpace        time.Time

	height int
	width  int
}

func NewGame(width, height int) *Game {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}

	return &Game{
		board:   board,
		paused:  false,
		started: false,

		height: height,
		width:  width,

		debounceInterval: 100 * time.Millisecond,
		lastSpace:        time.Now(),
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for y := range g.board {
		for x := range g.board[y] {
			if g.board[y][x] {
				vector.DrawFilledRect(screen, float32(x*cellSize), float32(y*cellSize), cellSize, cellSize, color.White, false)
			}
		}
	}

	if g.paused {
		ebitenutil.DebugPrint(screen, "Paused")
	} else if !g.started {
		ebitenutil.DebugPrint(screen, "Press space to start")
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return g.width * cellSize, g.height * cellSize
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		now := time.Now()
		if now.Sub(g.lastSpace) > g.debounceInterval {
			g.lastSpace = now
			if !g.started {
				g.started = true
			} else {
				g.paused = !g.paused
			}
		}
	}

	if g.paused || !g.started {
		return nil
	}

	next := make([][]bool, g.height)
	for i := range next {
		next[i] = make([]bool, g.width)
	}

	for y := range g.board {
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
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < g.width && ny >= 0 && ny < g.height {
				if g.board[ny][nx] {
					count++
				}
			}
		}
	}
	return count
}

func (g *Game) Initialize(leftOffset, topOffset int, state [][]bool) {
	for i := range state {
		translatedI := i + leftOffset
		for j := range state[i] {
			translatedJ := j + topOffset
			g.board[translatedJ][translatedI] = state[i][j]
		}
	}
}

func (g *Game) randomize() {
	for y := range g.board {
		for x := range g.board[y] {
			g.board[y][x] = rand.Intn(2) == 0
		}
	}
}
