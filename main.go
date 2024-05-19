package main

import (
	"tetris/renderer"
)

func main() {
	var screen renderer.Screen
	screen.Init()
	screen.Render()
}
