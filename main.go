package main

import (
	"tetris/gameloop"
	"tetris/screen"
)

func main() {
	inputChan, s := screen.Start()
	gameloop.Loop(inputChan, s)
}
