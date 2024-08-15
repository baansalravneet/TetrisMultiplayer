package pieces

import "tetris/component"

type Bar struct {
	X           int
	Y           int
	Orientation int
}

func NewBar() *Bar {
	return &Bar{1, 5, 0}
}

func (c *Bar) Position() (int, int) {
	return c.X, c.Y
}

func (c *Bar) Pixels() []component.Pixel {
	if c.Orientation == 0 {
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
	c.X = x
	c.Y = y
}

func (c *Bar) Rotate() {
	c.Orientation = 1 - c.Orientation
}

func (c *Bar) RotateBack() {
	c.Orientation = 1 - c.Orientation
}
