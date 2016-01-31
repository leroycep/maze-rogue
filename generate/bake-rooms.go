package generate

func BakeRooms(rooms []Room, w, h, regionid int) ([]int, int) {
	grid := make([]int, w*h)
	for _, room := range rooms {
		for i := room.X; i < room.X+room.W; i++ {
			for j := room.Y; j < room.Y+room.H; j++ {
				setAt(i, j, regionid, grid, w)
			}
		}
		regionid++
	}
	return grid, regionid
}

func setAt(x, y, value int, grid []int, gridWidth int) {
	grid[(y*gridWidth)+x] = value
}
