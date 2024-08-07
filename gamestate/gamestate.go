package gamestate

import (
	"fmt"
	"tetris/component"
	"time"
)

const SCREEN_HEIGHT int = 16
const SCREEN_WIDTH int = 10
const DROP_SPEED time.Duration = 1000 * time.Millisecond

type GameState struct {
	Gameover     bool
	GameExit     bool
	CurrentPiece component.Component
	NextPiece    component.Component
	Rubble       *component.Rubble
}

func Init() *GameState {
	return &GameState{
		NextPiece: component.NewRandomPiece(),
		Rubble:    component.NewRubble(SCREEN_HEIGHT, SCREEN_WIDTH),
	}
}

func (s *GameState) AddPiece() bool {
	s.CurrentPiece = s.NextPiece
	s.NextPiece = component.NewRandomPiece()
	return !s.collided()
}

func (s *GameState) HandleDrop() {
	s.moveComponent(2)
}

func (s *GameState) HandleInput(input rune) {
	switch input {
	case 'd':
		s.moveComponent(1)
	case 's':
		s.moveComponent(2)
	case 'a':
		s.moveComponent(3)
	case 'j':
		s.rotateComponent()
	case 'k':
		s.dropComponent()
	case 'x':
		fmt.Println("Exiting game...")
		s.GameExit = true
	}
}

func (s *GameState) dropComponent() {
	cx, cy := s.CurrentPiece.Position()
	for !s.collided() {
		s.CurrentPiece.NewPosition(cx+1, cy)
		cx++
	}
	s.CurrentPiece.NewPosition(cx-1, cy)
	s.changeToRubble()
}

func (s *GameState) rotateComponent() {
	c := s.CurrentPiece
	c.Rotate()
	if s.collided() {
		c.RotateBack()
	}
}

func (s *GameState) moveComponent(dir int) bool {
	c := s.CurrentPiece
	cx, cy := c.Position()

	switch dir {
	case 1:
		c.NewPosition(cx, cy+1)
	case 2:
		c.NewPosition(cx+1, cy)
	case 3:
		c.NewPosition(cx, cy-1)
	}

	if s.collided() {
		c.NewPosition(cx, cy)
		if dir == 2 {
			// component collided with something while moving down
			// change this to rubble
			s.changeToRubble()
			return false
		}
	}
	return true
}

func (s *GameState) collided() bool {
	c := s.CurrentPiece
	x, y := c.Position()
	for _, pixel := range c.Pixels() {
		if x+pixel.X < 1 || x+pixel.X >= SCREEN_HEIGHT-1 ||
			y+pixel.Y < 1 || y+pixel.Y >= SCREEN_WIDTH-1 {
			return true
		}
		if s.Rubble.Contains(x+pixel.X, y+pixel.Y) {
			return true
		}
	}
	return false
}

func (s *GameState) changeToRubble() {
	rubble := s.Rubble
	rubble.AddPixels(component.AbsolutePixels(s.CurrentPiece))
	s.CurrentPiece = nil
	s.updateRubble()
}

func (s *GameState) updateRubble() {
	pixels := s.Rubble.GetPixels()
	deleteLines := []int{}
	for i := range pixels {
		if lineFull(pixels[i]) {
			deleteLines = append(deleteLines, i)
		}
	}
	s.Rubble.Delete(deleteLines)
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
