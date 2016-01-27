package model

type Area struct {
	X, Y, W, H int
}

func (area Area) Contains(x, y int) bool {
	return x >= area.X && x < area.X+area.W && y >= area.Y && y < area.Y+area.H
}
