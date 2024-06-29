package gameloop

import (
	"fmt"
	"tetris/component"
	"tetris/screen"
)

func Loop(inputChan chan rune, s *screen.Screen) {
	dotId := s.AddComponent(component.NewDot())
GAME_LOOP:
	for {
		input := <-inputChan
		switch input {
		case 'w':
			s.MoveComponent(dotId, 0)
		case 'd':
			s.MoveComponent(dotId, 1)
		case 's':
			s.MoveComponent(dotId, 2)
		case 'a':
			s.MoveComponent(dotId, 3)
		case 'x':
			fmt.Println("Stopping game")
			break GAME_LOOP
		}
	}
}
