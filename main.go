package main

import (
	"tetris/input"
	"tetris/loop"
	"tetris/screen"
)

func main() {
	s := screen.Start()
	inputChan := input.Start()
	defer input.Stop()
	loop.Loop(inputChan, s)
}
