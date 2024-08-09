package screen

import (
	"fmt"
	"tetris/component"
	"tetris/gamestate"
	"time"
)

const screenHeight = 22
const screenWidth = 20

type Screen struct {
	pixels   [][]rune
	state    *gamestate.GameState
	border   component.Component
	npw      component.Component
	nextText component.Component
}

func (s *Screen) render() {
	s.clear()
	s.updateScreen()
	s.print()
}

func (s *Screen) updateScreen() {
	cx, cy := s.state.CurrentPiece.Position()
	for _, pixel := range s.state.CurrentPiece.Pixels() {
		s.update(pixel.C, cx+pixel.X, cy+pixel.Y)
	}
	cx, cy = 3, 16
	for _, pixel := range s.state.NextPiece.Pixels() {
		s.update(pixel.C, cx+pixel.X, cy+pixel.Y)
	}
	for _, pixel := range s.state.Rubble.Pixels() {
		s.update(pixel.C, pixel.X, pixel.Y)
	}
	if s.state.Gameover {
		gameoverComponent := component.NewGameOver()
		cx, cy = gameoverComponent.Position()
		for _, pixel := range gameoverComponent.Pixels() {
			x := cx + pixel.X
			y := cy + pixel.Y
			s.update(pixel.C, x, y)
		}
	}
}

func (s *Screen) print() {
	for _, row := range (*s).pixels {
		for _, v := range row {
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

func (s *Screen) clear() {
	fmt.Printf("\033[%dA", screenHeight)
	// clear the game board
	c := s.border.(*component.Border)
	for i := c.Top + 1; i < c.Bottom; i++ {
		for j := c.Left + 1; j < c.Right; j++ {
			s.update(' ', i, j)
		}
	}
	// clear npw
	c = s.npw.(*component.Border)
	for i := c.Top + 1; i < c.Bottom; i++ {
		for j := c.Left + 1; j < c.Right; j++ {
			s.update(' ', i, j)
		}
	}
}

func (s *Screen) update(newChar rune, x, y int) {
	(*s).pixels[x][y] = newChar
}

func (s *Screen) GameOver() {
	s.render()
}

func Start(state *gamestate.GameState) *Screen {
	s := newScreen()
	s.state = state
	s.border = component.NewBorder(0, 0, gamestate.BOARD_HEIGHT-1, gamestate.BOARD_WIDTH-1)
	s.npw = component.NewBorder(2, 14, 5, 19)
	s.nextText = component.NewText("NEXT", 1, 15)
	s.addStaticComponents()
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for range ticker.C {
			s.render()
		}
	}()
	return s
}

func newScreen() *Screen {
	s := &Screen{}
	s.pixels = make([][]rune, screenHeight)
	for i := range s.pixels {
		s.pixels[i] = make([]rune, screenWidth)
		for j := range s.pixels[i] {
			s.pixels[i][j] = ' '
		}
	}
	return s
}

func (s *Screen) addStaticComponents() {
	c := s.border
	cx, cy := c.Position()
	for _, pixel := range c.Pixels() {
		x := cx + pixel.X
		y := cy + pixel.Y
		s.update(pixel.C, x, y)
	}
	c = s.npw
	cx, cy = c.Position()
	for _, pixel := range c.Pixels() {
		x := cx + pixel.X
		y := cy + pixel.Y
		s.update(pixel.C, x, y)
	}
	c = s.nextText
	cx, cy = c.Position()
	for _, pixel := range c.Pixels() {
		x := cx + pixel.X
		y := cy + pixel.Y
		s.update(pixel.C, x, y)
	}
}
