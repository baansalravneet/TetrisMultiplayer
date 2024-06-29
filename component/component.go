package component

type Coords struct {
	X int
	Y int
}

type Component interface {
	Position() Coords
	Pixels() []Coords
	Char() rune
	NewPosition(int, int)
}

type Dot struct {
	x    int
	y    int
	char rune
}

func NewDot() *Dot {
	return &Dot{2, 2, 'o'}
}

func (c *Dot) Position() Coords {
	return Coords{c.x, c.y}
}

func (c *Dot) Char() rune {
	return c.char
}

func (c *Dot) Pixels() []Coords {
	return []Coords{{0, 0}}
}

func (c *Dot) NewPosition(x, y int) {
	c.x = x
	c.y = y
}
