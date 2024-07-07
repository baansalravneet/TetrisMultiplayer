package component

type LeftZ struct {
	x           int
	y           int
	orientation int
}

func NewLeftZ() *LeftZ {
	return &LeftZ{1, 5, 0}
}

func (c *LeftZ) Position() (int, int) {
	return c.x, c.y
}

func (c *LeftZ) Pixels() []Pixel {
	if c.orientation == 0 {
		return []Pixel{
			{0, 1, 'O'},
			{0, 0, 'O'},
			{1, 0, 'O'},
			{1, -1, 'O'},
		}
	}
	return []Pixel{
		{-1, 0, 'O'},
		{0, 1, 'O'},
		{0, 0, 'O'},
		{1, 1, 'O'},
	}
}

func (c *LeftZ) NewPosition(x, y int) {
	c.x = x
	c.y = y
}

func (c *LeftZ) Rotate() {
	c.orientation = 1 - c.orientation
}

func (c *LeftZ) RotateBack() {
	c.orientation = 1 - c.orientation
}
