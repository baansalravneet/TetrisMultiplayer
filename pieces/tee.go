package pieces

import "tetris/component"

type Tee struct {
	x           int
	y           int
	orientation int
}

func NewTee() *Tee {
	return &Tee{1, 5, 0}
}

func (c *Tee) Position() (int, int) {
	return c.x, c.y
}

func (c *Tee) Pixels() []component.Pixel {
	if c.orientation == 0 {
		return []component.Pixel{
			{X: 0, Y: 0, C: 'O'},
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 1, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
		}
	} else if c.orientation == 1 {
		return []component.Pixel{
			{X: 0, Y: -1, C: 'O'},
			{X: 0, Y: 0, C: 'O'},
			{X: 1, Y: 0, C: 'O'},
			{X: -1, Y: 0, C: 'O'},
		}
	} else if c.orientation == 2 {
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
	c.x = x
	c.y = y
}

func (c *Tee) Rotate() {
	c.orientation = (c.orientation + 1) % 4
}

func (c *Tee) RotateBack() {
	c.orientation = (c.orientation - 1 + 4) % 4
}
