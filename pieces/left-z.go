package pieces

import "tetris/component"

type LeftZ struct {
	Type        int
	X           int
	Y           int
	Orientation int
}

func NewLeftZ() *LeftZ {
	return &LeftZ{LEFT_Z, 1, 5, 0}
}

func (c *LeftZ) Position() (int, int) {
	return c.X, c.Y
}

func (c *LeftZ) Pixels() []component.Pixel {
	if c.Orientation == 0 {
		return []component.Pixel{
			{X: 0, Y: 1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
			{X: 1, Y: -1, C: 'O'},
		}
	}
	return []component.Pixel{
		{X: -1, Y: 0, C: 'O'},
		{X: 0, Y: 1, C: 'O'},
		{X: 0, Y: 0, C: 'O'},
		{X: 1, Y: 1, C: 'O'},
	}
}

func (c *LeftZ) NewPosition(x, y int) {
	c.X = x
	c.Y = y
}

func (c *LeftZ) Rotate() {
	c.Orientation = 1 - c.Orientation
}

func (c *LeftZ) RotateBack() {
	c.Orientation = 1 - c.Orientation
}
