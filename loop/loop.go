package loop

import (
	"tetris/gamestate"
	"tetris/screen"
	"time"
)

func Loop(inputChan chan rune, s *screen.Screen, state *gamestate.GameState) {
	ticker := time.NewTicker(gamestate.DROP_SPEED)
	for !state.Gameover && !state.GameExit {
		if state.CurrentPiece == nil {
			if !state.AddPiece() {
				s.GameOver()
				state.Gameover = true
			}
			ticker.Reset(gamestate.DROP_SPEED)
		}
		select {
		case input := <-inputChan:
			state.HandleInput(input)
		case <-ticker.C:
			state.HandleDrop()
		}
	}
}
