package main

import (
	"fmt"
	"io/fs"
	"syscall/js"
	"time"

	root "github.com/stephentt-me/gameoflife-wasm"
	"github.com/stephentt-me/gameoflife-wasm/algorithm"
)

var canvasHeight = 500
var canvasWidth = 1000
var cellSize = 10
var backgroundColor = "rgb(245, 246, 247)"
var cellColor = "rgb(150,212,168)"

func renderCanvas(world *algorithm.World, canvas *js.Value, canvasHeight, canvasWidth int) {
	board, height, width := world.Get()
	for i := 0; i < height; i++ { // y
		for j := 0; j < width; j++ { // x
			if board[i][j] {
				canvas.Set("fillStyle", cellColor)
			} else {
				canvas.Set("fillStyle", backgroundColor)
			}
			canvas.Call("fillRect", cellSize*j, cellSize*i, cellSize*(j+1), cellSize*(i+1))
		}
	}
}

func main() {
	doc := js.Global().Get("document")
	generationNumber := doc.Call("getElementById", "generationNumber")
	canvas := doc.Call("getElementById", "gof")
	canvasCtx := canvas.Call("getContext", "2d")
	generationNumber.Set("innerHTML", 0)
	dropdownNode := doc.Call("getElementById", "presetInput")

	dirs, err := fs.ReadDir(root.Patterns, "patterns")
	if err != nil {
		panic(err)
	}

	patternNames := make([]string, 0)
	for _, f := range dirs {
		fmt.Println("loaded pattern", f.Name())
		patternNames = append(patternNames, f.Name())
		optNode := doc.Call("createElement", "option")
		optNode.Set("innerHTML", f.Name())
		dropdownNode.Call("appendChild", optNode)
	}

	world := algorithm.NewWorld(canvasHeight/cellSize, canvasWidth/cellSize)

	canvas.Call("setAttribute", "width", canvasWidth)
	canvas.Call("setAttribute", "height", canvasHeight)
	canvasCtx.Set("fillStyle", backgroundColor)
	canvasCtx.Call("fillRect", 0, 0, canvasWidth, canvasHeight)

	renderCanvas(&world, &canvasCtx, canvasHeight, canvasWidth)
	time.Sleep(1 * time.Second)

	for i := 1; ; i++ {
		renderCanvas(&world, &canvasCtx, canvasHeight, canvasWidth)
		generationNumber.Set("innerHTML", i)
		time.Sleep(80 * time.Millisecond)

		world.Next()
	}
}
