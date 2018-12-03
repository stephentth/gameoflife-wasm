package main

import (
	"syscall/js"
	"time"

	gof "gitlab.com/trunghieu.truong.bit/gameoflife/algorithm"
)

func main() {
	doc := js.Global().Get("document")
	board := doc.Call("getElementById", "gof")
	generationNumber := doc.Call("getElementById", "generationNumber")

	world := gof.NewWorld(50, 30)
	worldHTML := gof.RenderWorld(world)
	board.Set("innerHTML", worldHTML)
	generationNumber.Set("innerHTML", 0)
	time.Sleep(2 * time.Second)

	for i := 1; i < 1000; i++ {
		world = gof.NextGeneration(world)
		worldHTML = gof.RenderWorld(world)
		board.Set("innerHTML", worldHTML)
		generationNumber.Set("innerHTML", i)
		time.Sleep(200 * time.Microsecond)
	}
}
