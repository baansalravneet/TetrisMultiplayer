package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var player1Address = flag.String("addr", "localhost:8080", "http service address")

type GameServer struct {
	Player1Connection *websocket.Conn
}

func New() *GameServer {
	gs := &GameServer{}
	http.HandleFunc("/connect", gs.connect)
	go func() {
		http.ListenAndServe(*player1Address, nil)
	}()
	fmt.Println("Connect Player 1")
	// TODO: find a better way to wait for the players
	for gs.Player1Connection == nil {
	}
	fmt.Println("Player 1 connected!")
	return gs
}

func (gs *GameServer) GetInputChannel(buffer int) chan rune {
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

func (gs *GameServer) connect(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
	}
	gs.Player1Connection = c
}

func (gs *GameServer) Close() {
	gs.Player1Connection.Close()
}
