package component

type RightZ struct {
	x           int
	y           int
	orientation int
}

func NewRightZ() *RightZ {
	return &RightZ{1, 5, 0}
}

func (c *RightZ) Position() (int, int) {
	return c.x, c.y
}

func (c *RightZ) Pixels() []Pixel {
	if c.orientation == 0 {
		return []Pixel{
			{0, -1, 'O'},
			{0, 0, 'O'},
			{1, 0, 'O'},
			{1, 1, 'O'},
		}
	}
	return []Pixel{
		{-1, 1, 'O'},
		{0, 1, 'O'},
		{0, 0, 'O'},
		{1, 0, 'O'},
	}
}

func (c *RightZ) NewPosition(x, y int) {
	c.x = x
	c.y = y
}

func (c *RightZ) Rotate() {
	c.orientation = 1 - c.orientation
}

func (c *RightZ) RotateBack() {
	c.orientation = 1 - c.orientation
}
