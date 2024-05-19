package gameloop

import "tetris/renderer"

func Loop() {
	screen := renderer.Init()
	screen.Render()
}
