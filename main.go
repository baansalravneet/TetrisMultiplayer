package main

import (
	"tetris/input"
	"tetris/loop"
	"tetris/screen"
)

func main() {
	s := screen.Start()
	inputChan := input.Start()
	loop.Loop(inputChan, s)
	input.Stop()
}
