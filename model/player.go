package model

type Player struct {
	Health int
	X, Y   int
}

func (p *Player) GetHealth() int {
	return p.Health
}

func (p *Player) SetHealth(value int) {
	p.Health = value
}

func (p *Player) GetX() int {
	return p.X
}

func (p *Player) SetX(value int) {
	p.X = value
}

func (p *Player) GetY() int {
	return p.Y
}

func (p *Player) SetY(value int) {
	p.Y = value
}
