package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var player1Address = flag.String("player1Address", "0.0.0.0:8080", "http service address")
var player2Address = flag.String("player2Address", "0.0.0.0:8081", "http service address")

type GameServer struct {
	Player1Connection *websocket.Conn
	Player2Connection *websocket.Conn
}

func New() *GameServer {
	gs := &GameServer{}
	http.HandleFunc("/connect", gs.connect)
	go func() {
		http.ListenAndServe(*player1Address, nil)
	}()
	go func() {
		http.ListenAndServe(*player2Address, nil)
	}()
	fmt.Println("Connect Player 1")
	// TODO: find a better way to wait for the players
	for gs.Player1Connection == nil {
	}
	fmt.Println("Player 1 connected!")
	fmt.Println("Connect Player 2")
	for gs.Player2Connection == nil {
	}
	fmt.Println("Player 2 connected!")
	return gs
}

func (gs *GameServer) GetPlayer1Input(buffer int) chan rune {
	inputChan := make(chan rune, buffer)
	go func() {
		for {
			_, message, err := gs.Player1Connection.ReadMessage()
			if err != nil {
				fmt.Println("Error reading client input")
				break
			}
			inputChan <- rune(string(message)[0])
		}
	}()
	return inputChan
}

func (gs *GameServer) GetPlayer2Input(buffer int) chan rune {
	inputChan := make(chan rune, buffer)
	go func() {
		for {
			_, message, err := gs.Player2Connection.ReadMessage()
			if err != nil {
				fmt.Println("Error reading client input")
				break
			}
			inputChan <- rune(string(message)[0])
		}
	}()
	return inputChan
}

func (gs *GameServer) connect(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
	}
	if gs.Player1Connection == nil {
		gs.Player1Connection = c
	} else {
		gs.Player2Connection = c
	}
}
