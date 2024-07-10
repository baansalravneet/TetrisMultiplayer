package component

type Rubble struct {
	pixels [][]bool

	id int
}

func NewRubble(height, width int) *Rubble {
	pixels := make([][]bool, height)
	for i := range pixels {
		pixels[i] = make([]bool, width)
	}
	return &Rubble{pixels, RUBBLE_ID}
}

func (c *Rubble) Position() (int, int) {
	return 0, 0
}

func (c *Rubble) Pixels() []Pixel {
	p := []Pixel{}
	for i := range c.pixels {
		for j := range c.pixels[i] {
			if c.pixels[i][j] {
				p = append(p, Pixel{i, j, 'X'})
			}
		}
	}
	return p
}

func (c *Rubble) NewPosition(x, y int) {
	// doesn't move
}

func (c *Rubble) Rotate() {
	// redundant
}

func (c *Rubble) RotateBack() {
	// redundant
}

func (c *Rubble) AddPixels(pixels []Pixel) {
	for _, p := range pixels {
		c.pixels[p.X][p.Y] = true
	}
}

func (c *Rubble) Id() int {
	return c.id
}
