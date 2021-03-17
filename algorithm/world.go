package algorithm

import (
	"fmt"
	"math/rand"
)

// World represent a configuration of GOF
type World struct {
	board  [][]bool
	height int
	width  int
}

// NewWorld ..
func NewWorld(height, width int) World {
	world := make([][]bool, height)
	for i := range world {
		world[i] = make([]bool, width)
		for j := range world[i] {
			// Seed init state
			// TODO: improve random init state
			if rand.Int()%3 == 0 {
				world[i][j] = true
			} else {
				world[i][j] = false
			}
		}
	}
	return World{world, height, width}
}

// isAlive check if a specific cell is alive
func (w *World) isAlive(x, y int) bool {
	x += w.width
	x %= w.width
	y += w.height
	y %= w.height
	return w.board[y][x]
}

// willAliveNextGen ..
func (w *World) willAliveNextGen(x, y int) bool {
	var count int8
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if w.isAlive(x+j, y+i) {
				count++
			}
		}
	}
	if count == 3 || (w.board[y][x] && count == 2) {
		return true
	}
	return false
}

func (w *World) Get() ([][]bool, int, int) {
	return w.board, w.height, w.width
}

// Next generate next configuration
func (w *World) Next() {
	newBoard := make([][]bool, w.height)

	for i := 0; i < w.height; i++ { // y
		newBoard[i] = make([]bool, w.width)
		for j := 0; j < w.width; j++ { // x
			newBoard[i][j] = w.willAliveNextGen(j, i)
		}
	}
	w.board = newBoard
}

// Print World into std.out
func (w *World) Print() {
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			if w.board[i][j] {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
