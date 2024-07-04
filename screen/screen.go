package screen

import (
	"fmt"
	"tetris/component"
	"time"
)

type Screen struct {
	Height     int
	Width      int
	pixels     [][]rune
	components map[int]*component.Component
}

func (s *Screen) render() {
	s.Clear()
	for _, c := range s.components {
		cx, cy := (*c).Position()
		for _, pixel := range (*c).Pixels() {
			x := cx + pixel.X
			y := cy + pixel.Y
			s.Update(pixel.C, x, y)
		}
	}
	for _, row := range (*s).pixels {
		for _, v := range row {
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

func (s *Screen) Clear() {
	fmt.Printf("\033[%dA", s.Height)
	for i := range s.pixels {
		for j := range s.pixels[i] {
			s.pixels[i][j] = ' '
		}
	}
}

func (s *Screen) Update(newChar rune, x, y int) {
	(*s).pixels[x][y] = newChar
}

func (s *Screen) MoveComponent(id int, dir int) {
	c := *s.components[id]
	cx, cy := c.Position()
	switch dir {
	case 0:
		c.NewPosition(cx-1, cy)
	case 1:
		c.NewPosition(cx, cy+1)
	case 2:
		c.NewPosition(cx+1, cy)
	case 3:
		c.NewPosition(cx, cy-1)
	}
	if !s.checkBoundsAndCollisions(c) {
		c.NewPosition(cx, cy)
	}
	s.components[id] = &c
}

func (s *Screen) AddComponent(c component.Component) int {
	if !s.checkBoundsAndCollisions(c) {
		return -1
	}
	findFreeId := func() int {
		m := (*s).components
		i := 0
		for i < len(m) {
			if _, ok := m[i]; !ok {
				return i
			}
			i++
		}
		return i
	}
	id := findFreeId()
	s.components[id] = &c
	return id
}

// two components will not collide if the pixels are the same
// possible merge
func (s *Screen) checkBoundsAndCollisions(c component.Component) bool {
	x, y := c.Position()
	for _, pixel := range c.Pixels() {
		if x+pixel.X < 0 || x+pixel.X >= s.Height ||
			y+pixel.Y < 0 || y+pixel.Y >= s.Width {
			return false
		}
		if s.pixels[x+pixel.X][y+pixel.Y] != ' ' &&
			s.pixels[x+pixel.X][y+pixel.Y] != pixel.C {
			return false
		}
	}
	return true
}

func Start() *Screen {
	s := newScreen()
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for range ticker.C {
			s.render()
		}
	}()
	return &s
}

func newScreen() Screen {
	s := Screen{
		Height:     10,
		Width:      10,
		components: make(map[int]*component.Component),
	}
	s.pixels = make([][]rune, s.Height)
	for i := range s.pixels {
		s.pixels[i] = make([]rune, s.Width)
		for j := range s.pixels[i] {
			s.pixels[i][j] = ' '
		}
	}
	return s
}
