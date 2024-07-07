package component

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
}
