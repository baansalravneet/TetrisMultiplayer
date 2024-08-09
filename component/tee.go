package component

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

func (c *Tee) Pixels() []Pixel {
	if c.orientation == 0 {
		return []Pixel{
			{0, 0, 'O'},
			{0, -1, 'O'},
			{0, 1, 'O'},
			{1, 0, 'O'},
		}
	} else if c.orientation == 1 {
		return []Pixel{
			{0, -1, 'O'},
			{0, 0, 'O'},
			{1, 0, 'O'},
			{-1, 0, 'O'},
		}
	} else if c.orientation == 2 {
		return []Pixel{
			{0, 0, 'O'},
			{-1, 0, 'O'},
			{0, 1, 'O'},
			{0, -1, 'O'},
		}
	}
	return []Pixel{
		{0, 0, 'O'},
		{0, 1, 'O'},
		{-1, 0, 'O'},
		{1, 0, 'O'},
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
