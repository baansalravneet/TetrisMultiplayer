package component

import "math/rand/v2"

const BORDER_ID = 0
const RUBBLE_ID = 1

const BAR_ID = 2
const BOX_ID = 3
const LEFT_L_ID = 4
const LEFT_Z_ID = 5
const RIGHT_L_ID = 6
const RIGHT_Z_ID = 7
const TEE_ID = 8

const GAME_OVER_ID = 9

type Pixel struct {
	X int
	Y int
	C rune
}

type Component interface {
	Position() (int, int)
	Pixels() []Pixel
	NewPosition(int, int)
	Rotate()
	RotateBack()
	Id() int
}

func AbsolutePixels(c Component) []Pixel {
	p := []Pixel{}
	x, y := c.Position()
	for _, i := range c.Pixels() {
		p = append(p, Pixel{x + i.X, y + i.Y, i.C})
	}
	return p
}

func NewRandomComponent() Component {
	c := rand.IntN(7)
	switch c {
	case 0:
		return NewBar()
	case 1:
		return NewBox()
	case 2:
		return NewLeftL()
	case 3:
		return NewLeftZ()
	case 4:
		return NewRightL()
	case 5:
		return NewRightZ()
	default:
		return NewTee()
	}
}
