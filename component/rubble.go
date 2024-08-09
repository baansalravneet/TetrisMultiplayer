package component

type Rubble struct {
	pixels [][]bool
}

func NewRubble(height, width int) *Rubble {
	pixels := make([][]bool, height)
	for i := range pixels {
		pixels[i] = make([]bool, width)
	}
	return &Rubble{pixels}
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

func (c *Rubble) GetPixels() [][]bool {
	return c.pixels
}

func (c *Rubble) Delete(lines []int) {
	if len(lines) == 0 {
		return
	}
	i, j, count := len(lines)-1, len(c.pixels)-1, 0
	for j-count >= 0 {
		if i >= 0 && j-count == lines[i] {
			count++
			i--
			copy(c.pixels[j], c.pixels[j-count])
		} else {
			copy(c.pixels[j], c.pixels[j-count])
			j--
		}
	}
}

func (c *Rubble) Contains(x, y int) bool {
	return c.pixels[x][y]
}
