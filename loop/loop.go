package loop

import (
	"encoding/json"
	"tetris/gamestate"
	"time"

	"github.com/gorilla/websocket"
)

func Loop(state, opState *gamestate.GameState, inputChan chan rune, c *websocket.Conn) {
	gameTicker := time.NewTicker(gamestate.DROP_SPEED)
	connectionTicker := time.NewTicker(100 * time.Millisecond) // TODO: remove this and send state for every state update
	for !state.Gameover && !state.GameExit && !opState.Gameover && !opState.GameExit {
		if state.CurrentPiece == nil {
			if !state.AddPiece() {
				state.Gameover = true
			}
			gameTicker.Reset(gamestate.DROP_SPEED)
		}
		select {
		case input := <-inputChan:
			state.HandleInput(input)
		case <-gameTicker.C:
			state.HandleDrop()
		case <-connectionTicker.C:
			// TODO: make this async and do it for every state update
			stateData := []gamestate.GameState{*state, *opState}
			stateString, _ := json.Marshal(stateData)
			c.WriteMessage(websocket.TextMessage, stateString)
		}
	}
}
