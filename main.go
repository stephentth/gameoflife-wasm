package main

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	gof "github.com/stephentt-me/gameoflife-wasm/algorithm"
)

func clearScreen() {
	// Clear screen only in Linux or MacOS environment
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func run() {
	world := gof.NewWorld(15, 15)
	world.Print()
	time.Sleep(2 * time.Second)
	clearScreen()

	for i := 0; i < 1000; i++ {
		world = world.Next()
		world.Print()
		time.Sleep(100 * time.Millisecond)
		clearScreen()
	}
}

func main() {
	run()
}
