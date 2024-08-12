package pieces

import "tetris/component"

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

func (c *RightZ) Pixels() []component.Pixel {
	if c.orientation == 0 {
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
	c.x = x
	c.y = y
}

func (c *RightZ) Rotate() {
	c.orientation = 1 - c.orientation
}

func (c *RightZ) RotateBack() {
	c.orientation = 1 - c.orientation
}
