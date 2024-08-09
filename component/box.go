package component

type Box struct {
	x int
	y int
}

func NewBox() *Box {
	return &Box{1, 5}
}

func (c *Box) Position() (int, int) {
	return c.x, c.y
}

func (c *Box) Pixels() []Pixel {
	return []Pixel{
		{0, 0, 'O'},
		{1, 0, 'O'},
		{0, 1, 'O'},
		{1, 1, 'O'},
	}
}

func (c *Box) NewPosition(x, y int) {
	c.x = x
	c.y = y
}

func (c *Box) Rotate() {
	// redundant
}

func (c *Box) RotateBack() {
	// redundant
}
