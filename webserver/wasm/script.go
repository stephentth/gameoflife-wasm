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

	world := gof.NewWorld(30, 80)
	worldHTML := world.RenderHTML()
	board.Set("innerHTML", worldHTML)
	generationNumber.Set("innerHTML", 0)
	time.Sleep(2 * time.Second)

	for i := 1; i < 1000; i++ {
		world = world.Next()
		worldHTML = world.RenderHTML()
		board.Set("innerHTML", worldHTML)
		generationNumber.Set("innerHTML", i)
		time.Sleep(200 * time.Microsecond)
	}
}
