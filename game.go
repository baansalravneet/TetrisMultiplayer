package main

import (
	"tetris/gamestate"
	"tetris/input"
	"tetris/loop"
	"tetris/screen"
)

func main() {
	state := gamestate.Init()
	s := screen.Start()
	inputChan := input.Start(100)
	defer input.Stop()
	loop.Loop(inputChan, s, state)
}
