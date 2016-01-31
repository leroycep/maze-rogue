package generate

func ConnectRooms(maze []int, width, height, regionid int) ([]int, int) {
	connected := make([]int, len(maze))
	for idx, value := range maze {
		connected[idx] = value
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if isEdge(i, j, maze, width, height) {
				connected[(j*width)+i] = -1
			}
		}
	}
	return connected, regionid
}

func isEdge(x, y int, maze []int, width, height int) bool {
	if t1, t2 := getTile(x-1, y, maze, width, height), getTile(x+1, y, maze, width, height); t1 != t2 &&
		t1 != 0 && t2 != 0 && getTile(x, y, maze, width, height) == 0 {
		return true
	} else if t1, t2 := getTile(x, y-1, maze, width, height), getTile(x, y+1, maze, width, height); t1 != t2 &&
		t1 != 0 && t2 != 0 && getTile(x, y, maze, width, height) == 0 {
		return true
	} else {
		return false
	}
}

func getTile(x, y int, maze []int, width, height int) int {
	if x < 0 || y < 0 || x >= width || y >= height {
		return 0
	}
	return maze[(y*width)+x]
}
