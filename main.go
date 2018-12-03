package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	gof "gitlab.com/trunghieu.truong.bit/gameoflife/algorithm"
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
	fmt.Println(runtime.GOOS)
	world := gof.NewWorld(15, 15)
	gof.PrintWorld(world)
	for i := 0; i < 1000; i++ {
		time.Sleep(100 * time.Millisecond)
		clearScreen()
		world = gof.NextGeneration(world)
		gof.PrintWorld(world)
	}
}

func main() {
	run()
}
