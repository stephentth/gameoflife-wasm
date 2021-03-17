package algorithm

import (
	"log"
	"testing"
)

func compareWorld(a, b World) bool {
	if a.width != b.width && a.height != b.height {
		return false
	}
	for i := 0; i < a.height; i++ {
		for j := 0; j < a.width; j++ {
			if a.board[i][j] != b.board[i][j] {
				return false
			}
		}
	}
	return true
}

var testCasesNextGen = []struct {
	have World
	want World
}{
	{
		World{[][]bool{
			{true, true, false},
			{false, false, false},
			{false, false, false}}, 3, 3},
		World{[][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false}}, 3, 3},
	},
	{ // Block
		World{[][]bool{
			{true, true, false, false},
			{true, false, false, false},
			{false, false, false, false},
			{false, false, false, false}}, 4, 4},
		World{[][]bool{
			{true, true, false, false},
			{true, true, false, false},
			{false, false, false, false},
			{false, false, false, false}}, 4, 4},
	},
	{ // Blinker
		World{[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, true, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false}}, 5, 5},
		World{[][]bool{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false}}, 5, 5},
	},
	{ // Glider
		World{[][]bool{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, false, true, false},
			{false, true, true, true, false},
			{false, false, false, false, false}}, 5, 5},
		World{[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, true, false, true, false},
			{false, false, true, true, false},
			{false, false, true, false, false}}, 5, 5},
	},
	{ // Glider stg 2
		World{[][]bool{
			{false, true, false, true, false},
			{false, false, true, true, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false}}, 5, 5},
		World{[][]bool{
			{false, false, false, true, false},
			{false, true, false, true, false},
			{false, false, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false}}, 5, 5},
	},
	{ // Tub
		World{[][]bool{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, true, false, true, false},
			{false, false, true, false, false},
			{false, false, false, false, false}}, 5, 5},
		World{[][]bool{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, true, false, true, false},
			{false, false, true, false, false},
			{false, false, false, false, false}}, 5, 5},
	},
}

func TestNextGen(t *testing.T) {
	for i, tc := range testCasesNextGen {
		log.Printf("Testcase #%v\n", i)

		tc.have.Next()
		if !compareWorld(tc.have, tc.want) {
			t.Fatalf("testcase #%v: got <%v> but want <%v>", i, tc.have.board, tc.want.board)
		}
	}
}

var testCasesWillAliveNextGen = []struct {
	have [][]bool
	want bool
}{
	{
		// Living cell have 3 neighbor -> Keep alive
		[][]bool{
			{true, true, true},
			{false, true, false},
			{false, false, false}},
		true,
	},
	{
		// Living cell have 4 neighbor -> Overpopulation
		[][]bool{
			{true, true, true},
			{false, true, false},
			{false, false, true}},
		false,
	},
	{
		// Dead cell have 3 neighbor -> Reproduce
		[][]bool{
			{true, false, false},
			{false, false, true},
			{true, false, false}},
		true,
	},
	{
		// Dead cell have 4 neighbor -> Keep it dead
		[][]bool{
			{true, false, false},
			{true, false, true},
			{true, false, false}},
		false,
	},
}

func TestWillAliveNextGen(t *testing.T) {
	for i, testCase := range testCasesWillAliveNextGen {
		world := World{testCase.have, 3, 3}
		state := world.willAliveNextGen(1, 1)
		if state != testCase.want {
			t.Fatalf("case no. %v: configuration <%v> need <%v> but have <%v>\n", i, world.board, testCase.want, state)
		}
	}
}
