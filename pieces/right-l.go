package pieces

import "tetris/component"

type RightL struct {
	Type        int
	X           int
	Y           int
	Orientation int
}

func NewRightL() *RightL {
	return &RightL{RIGHT_L, 1, 5, 0}
}

func (c *RightL) Position() (int, int) {
	return c.X, c.Y
}

func (c *RightL) Pixels() []component.Pixel {
	if c.Orientation == 0 {
		return []component.Pixel{
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
			{X: 1, Y: -1, C: 'O'},
		}
	} else if c.Orientation == 1 {
		return []component.Pixel{
			{X: -1, Y: 0, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
			{X: -1, Y: -1, C: 'O'},
		}
	} else if c.Orientation == 2 {
		return []component.Pixel{
			{X: -1, Y: 1, C: 'O'},
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
		}
	}
	return []component.Pixel{
		{X: 1, Y: 0, C: 'O'},
		{X: 0, Y: 0, C: 'O'},
		{X: -1, Y: 0, C: 'O'},
		{X: 1, Y: 1, C: 'O'},
	}
}

func (c *RightL) NewPosition(x, y int) {
	c.X = x
	c.Y = y
}

func (c *RightL) Rotate() {
	c.Orientation = (c.Orientation + 1) % 4
}

func (c *RightL) RotateBack() {
	c.Orientation = (c.Orientation - 1 + 4) % 4
}
