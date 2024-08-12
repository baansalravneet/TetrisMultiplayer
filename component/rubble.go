package component

type Rubble struct {
	P [][]bool
}

func NewRubble(height, width int) Rubble {
	pixels := make([][]bool, height)
	for i := range pixels {
		pixels[i] = make([]bool, width)
	}
	return Rubble{pixels}
}

func (c *Rubble) Position() (int, int) {
	return 0, 0
}

func (c *Rubble) Pixels() []Pixel {
	p := []Pixel{}
	for i := range c.P {
		for j := range c.P[i] {
			if c.P[i][j] {
				p = append(p, Pixel{i, j, 'X'})
			}
		}
	}
	return p
}

func (c *Rubble) AddPixels(pixels []Pixel) {
	for _, p := range pixels {
		c.P[p.X][p.Y] = true
	}
}

func (c *Rubble) GetPixels() [][]bool {
	return c.P
}

func (c *Rubble) Delete(lines []int) {
	if len(lines) == 0 {
		return
	}
	i, j, count := len(lines)-1, len(c.P)-1, 0
	for j-count >= 0 {
		if i >= 0 && j-count == lines[i] {
			count++
			i--
			copy(c.P[j], c.P[j-count])
		} else {
			copy(c.P[j], c.P[j-count])
			j--
		}
	}
}

func (c *Rubble) Contains(x, y int) bool {
	return c.P[x][y]
}
