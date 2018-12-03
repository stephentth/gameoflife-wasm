package algorithm

import (
	"fmt"
	"math/rand"
)

const LIVECELL = "<td class='live cell'></td>"
const DEADCELL = "<td class='dead cell'></td>"

// NewWorld make 2D board pre-fill with cell
func NewWorld(width, height int) [][]bool {
	world := make([][]bool, height)
	for i := range world {
		world[i] = make([]bool, width)
		// populate random cell
		for j := range world[i] {
			// TODO: dehardcode random function
			if rand.Int()%3 == 0 {
				world[i][j] = true
			} else {
				world[i][j] = false
			}
		}
	}
	return world
}

// PrintWorld show the 2D array of world to std.out
func PrintWorld(world [][]bool) {
	for _, row := range world {
		for _, value := range row {
			if value {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// RenderWorld convert game'world to HTML Table string
func RenderWorld(world [][]bool) string {
	var table string
	table += "<table>\n"
	for _, row := range world {
		table += "<tr>"
		for _, value := range row {
			if value {
				table += LIVECELL
			} else {
				table += DEADCELL
			}
		}
		table += "</tr>\n"
	}
	table += "</table>\n"
	return table
}

// DuplicateWorld ..
func DuplicateWorld(world [][]bool) [][]bool {
	height := len(world)
	newWorld := make([][]bool, height)

	for i, row := range world {
		newWorld[i] = make([]bool, len(row))
		copy(newWorld[i], row)
	}
	return newWorld
}
