package main

import (
	"tetris/gameloop"
	"tetris/screen"
)

func main() {
	quitChan, s := screen.Start()
	gameloop.Loop(quitChan, s)
}
