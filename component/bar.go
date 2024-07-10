package component

type Bar struct {
	x           int
	y           int
	orientation int

	id int
}

func NewBar() *Bar {
	return &Bar{1, 5, 0, BAR_ID}
}

func (c *Bar) Position() (int, int) {
	return c.x, c.y
}

func (c *Bar) Pixels() []Pixel {
	if c.orientation == 0 {
		return []Pixel{
			{0, -1, 'O'},
			{0, 0, 'O'},
			{0, 1, 'O'},
			{0, 2, 'O'},
		}
	}
	return []Pixel{
		{-1, 0, 'O'},
		{0, 0, 'O'},
		{1, 0, 'O'},
		{2, 0, 'O'},
	}
}

func (c *Bar) NewPosition(x, y int) {
	c.x = x
	c.y = y
}

func (c *Bar) Rotate() {
	c.orientation = 1 - c.orientation
}

func (c *Bar) RotateBack() {
	c.orientation = 1 - c.orientation
}

func (c *Bar) Id() int {
	return c.id
}
