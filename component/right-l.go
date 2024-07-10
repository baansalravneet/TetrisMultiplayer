package component

type RightL struct {
	x           int
	y           int
	orientation int

	id int
}

func NewRightL() *RightL {
	return &RightL{1, 5, 0, RIGHT_L_ID}
}

func (c *RightL) Position() (int, int) {
	return c.x, c.y
}

func (c *RightL) Pixels() []Pixel {
	if c.orientation == 0 {
		return []Pixel{
			{0, -1, 'O'},
			{0, 0, 'O'},
			{0, 1, 'O'},
			{1, -1, 'O'},
		}
	} else if c.orientation == 1 {
		return []Pixel{
			{-1, 0, 'O'},
			{0, 0, 'O'},
			{1, 0, 'O'},
			{-1, -1, 'O'},
		}
	} else if c.orientation == 2 {
		return []Pixel{
			{-1, 1, 'O'},
			{0, -1, 'O'},
			{0, 0, 'O'},
			{0, 1, 'O'},
		}
	}
	return []Pixel{
		{1, 0, 'O'},
		{0, 0, 'O'},
		{-1, 0, 'O'},
		{1, 1, 'O'},
	}
}

func (c *RightL) NewPosition(x, y int) {
	c.x = x
	c.y = y
}

func (c *RightL) Rotate() {
	c.orientation = (c.orientation + 1) % 4
}

func (c *RightL) RotateBack() {
	c.orientation = (c.orientation - 1 + 4) % 4
}

func (c *RightL) Id() int {
	return c.id
}
