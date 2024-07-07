package component

type Dot struct {
	x int
	y int
}

func NewDot() *Dot {
	return &Dot{5, 5}
}

func (c *Dot) Position() (int, int) {
	return c.x, c.y
}

func (c *Dot) Pixels() []Pixel {
	return []Pixel{{0, 0, 'O'}}
}

func (c *Dot) NewPosition(x, y int) {
	c.x = x
	c.y = y
}

func (c *Dot) Rotate() {
	// redundant
}

func (c *Dot) RotateBack() {
	// redundant
}
