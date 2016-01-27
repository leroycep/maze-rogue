package model

type Beast struct {
	Health int
	X, Y   int
}

func (p *Beast) GetHealth() int {
	return p.Health
}

func (p *Beast) SetHealth(value int) {
	p.Health = value
}

func (p *Beast) GetX() int {
	return p.X
}

func (p *Beast) SetX(value int) {
	p.X = value
}

func (p *Beast) GetY() int {
	return p.Y
}

func (p *Beast) SetY(value int) {
	p.Y = value
}
