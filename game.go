package main

import (
	"fmt"
	"tetris/gamestate"
	"tetris/loop"
	"tetris/server"
)

func main() {
	fmt.Println("Initializing game state...")
	state := gamestate.Init()

	fmt.Println("Initializing server...")
	gs := server.New()
	defer gs.Close()

	fmt.Println("Initializing controls...")
	clientInput := gs.GetInputChannel(100)

	fmt.Println("Starting Game...")
	loop.Loop(state, clientInput, gs)
}
