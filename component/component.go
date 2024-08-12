package component

type Pixel struct {
	X int
	Y int
	C rune
}

type Component interface {
	Position() (int, int)
	Pixels() []Pixel
}

func AbsolutePixels(c Component) []Pixel {
	p := []Pixel{}
	x, y := c.Position()
	for _, i := range c.Pixels() {
		p = append(p, Pixel{x + i.X, y + i.Y, i.C})
	}
	return p
}
