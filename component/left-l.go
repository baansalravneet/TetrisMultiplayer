package component

type LeftL struct {
	x           int
	y           int
	orientation int

	id int
}

func NewLeftL() *LeftL {
	return &LeftL{1, 5, 0, LEFT_L_ID}
}

func (c *LeftL) Position() (int, int) {
	return c.x, c.y
}

func (c *LeftL) Pixels() []Pixel {
	if c.orientation == 0 {
		return []Pixel{
			{0, -1, 'O'},
			{0, 0, 'O'},
			{0, 1, 'O'},
			{1, 1, 'O'},
		}
	} else if c.orientation == 1 {
		return []Pixel{
			{-1, 0, 'O'},
			{0, 0, 'O'},
			{1, 0, 'O'},
			{1, -1, 'O'},
		}
	} else if c.orientation == 2 {
		return []Pixel{
			{-1, -1, 'O'},
			{0, -1, 'O'},
			{0, 0, 'O'},
			{0, 1, 'O'},
		}
	}
	return []Pixel{
		{1, 0, 'O'},
		{0, 0, 'O'},
		{-1, 0, 'O'},
		{-1, 1, 'O'},
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

func (c *LeftL) Id() int {
	return c.id
}
