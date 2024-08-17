package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"tetris/gamestate"
	"tetris/input"
	"tetris/screen"
	"time"

	"github.com/gorilla/websocket"
)

var playerAddress *string

func main() {

	player := os.Args[0]
	if player == "1" {
		playerAddress = flag.String("player1Address", "localhost:8080", "http service address")
	} else {
		playerAddress = flag.String("player2Address", "localhost:8081", "http service address")
	}

	screen := screen.Init()

	c := connectToServer()
	defer c.Close()

	inputChan := input.Start(100)
	defer input.Stop()

	clientLoop(screen, inputChan, c)

}

func clientLoop(screen screen.Screen, inputChan chan rune, c *websocket.Conn) {
	screenTicker := time.NewTicker(gamestate.SCREEN_REFRESH_RATE)
	state, opState, err := getNewState(c)
	for err == nil {
		select {
		case <-screenTicker.C:
			screen.Update(*state, *opState)
		case input := <-inputChan:
			sendInput(c, input)
		default:
			state, opState, err = getNewState(c)
		}
	}

}

func sendInput(c *websocket.Conn, input rune) {
	// ignore all errors
	c.WriteMessage(websocket.TextMessage, []byte(string(input)))
}

func getNewState(c *websocket.Conn) (*gamestate.GameState, *gamestate.GameState, error) {
	var state []gamestate.GameState
	_, message, err := c.ReadMessage()
	if err != nil {
		fmt.Println("Server sent an error", err)
		return nil, nil, err
	}
	err = json.Unmarshal(message, &state)
	return &state[0], &state[1], err
}

func connectToServer() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: *playerAddress, Path: "/connect"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("Error connecting to the server", err)
		panic(err)
	}
	return c
}
