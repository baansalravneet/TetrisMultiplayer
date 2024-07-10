package main

import (
	"tetris/input"
	"tetris/loop"
	"tetris/screen"
)

func main() {
	// starting the rendering service
	s := screen.Start()

	// initializing the buffered input channel
	inputChan := input.Start(100)
	defer input.Stop()
	
	// start the loop
	loop.Loop(inputChan, s)
}
