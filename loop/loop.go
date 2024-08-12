package loop

import (
	"tetris/gamestate"
	"tetris/screen"
	"time"
)

func Loop(inputChan chan rune, s screen.Screen, state *gamestate.GameState) {
	// connect player 1 - WIP
	// connect player 2 - WIP
	// start game
	gameTicker := time.NewTicker(gamestate.DROP_SPEED)
	screenTicker := time.NewTicker(gamestate.SCREEN_REFRESH_RATE)
	for !state.Gameover && !state.GameExit {
		if state.CurrentPiece == nil {
			if !state.AddPiece() {
				s.GameOver(*state)
				state.Gameover = true
			}
			gameTicker.Reset(gamestate.DROP_SPEED)
		}
		select {
		case input := <-inputChan:
			state.HandleInput(input)
		case <-gameTicker.C:
			state.HandleDrop()
		case <-screenTicker.C:
			s.Update(*state)
		}
	}
}
