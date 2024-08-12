package pieces

import "tetris/component"

type Box struct {
	Type        int
	X int
	Y int
}

func NewBox() *Box {
	return &Box{BOX, 1, 5}
}

func (c *Box) Position() (int, int) {
	return c.X, c.Y
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
	c.X = x
	c.Y = y
}

func (c *Box) Rotate() {
	// redundant
}

func (c *Box) RotateBack() {
	// redundant
}
