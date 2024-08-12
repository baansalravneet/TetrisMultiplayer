package pieces

import "tetris/component"

type Box struct {
	x int
	y int
}

func NewBox() *Box {
	return &Box{1, 5}
}

func (c *Box) Position() (int, int) {
	return c.x, c.y
}

func (c *Box) Pixels() []component.Pixel {
	return []component.Pixel{
		{X: 0, Y: 0, C: 'O'},
		{X: 1, Y: 0, C: 'O'},
		{X: 0, Y: 1, C: 'O'},
		{X: 1, Y: 1, C: 'O'},
	}
}

func (c *Box) NewPosition(x, y int) {
	c.x = x
	c.y = y
}

func (c *Box) Rotate() {
	// redundant
}

func (c *Box) RotateBack() {
	// redundant
}
