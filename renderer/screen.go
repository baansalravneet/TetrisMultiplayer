package renderer

import "fmt"

var TOP_LEFT = '╔'
var BOTTOM_LEFT = '╚'
var TOP_RIGHT = '╗'
var BOTTOM_RIGHT = '╝'
var HORIZONTAL = '═'
var VERTICAL = '║'
var BLANK = ' '

var SCREEN_WIDTH = 10
var SCREEN_HEIGHT = 10

type Screen struct {
	pixels [][]rune
}

func Init() *Screen {
	pixels := make([][]rune, SCREEN_HEIGHT)
	for i := range pixels {
		pixels[i] = make([]rune, SCREEN_WIDTH)
		for j := range pixels[i] {
			pixels[i][j] = BLANK
		}
	}
	pixels[0][0] = TOP_LEFT
	pixels[0][SCREEN_WIDTH-1] = TOP_RIGHT
	pixels[SCREEN_HEIGHT-1][0] = BOTTOM_LEFT
	pixels[SCREEN_HEIGHT-1][SCREEN_WIDTH-1] = BOTTOM_RIGHT
	for i := 1; i < SCREEN_WIDTH-1; i++ {
		pixels[0][i] = HORIZONTAL
		pixels[SCREEN_HEIGHT-1][i] = HORIZONTAL
	}
	for i := 1; i < SCREEN_HEIGHT-1; i++ {
		pixels[i][0] = VERTICAL
		pixels[i][SCREEN_WIDTH-1] = VERTICAL
	}
	return &Screen{pixels}
}

func (s *Screen) Render() {
	for _, line := range s.pixels {
		fmt.Println(string(line))
	}
}
