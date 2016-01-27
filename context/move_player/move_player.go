package move_player

type Player interface {
	GetX() int
	SetX(value int)
	GetY() int
	SetY(value int)
}

type Room interface {
	Contains(x, y int) bool
}

func MovePlayer(player Player, rooms []Room, x, y int) {
	if x == 0 && y == 0 {
		// Do nothing
		return
	}
	nx, ny := player.GetX()+x, player.GetY()+y
	isInRoom := false
	for _, room := range rooms {
		if room.Contains(nx, ny) {
			isInRoom = true
		}
	}
	if !isInRoom {
		// Can't move there
		return
	}
	player.SetX(nx)
	player.SetY(ny)
}
