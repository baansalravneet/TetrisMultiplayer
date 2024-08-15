package pieces

import "tetris/component"

type LeftL struct {
	X           int
	Y           int
	Orientation int
}

func NewLeftL() *LeftL {
	return &LeftL{1, 5, 0}
}

func (c *LeftL) Position() (int, int) {
	return c.X, c.Y
}

func (c *LeftL) Pixels() []component.Pixel {
	if c.Orientation == 0 {
		return []component.Pixel{
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
			{X: 1, Y: 1, C: 'O'},
		}
	} else if c.Orientation == 1 {
		return []component.Pixel{
			{X: -1, Y: 0, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
			{X: 1, Y: -1, C: 'O'},
		}
	} else if c.Orientation == 2 {
		return []component.Pixel{
			{X: -1, Y: -1, C: 'O'},
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
		}
	}
	return []component.Pixel{
		{X: 1, Y: 0, C: 'O'},
		{X: 0, Y: 0, C: 'O'},
		{X: -1, Y: 0, C: 'O'},
		{X: -1, Y: 1, C: 'O'},
	}
}

func (c *LeftL) NewPosition(x, y int) {
	c.X = x
	c.Y = y
}

func (c *LeftL) Rotate() {
	c.Orientation = (c.Orientation + 1) % 4
}

func (c *LeftL) RotateBack() {
	c.Orientation = (c.Orientation - 1 + 4) % 4
}
