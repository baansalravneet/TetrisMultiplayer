package pieces

import "tetris/component"

type LeftL struct {
	x           int
	y           int
	orientation int
}

func NewLeftL() *LeftL {
	return &LeftL{1, 5, 0}
}

func (c *LeftL) Position() (int, int) {
	return c.x, c.y
}

func (c *LeftL) Pixels() []component.Pixel {
	if c.orientation == 0 {
		return []component.Pixel{
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
			{X: 1, Y: 1, C: 'O'},
		}
	} else if c.orientation == 1 {
		return []component.Pixel{
			{X: -1, Y: 0, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
			{X: 1, Y: -1, C: 'O'},
		}
	} else if c.orientation == 2 {
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
	c.x = x
	c.y = y
}

func (c *LeftL) Rotate() {
	c.orientation = (c.orientation + 1) % 4
}

func (c *LeftL) RotateBack() {
	c.orientation = (c.orientation - 1 + 4) % 4
}
