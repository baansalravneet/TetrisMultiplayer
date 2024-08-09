package component

type Text struct {
	s string
	x int
	y int
}

func NewText(s string, x, y int) *Text {
	return &Text{s, x, y}
}

func (c *Text) Position() (int, int) {
	return c.x, c.y
}

func (c *Text) Pixels() []Pixel {
	pixels := []Pixel{}
	for j, c := range c.s {
		pixels = append(pixels, Pixel{0, j, c})
	}
	return pixels
}

func (c *Text) NewPosition(x, y int) {
	// you cannot change the position of a border
}

func (c *Text) Rotate() {
	// you cannot rotate the border
}

func (c *Text) RotateBack() {
	// you cannot rotate the border
}
