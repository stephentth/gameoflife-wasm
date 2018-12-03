package algorithm

import (
	"log"
	"testing"
)

var testCases = []struct {
	have World
	want World
}{
	{
		World{[][]bool{
			[]bool{true, true, false},
			[]bool{false, false, false},
			[]bool{false, false, false}}, 3, 3},
		World{[][]bool{
			[]bool{false, false, false},
			[]bool{false, false, false},
			[]bool{false, false, false}}, 3, 3},
	},
	{ // Block
		World{[][]bool{
			[]bool{true, true, false, false},
			[]bool{true, false, false, false},
			[]bool{false, false, false, false},
			[]bool{false, false, false, false}}, 4, 4},
		World{[][]bool{
			[]bool{true, true, false, false},
			[]bool{true, true, false, false},
			[]bool{false, false, false, false},
			[]bool{false, false, false, false}}, 4, 4},
	},
	{ // Blinker
		World{[][]bool{
			[]bool{false, false, false, false, false},
			[]bool{false, false, false, false, false},
			[]bool{false, true, true, true, false},
			[]bool{false, false, false, false, false},
			[]bool{false, false, false, false, false}}, 5, 5},
		World{[][]bool{
			[]bool{false, false, false, false, false},
			[]bool{false, false, true, false, false},
			[]bool{false, false, true, false, false},
			[]bool{false, false, true, false, false},
			[]bool{false, false, false, false, false}}, 5, 5},
	},
	{ // Glider
		World{[][]bool{
			[]bool{false, false, false, false, false},
			[]bool{false, false, true, false, false},
			[]bool{false, false, false, true, false},
			[]bool{false, true, true, true, false},
			[]bool{false, false, false, false, false}}, 5, 5},
		World{[][]bool{
			[]bool{false, false, false, false, false},
			[]bool{false, false, false, false, false},
			[]bool{false, true, false, true, false},
			[]bool{false, false, true, true, false},
			[]bool{false, false, true, false, false}}, 5, 5},
	},
	{ // Glider stg 2
		World{[][]bool{
			[]bool{false, true, false, true, false},
			[]bool{false, false, true, true, false},
			[]bool{false, false, true, false, false},
			[]bool{false, false, false, false, false},
			[]bool{false, false, false, false, false}}, 5, 5},
		World{[][]bool{
			[]bool{false, false, false, true, false},
			[]bool{false, true, false, true, false},
			[]bool{false, false, true, true, false},
			[]bool{false, false, false, false, false},
			[]bool{false, false, false, false, false}}, 5, 5},
	},
	{ // Tub
		World{[][]bool{
			[]bool{false, false, false, false, false},
			[]bool{false, false, true, false, false},
			[]bool{false, true, false, true, false},
			[]bool{false, false, true, false, false},
			[]bool{false, false, false, false, false}}, 5, 5},
		World{[][]bool{
			[]bool{false, false, false, false, false},
			[]bool{false, false, true, false, false},
			[]bool{false, true, false, true, false},
			[]bool{false, false, true, false, false},
			[]bool{false, false, false, false, false}}, 5, 5},
	},
}

func TestGenerate(t *testing.T) {
	for i, testCase := range testCases {
		nextWorld := testCase.have.Next()

		log.Printf("Testcase #%v\n", i)
		testCase.have.Print()
		nextWorld.Print()
		if !CompareWorld(nextWorld, testCase.want) {
			t.Fatalf("testcase #%v: configuration <%v> want <%v> but have <%v>", i, testCase.have.board, testCase.want.board, nextWorld.board)
		}
	}
}
