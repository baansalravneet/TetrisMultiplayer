package main

import (
	"fmt"
	"sync"
	"tetris/gamestate"
	"tetris/loop"
	"tetris/server"
)

func main() {
	fmt.Println("Initializing game state...")
	player1State := gamestate.Init()
	player2State := gamestate.Init()

	fmt.Println("Initializing server...")
	gs := server.New()

	fmt.Println("Initializing controls...")
	player1Input := gs.GetPlayer1Input(100)
	player2Input := gs.GetPlayer2Input(100)

	fmt.Println("Starting Game...")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		loop.Loop(player1State, player2State, player1Input, gs.Player1Connection)
		wg.Done()
	}()
	go func() {
		loop.Loop(player2State, player1State, player2Input, gs.Player2Connection)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All games finished")
}
