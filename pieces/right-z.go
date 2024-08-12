package pieces

import "tetris/component"

type RightZ struct {
	Type        int
	X           int
	Y           int
	Orientation int
}

func NewRightZ() *RightZ {
	return &RightZ{RIGHT_Z, 1, 5, 0}
}

func (c *RightZ) Position() (int, int) {
	return c.X, c.Y
}

func (c *RightZ) Pixels() []component.Pixel {
	if c.Orientation == 0 {
		return []component.Pixel{
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
			{X: 1, Y: 1, C: 'O'},
		}
	}
	return []component.Pixel{
		{X: -1, Y: 1, C: 'O'},
		{X: 0, Y: 1, C: 'O'},
		{X: 0, Y: 0, C: 'O'},
		{X: 1, Y: 0, C: 'O'},
	}
}

func (c *RightZ) NewPosition(x, y int) {
	c.X = x
	c.Y = y
}

func (c *RightZ) Rotate() {
	c.Orientation = 1 - c.Orientation
}

func (c *RightZ) RotateBack() {
	c.Orientation = 1 - c.Orientation
}
