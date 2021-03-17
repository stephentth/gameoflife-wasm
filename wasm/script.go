package main

import (
	"syscall/js"
	"time"

	"github.com/stephentt-me/gameoflife-wasm/algorithm"
)

func renderCanvas(world *algorithm.World, canvas *js.Value, canvasHeight, canvasWidth int) {
	canvas.Set("fillStyle", "rgb(255, 0, 255)")
	canvas.Call("fillRect", 0, 0, canvasHeight, canvasWidth)

	board, height, width := world.Get()
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] {
				canvas.Call("fillReact")
			}
		}
	}
}

func main() {
	doc := js.Global().Get("document")
	generationNumber := doc.Call("getElementById", "generationNumber")
	canvas := doc.Call("getElementById", "gof")
	canvasCtx := canvas.Call("getContext", "2d")
	generationNumber.Set("innerHTML", 0)

	canvasHeight := 250
	canvasWidth := 750

	world := algorithm.NewWorld(30, 80)
	// init the canvas
	canvas.Call("setAttribute", "width", canvasWidth)
	canvas.Call("setAttribute", "height", canvasHeight)
	canvas.Set("fillStyle", "rgb(255, 0, 255)")
	canvas.Call("fillRect", 0, 0, canvasWidth, canvasHeight)

	// renderCanvas(&world, &canvasCtx, canvasHeight, canvasWidth)
	// time.Sleep(200 * time.Microsecond)

	for i := 1; i < 9999; i++ {
		renderCanvas(&world, &canvasCtx, canvasHeight, canvasWidth)
		generationNumber.Set("innerHTML", i)
		time.Sleep(200 * time.Microsecond)

		world.Next()
	}
}
