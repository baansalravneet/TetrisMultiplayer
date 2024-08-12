package pieces

import "tetris/component"

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

func (c *LeftZ) Pixels() []component.Pixel {
	if c.orientation == 0 {
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
	c.x = x
	c.y = y
}

func (c *LeftZ) Rotate() {
	c.orientation = 1 - c.orientation
}

func (c *LeftZ) RotateBack() {
	c.orientation = 1 - c.orientation
}
