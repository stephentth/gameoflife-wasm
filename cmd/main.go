package main

import "github.com/stephentt-me/gameoflife-wasm/algorithm"

func main() {
	world := algorithm.NewWorld(20, 20)
	world.Print()
}
