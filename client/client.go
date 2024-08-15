package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"tetris/gamestate"
	"tetris/input"
	"tetris/screen"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	screen := screen.Init()

	c := connectToServer()
	defer c.Close()

	inputChan := input.Start(100)
	defer input.Stop()

	clientLoop(screen, inputChan, c)

}

func clientLoop(screen screen.Screen, inputChan chan rune, c *websocket.Conn) {
	screenTicker := time.NewTicker(gamestate.SCREEN_REFRESH_RATE)
	state, err := getNewState(c)
	for err == nil {
		select {
		case <-screenTicker.C:
			screen.Update(*state)
		case input := <-inputChan:
			sendInput(c, input)
		default:
			state, err = getNewState(c)
		}
	}

}

func sendInput(c *websocket.Conn, input rune) {
	// ignore all errors
	c.WriteMessage(websocket.TextMessage, []byte(string(input)))
}

func getNewState(c *websocket.Conn) (*gamestate.GameState, error) {
	var state *gamestate.GameState
	_, message, err := c.ReadMessage()
	if err != nil {
		fmt.Println("Server sent an error", err)
		return nil, err
	}
	err = json.Unmarshal(message, &state)
	return state, err
}

func connectToServer() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/connect"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("Error connecting to the server", err)
		panic(err)
	}
	return c
}
