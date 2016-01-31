package generate

func TrimPaths(maze []int, width, height int) []int {
	trimmed := make([]int, len(maze))
	for idx, value := range maze {
		trimmed[idx] = value
	}
LOOP:
	for {
		found := false
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				if trimmed[(j*width)+i] != 0 {
					count := 0
					if getTile(i, j-1, trimmed, width, height) == 0 {
						count++
					}
					if getTile(i, j+1, trimmed, width, height) == 0 {
						count++
					}
					if getTile(i-1, j, trimmed, width, height) == 0 {
						count++
					}
					if getTile(i+1, j, trimmed, width, height) == 0 {
						count++
					}
					if count >= 3 {
						found = true
						trimmed[(j*width)+i] = 0
					}
				}
			}
		}
		if !found {
			break LOOP
		}
	}
	return trimmed
}
