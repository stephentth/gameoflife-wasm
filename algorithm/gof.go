package algorithm

// CheckNeighbor ..
func CheckNeighbor(neighbor [3][3]bool) bool {
	var count int8
	me := neighbor[1][1]
	for _, row := range neighbor {
		for _, value := range row {
			if value {
				count++
			}
		}
	}
	if me {
		count--
	}
	// TODO: LOL report golang's bug
	if (me) && (count >= 2) && (count <= 3) {
		return true
	}
	if !me && (count == 3) {
		return true
	}
	return false
}

// NextGeneration ..
func NextGeneration(world [][]bool) [][]bool {
	var c, tl, t, tr, r, br, b, bl, l bool
	var neighbor [3][3]bool

	height := len(world)
	width := len(world[0]) // TODO: failsafe
	newWorld := DuplicateWorld(world)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c = world[i][j]
			if i-1 < 0 || j-1 < 0 {
				tl = false
			} else {
				tl = world[i-1][j-1]
			}
			if i-1 < 0 {
				t = false
			} else {
				t = world[i-1][j]
			}
			if i-1 < 0 || j+1 >= width {
				tr = false
			} else {
				tr = world[i-1][j+1]
			}
			if j+1 >= width {
				r = false
			} else {
				r = world[i][j+1]
			}
			if i+1 >= height || j+1 >= width {
				br = false
			} else {
				br = world[i+1][j+1]
			}
			if i+1 >= height {
				b = false
			} else {
				b = world[i+1][j]
			}
			if i+1 >= height || j-1 < 0 {
				bl = false
			} else {
				bl = world[i+1][j-1]
			}
			if j-1 < 0 {
				l = false
			} else {
				l = world[i][j-1]
			}
			neighbor = [3][3]bool{
				[3]bool{tl, t, tr},
				[3]bool{l, c, r},
				[3]bool{bl, b, br},
			}
			newWorld[i][j] = CheckNeighbor(neighbor)
		}
	}
	return newWorld
}
