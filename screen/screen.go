package screen

import (
	"fmt"
	"tetris/component"
	"tetris/gamestate"
	"time"
)

type Screen struct {
	pixels [][]rune
	state  *gamestate.GameState
}

func (s *Screen) render() {
	s.clear()
	cx, cy := s.state.CurrentPiece.Position()
	for _, pixel := range s.state.CurrentPiece.Pixels() {
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
	for _, row := range (*s).pixels {
		for _, v := range row {
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

func (s *Screen) clear() {
	fmt.Printf("\033[%dA", gamestate.SCREEN_HEIGHT)
	for i := 1; i < gamestate.SCREEN_HEIGHT-1; i++ {
		for j := 1; j < gamestate.SCREEN_WIDTH-1; j++ {
			s.pixels[i][j] = ' '
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
	s.addBorder(component.NewBorder(0, 0, gamestate.SCREEN_HEIGHT-1, gamestate.SCREEN_WIDTH-1))
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
	s.pixels = make([][]rune, gamestate.SCREEN_HEIGHT)
	for i := range s.pixels {
		s.pixels[i] = make([]rune, gamestate.SCREEN_WIDTH)
		for j := range s.pixels[i] {
			s.pixels[i][j] = ' '
		}
	}
	return s
}

func (s *Screen) addBorder(c component.Component) {
	cx, cy := c.Position()
	for _, pixel := range c.Pixels() {
		x := cx + pixel.X
		y := cy + pixel.Y
		s.update(pixel.C, x, y)
	}
}
