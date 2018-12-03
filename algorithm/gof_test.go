package algorithm

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	have [3][3]bool
	want bool
}{
	{
		// Living cell have 3 neighbor -> Keep alive
		[3][3]bool{
			[3]bool{true, true, true},
			[3]bool{false, true, false},
			[3]bool{false, false, false}},
		true,
	},
	{
		// Living cell have 4 neighbor -> Overpopulation
		[3][3]bool{
			[3]bool{true, true, true},
			[3]bool{false, true, false},
			[3]bool{false, false, true}},
		false,
	},
	{
		// Dead cell have 3 neighbor -> Reproduce
		[3][3]bool{
			[3]bool{true, false, false},
			[3]bool{false, false, true},
			[3]bool{true, false, false}},
		true,
	},
	{
		// Dead cell have 4 neighbor -> Keep it dead
		[3][3]bool{
			[3]bool{true, false, false},
			[3]bool{true, false, true},
			[3]bool{true, false, false}},
		false,
	},
}

func TestCheckNeighbor(t *testing.T) {
	for _, testCase := range testCases {
		if result := CheckNeighbor(testCase.have); result != testCase.want {
			t.Fatalf("FAILED: Configuation %v want <<%v>> but gave <<%v>>", testCase.have, testCase.want, result)
		}
	}
}

func TestNextGeneration(t *testing.T) {
	world := [][]bool{
		[]bool{false, false, false, false, false},
		[]bool{false, false, false, false, false},
		[]bool{false, true, true, true, false},
		[]bool{false, false, false, false, false},
		[]bool{false, false, false, false, false},
	}
	PrintWorld(world)
	world = NextGeneration(world)
	PrintWorld(world)
	world = NextGeneration(world)
	PrintWorld(world)
	world = NextGeneration(world)
	fmt.Printf("%v", RenderWorld(world))
}
