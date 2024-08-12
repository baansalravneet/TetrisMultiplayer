package pieces

import "tetris/component"

type Bar struct {
	x           int
	y           int
	orientation int
}

func NewBar() *Bar {
	return &Bar{1, 5, 0}
}

func (c *Bar) Position() (int, int) {
	return c.x, c.y
}

func (c *Bar) Pixels() []component.Pixel {
	if c.orientation == 0 {
		return []component.Pixel{
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
			{X: 0, Y: 2, C: 'O'},
		}
	}
	return []component.Pixel{
		{X: -1, Y: 0, C: 'O'},
		{X: 0, Y: 0, C: 'O'},
		{X: 1, Y: 0, C: 'O'},
		{X: 2, Y: 0, C: 'O'},
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
