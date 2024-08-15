package pieces

import "tetris/component"

type Tee struct {
	X           int
	Y           int
	Orientation int
}

func NewTee() *Tee {
	return &Tee{1, 5, 0}
}

func (c *Tee) Position() (int, int) {
	return c.X, c.Y
}

func (c *Tee) Pixels() []component.Pixel {
	if c.Orientation == 0 {
		return []component.Pixel{
			{X: 0, Y: 0, C: 'O'},
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
		}
	} else if c.Orientation == 1 {
		return []component.Pixel{
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
			{X: -1, Y: 0, C: 'O'},
		}
	} else if c.Orientation == 2 {
		return []component.Pixel{
			{X: 0, Y: 0, C: 'O'},
			{X: -1, Y: 0, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
			{X: 0, Y: -1, C: 'O'},
		}
	}
	return []component.Pixel{
		{X: 0, Y: 0, C: 'O'},
		{X: 0, Y: 1, C: 'O'},
		{X: -1, Y: 0, C: 'O'},
		{X: 1, Y: 0, C: 'O'},
	}
}

func (c *Tee) NewPosition(x, y int) {
	c.X = x
	c.Y = y
}

func (c *Tee) Rotate() {
	c.Orientation = (c.Orientation + 1) % 4
}

func (c *Tee) RotateBack() {
	c.Orientation = (c.Orientation - 1 + 4) % 4
}
