package generate

import "math/rand"

type edge struct {
	x, y    int
	regions [2]int
}

func ConnectRooms(maze []int, width, height, regionid int) ([]int, int) {
	connected := make([]int, len(maze))
	for idx, value := range maze {
		connected[idx] = value
	}
	edges := []edge{}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if is, e := isEdge(i, j, maze, width, height); is {
				edges = append(edges, e)
			}
		}
	}
	unified := make([]bool, regionid)
	unified[1] = true
	for len(edges) > 0 {
		connectedEdges := []edge{}
		for _, edge := range edges {
			if unified[edge.regions[0]] || unified[edge.regions[1]] {
				connectedEdges = append(connectedEdges, edge)
			}
		}
		e := connectedEdges[rand.Intn(len(connectedEdges))]
		connected[(e.y*width)+e.x] = regionid
		unified[e.regions[0]] = true
		unified[e.regions[1]] = true
		temp := []edge{}
		for _, value := range edges {
			if !(unified[value.regions[0]] && unified[value.regions[1]]) {
				if rand.Intn(100) == 0 {
					connected[(value.y*width)+value.x] = regionid
				}
				temp = append(temp, value)
			}
		}
		edges = temp
	}
	return connected, regionid
}

func isEdge(x, y int, maze []int, width, height int) (bool, edge) {
	if t1, t2 := getTile(x-1, y, maze, width, height), getTile(x+1, y, maze, width, height); t1 != t2 &&
		t1 != 0 && t2 != 0 && getTile(x, y, maze, width, height) == 0 {
		return true, edge{x, y, [2]int{t1, t2}}
	} else if t1, t2 := getTile(x, y-1, maze, width, height), getTile(x, y+1, maze, width, height); t1 != t2 &&
		t1 != 0 && t2 != 0 && getTile(x, y, maze, width, height) == 0 {
		return true, edge{x, y, [2]int{t1, t2}}
	} else {
		return false, edge{-1, -1, [2]int{-1, -1}}
	}
}

func getTile(x, y int, maze []int, width, height int) int {
	if x < 0 || y < 0 || x >= width || y >= height {
		return 0
	}
	return maze[(y*width)+x]
}
