package algorithm

// DuplicateWorld ..
func DuplicateWorld(world World) World {
	newBoard := make([][]bool, world.height)

	for i, row := range world.board {
		newBoard[i] = make([]bool, len(row))
		copy(newBoard[i], row)
	}
	return World{newBoard, world.height, world.width}
}

// CompareWorld ..
func CompareWorld(world1, world2 World) bool {
	if world1.width != world2.width && world1.height != world2.height {
		return false
	}
	for i := 0; i < world1.height; i++ {
		for j := 0; j < world1.width; j++ {
			if world1.board[i][j] != world2.board[i][j] {
				return false
			}
		}
	}
	return true
}
