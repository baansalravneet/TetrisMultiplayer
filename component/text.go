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
