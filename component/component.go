package component

type Pixel struct {
	X int
	Y int
	C rune
}

type Component interface {
	Position() (int, int)
	Pixels() []Pixel
	NewPosition(int, int)
}

type Dot struct {
	x    int
	y    int
	char rune
}

func NewDot() *Dot {
	return &Dot{5, 5, 'o'}
}

func (c *Dot) Position() (int, int) {
	return c.x, c.y
}

func (c *Dot) Pixels() []Pixel {
	return []Pixel{{0, 0, 'O'}}
}

func (c *Dot) NewPosition(x, y int) {
	c.x = x
	c.y = y
}

type Border struct {
	topLeft     rune
	bottomLeft  rune
	topRight    rune
	bottomRight rune
	horizontal  rune
	vertical    rune
	top         int
	left        int
	bottom      int
	right       int
}

func NewBorder(top, left, bottom, right int) *Border {
	return &Border{'╔', '╚', '╗', '╝', '═', '║', top, left, bottom, right}
}

func (c *Border) Position() (int, int) {
	return 0, 0
}

func (c *Border) Pixels() []Pixel {
	pixels := []Pixel{}
	pixels = append(pixels, Pixel{c.top, c.left, c.topLeft})
	pixels = append(pixels, Pixel{c.top, c.right, c.topRight})
	pixels = append(pixels, Pixel{c.bottom, c.left, c.bottomLeft})
	pixels = append(pixels, Pixel{c.bottom, c.right, c.bottomRight})
	for i := c.top + 1; i < c.bottom; i++ {
		pixels = append(pixels, Pixel{i, c.left, c.vertical})
		pixels = append(pixels, Pixel{i, c.right, c.vertical})
	}
	for j := c.left + 1; j < c.right; j++ {
		pixels = append(pixels, Pixel{c.top, j, c.horizontal})
		pixels = append(pixels, Pixel{c.bottom, j, c.horizontal})
	}
	return pixels
}

func (c *Border) NewPosition(x, y int) {
	// you cannot change the position of a border
}
