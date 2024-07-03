package screen

import (
	"bufio"
	"fmt"
	"os"
	"tetris/component"
	"time"
)

type Screen struct {
	pixels     [][]rune
	components map[int]*component.Component
}

func (s *Screen) render() {
	s.Clear()
	for _, c := range s.components {
		coords := (*c).Position()
		for _, relativeCoords := range (*c).Pixels() {
			x := coords.X + relativeCoords.X
			y := coords.Y + relativeCoords.Y
			s.Update((*c).Char(), x, y)
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
	fmt.Printf("\033[%dA", screenHeight+1)
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
	coords := c.Position()
	switch dir {
	case 0:
		if coords.X > 0 {
			c.NewPosition(coords.X-1, coords.Y)
		}
	case 1:
		if coords.Y < screenWidth-1 {
			c.NewPosition(coords.X, coords.Y+1)
		}
	case 2:
		if coords.X < screenHeight-1 {
			c.NewPosition(coords.X+1, coords.Y)
		}
	case 3:
		if coords.Y > 0 {
			c.NewPosition(coords.X, coords.Y-1)
		}
	}
	s.components[id] = &c
}

func (s *Screen) AddComponent(c component.Component) int {
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

const screenHeight = 5
const screenWidth = 5

func Start() (chan rune, *Screen) {
	s := newScreen()
	inputs := make(chan rune)
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for range ticker.C {
			s.render()
		}
	}()
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			key := readKey(reader)
			inputs <- key
		}
	}()
	return inputs, &s
}

func readKey(reader *bufio.Reader) rune {
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading key: ", err)
	}
	return char
}

func newScreen() Screen {
	s := Screen{
		pixels:     make([][]rune, screenHeight),
		components: make(map[int]*component.Component),
	}
	for i := range s.pixels {
		s.pixels[i] = make([]rune, screenWidth)
		for j := range s.pixels[i] {
			s.pixels[i][j] = ' '
		}
	}
	return s
}
