package main

import (
	"tetris/gameloop"
	"tetris/screen"
	"tetris/input"
)

func main() {
	s := screen.Start()
	inputChan := input.Start()
	gameloop.Loop(inputChan, s)
	input.Stop()
}
