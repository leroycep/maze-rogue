package generate

import "math/rand"

func MakeMazes(bakedRooms []int, width, height int) []int {
	maze := make([]int, width*height)
	for idx, value := range bakedRooms {
		maze[idx] = value
	}
	for i := 0; i < width; i += 2 {
	YLoop:
		for j := 0; j < height; j += 2 {
			// Check for a place to start the maze
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if isOccupied(i+di, j+dj, width, height, maze) {
						// Oh, not a valid place to start a maze
						continue YLoop
					}
				}
			}
			// There is a place to start the maze!
			maze = digMaze(i, j, maze, width, height)
		}
	}
	return maze
}

func isOccupied(x, y, gridWidth, gridHeight int, grid []int) bool {
	if x < 0 || y < 0 || x >= gridWidth || y >= gridHeight {
		return false
	}
	return grid[(y*gridWidth)+x] != 0
}

type vector struct {
	x, y int
}

func digMaze(x, y int, maze []int, width, height int) []int {
	mymaze := make([]int, len(maze))
	for idx, t := range maze {
		mymaze[idx] = t
	}
	directions := []vector{vector{0, -2}, vector{0, 2}, vector{-2, 0}, vector{2, 0}}
	for i := 0; i < len(directions); i++ {
		n := rand.Intn(len(directions))
		directions[i], directions[n] = directions[n], directions[i]
	}
DirLoop:
	for _, dir := range directions {
		if x+dir.x < 0 || x+dir.x >= width || y+dir.y < 0 || y+dir.y >= height || isOccupied(x+dir.x, y+dir.y, width, height, mymaze) {
			// Oh, not a valid place to start a maze
			continue DirLoop
		}
		mymaze[((y+dir.y)*width)+x+dir.x] = 2
		mymaze[((y+dir.y/2)*width)+x+dir.x/2] = 2
		mymaze = digMaze(x+dir.x, y+dir.y, mymaze, width, height)
	}
	return mymaze
}
