package algorithm

import (
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
			// TODO: dehardcode random function
			if rand.Int()%3 == 0 {
				world[i][j] = true
			} else {
				world[i][j] = false
			}
		}
	}
	return World{world, height, width}
}

// IsAlive check if a specific cell is alive
func (w World) IsAlive(x, y int) bool {
	x += w.width
	x %= w.width
	y += w.height
	y %= w.height
	return w.board[y][x]
}

// IsAliveNext ..
func (w World) IsAliveNext(x, y int) bool {
	var count int8
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if w.IsAlive(x+j, y+i) {
				count++
			}
		}
	}
	if (w.board[y][x] && count >= 2 && count <= 3) || (w.board[y][x] == false && count == 3) {
		return true
	}
	return false
}

// Next generate next configuration
func (w World) Next() World {
	newWorld := DuplicateWorld(w)
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			newWorld.board[i][j] = w.IsAliveNext(j, i)
		}
	}
	return newWorld
}
