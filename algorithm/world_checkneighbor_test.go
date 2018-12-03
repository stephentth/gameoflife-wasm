package algorithm

import (
	"testing"
)

var testCasesCalculateNeighbor = []struct {
	have [][]bool
	want bool
}{
	{
		// Living cell have 3 neighbor -> Keep alive
		[][]bool{
			[]bool{true, true, true},
			[]bool{false, true, false},
			[]bool{false, false, false}},
		true,
	},
	{
		// Living cell have 4 neighbor -> Overpopulation
		[][]bool{
			[]bool{true, true, true},
			[]bool{false, true, false},
			[]bool{false, false, true}},
		false,
	},
	{
		// Dead cell have 3 neighbor -> Reproduce
		[][]bool{
			[]bool{true, false, false},
			[]bool{false, false, true},
			[]bool{true, false, false}},
		true,
	},
	{
		// Dead cell have 4 neighbor -> Keep it dead
		[][]bool{
			[]bool{true, false, false},
			[]bool{true, false, true},
			[]bool{true, false, false}},
		false,
	},
}

func TestCheckNeighbor(t *testing.T) {
	for i, testCase := range testCasesCalculateNeighbor {
		world := World{testCase.have, 3, 3}
		state := world.IsAliveNext(1, 1)
		if state != testCase.want {
			t.Fatalf("case no. %v: configuration <%v> need <%v> but have <%v>\n", i, world.board, testCase.want, state)
		}
	}
}
