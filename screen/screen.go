package screen

import (
	"fmt"
	"sync"
	"tetris/component"
	"time"
)

type Screen struct {
	Height     int
	Width      int
	pixels     [][]rune
	components map[int]component.Component
	mu         sync.RWMutex
}

func (s *Screen) render() {
	s.clear()
	for _, c := range s.components {
		cx, cy := c.Position()
		for _, pixel := range c.Pixels() {
			x := cx + pixel.X
			y := cy + pixel.Y
			s.update(pixel.C, x, y)
		}
	}
	for _, row := range (*s).pixels {
		for _, v := range row {
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

func (s *Screen) clear() {
	fmt.Printf("\033[%dA", s.Height)
	for i := range s.pixels {
		for j := range s.pixels[i] {
			s.pixels[i][j] = ' '
		}
	}
}

func (s *Screen) update(newChar rune, x, y int) {
	(*s).pixels[x][y] = newChar
}

func (s *Screen) MoveComponent(cId int, dir int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.components[cId]; !ok {
		return false
	}

	c := s.components[cId]
	cx, cy := c.Position()

	switch dir {
	case 1:
		c.NewPosition(cx, cy+1)
	case 2:
		c.NewPosition(cx+1, cy)
	case 3:
		c.NewPosition(cx, cy-1)
	}

	if !s.checkBoundsAndCollisions(c) {
		c.NewPosition(cx, cy)
		if dir == 2 {
			// component collided with something while moving down
			// change this to rubble
			s.changeToRubble(c)
			return false
		}
	}
	return true
}

func (s *Screen) changeToRubble(c component.Component) {
	rubble := s.components[component.RUBBLE_ID].(*component.Rubble)
	rubble.AddPixels(component.AbsolutePixels(c))
	delete(s.components, c.Id())
	updateRubble(rubble)
}

func (s *Screen) DropComponent(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c := s.components[id]
	cx, cy := c.Position()

	for s.checkBoundsAndCollisions(c) {
		c.NewPosition(cx+1, cy)
		cx++
	}
	c.NewPosition(cx-1, cy)
	s.changeToRubble(c)
}

func (s *Screen) RotateComponent(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.components[id]; !ok {
		return
	}
	c := s.components[id]
	c.Rotate()
	if !s.checkBoundsAndCollisions(c) {
		c.RotateBack()
	}
}

func (s *Screen) AddComponent(c component.Component) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.checkBoundsAndCollisions(c) {
		return false
	}
	cId := c.Id()
	s.components[cId] = c
	if cId >= 2 && cId <= 7 {
		go func() {
			ticker := time.NewTicker(1 * time.Second)
			for range ticker.C {
				if !s.MoveComponent(cId, 2) {
					ticker.Stop()
				}
			}
		}()
	}
	return true
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
	s.AddComponent(component.NewBorder(0, 0, s.Height-1, s.Width-1))
	s.AddComponent(component.NewRubble(s.Height, s.Width))
	return s
}

func newScreen() *Screen {
	s := &Screen{
		Height:     16,
		Width:      10,
		components: make(map[int]component.Component),
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

func (s *Screen) ActiveComponent() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := 2; i <= 7; i++ {
		if _, ok := s.components[i]; ok {
			return i
		}
	}
	return 0
}

func updateRubble(rubble *component.Rubble) {
	pixels := rubble.GetPixels()
	deleteLines := []int{}
	for i := range pixels {
		if lineFull(pixels[i]) {
			deleteLines = append(deleteLines, i)
		}
	}
	rubble.Delete(deleteLines)
}

func lineFull(pixels []bool) bool {
	count := 0
	for _, v := range pixels {
		if v {
			count++
		}
	}
	return count == len(pixels)-2
}
