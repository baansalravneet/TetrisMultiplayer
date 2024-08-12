package component

type Border struct {
	topLeft     rune
	bottomLeft  rune
	topRight    rune
	bottomRight rune
	horizontal  rune
	vertical    rune
	Top         int
	Left        int
	Bottom      int
	Right       int
}

func NewBorder(top, left, bottom, right int) *Border { // these are absolute positions
	return &Border{'╔', '╚', '╗', '╝', '═', '║', top, left, bottom, right}
}

func (c *Border) Position() (int, int) {
	return c.Top, c.Left
}

func (c *Border) Pixels() []Pixel {
	pixels := []Pixel{}
	pixels = append(pixels, Pixel{0, 0, c.topLeft})
	pixels = append(pixels, Pixel{0, c.Right - c.Left, c.topRight})
	pixels = append(pixels, Pixel{c.Bottom - c.Top, 0, c.bottomLeft})
	pixels = append(pixels, Pixel{c.Bottom - c.Top, c.Right - c.Left, c.bottomRight})
	for i := 1; i < c.Bottom-c.Top; i++ {
		pixels = append(pixels, Pixel{i, 0, c.vertical})
		pixels = append(pixels, Pixel{i, c.Right - c.Left, c.vertical})
	}
	for j := 1; j < c.Right-c.Left; j++ {
		pixels = append(pixels, Pixel{0, j, c.horizontal})
		pixels = append(pixels, Pixel{c.Bottom - c.Top, j, c.horizontal})
	}
	return pixels
}
