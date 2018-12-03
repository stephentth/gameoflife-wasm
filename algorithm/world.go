package algorithm

import (
	"fmt"
	"math/rand"
)

const LIVECELL = "<td class='cell live'></td>"
const DEADCELL = "<td class='cell dead'></td>"

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
	// fmt.Printf("x=%v y=%v result=%v\n", x, y, w.board[y][x])
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
		// fmt.Printf("this cell (%v, %v) is alive\n", x, y)
		return true
	}
	// fmt.Printf("this cell (%v, %v) is dead\n", x, y)
	return false
}

// Next generate next configuration
func (w World) Next() World {
	newWorld := DuplicateWorld(w)
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			// fmt.Printf("Scan at col=%v line=%v\n", j, i)
			newWorld.board[i][j] = w.IsAliveNext(j, i)
		}
	}
	return newWorld
}

// Print World into std.out
func (w World) Print() {
	for i := 0; i < w.height; i++ {
		for j := 0; j < w.width; j++ {
			if w.board[i][j] {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// RenderHTML ..
func (w World) RenderHTML() string {
	html := ""
	html += "<table>"
	for i := 0; i < w.height; i++ {
		html += "<tr>"
		for j := 0; j < w.width; j++ {
			if w.board[i][j] {
				html += LIVECELL
			} else {
				html += DEADCELL
			}
		}
		html += "</tr>\n"
	}
	html += "</table>\n"
	return html
}
