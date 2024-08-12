package pieces

import (
	"math/rand/v2"
	"tetris/component"
)

type Piece interface {
	Position() (int, int)
	Pixels() []component.Pixel
	NewPosition(int, int)
	Rotate()
	RotateBack()
}

func AbsolutePixels(c Piece) []component.Pixel {
	p := []component.Pixel{}
	x, y := c.Position()
	for _, i := range c.Pixels() {
		p = append(p, component.Pixel{X: x + i.X, Y: y + i.Y, C: i.C})
	}
	return p
}

func NewRandomPiece() Piece {
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
