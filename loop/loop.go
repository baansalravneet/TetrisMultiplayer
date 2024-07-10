package loop

import (
	"errors"
	"fmt"
	"tetris/component"
	"tetris/screen"
)

func Loop(inputChan chan rune, s *screen.Screen) {
GAME_LOOP:
	for {
		var cId int
		if cId = s.ActiveComponent(); cId == 0 {
			c := component.NewRandomComponent()
			s.AddComponent(c)
			cId = c.Id()
		}
		err := handleInputs(cId, inputChan, s)
		if err != nil {
			fmt.Println(err)
			break GAME_LOOP
		}
	}
}

func handleInputs(cId int, inputChan chan rune, s *screen.Screen) error {
	for i := 0; i < len(inputChan); i++ {
		switch <-inputChan {
		case 'd':
			s.MoveComponent(cId, 1)
		case 's':
			s.MoveComponent(cId, 2)
		case 'a':
			s.MoveComponent(cId, 3)
		case 'j':
			s.RotateComponent(cId)
		case 'k':
			s.DropComponent(cId)
		case 'x':
			return errors.New("stoping game loop")
		}
	}
	return nil
}
