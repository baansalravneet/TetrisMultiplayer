package component

type GameOver struct {}

func NewGameOver() *GameOver {
	return &GameOver{}
}

func (c *GameOver) Position() (int, int) {
	return 1, 1
}

func (c *GameOver) Pixels() []Pixel {
	pixels := []Pixel{}
	for i := 0; i < 20; i++ {
		for j, c := range " GAMEOVER " {
			pixels = append(pixels, Pixel{i, j, c})
		}
	}
	return pixels
}

func (c *GameOver) NewPosition(x, y int) {
	// you cannot change the position of a border
}

func (c *GameOver) Rotate() {
	// you cannot rotate the border
}

func (c *GameOver) RotateBack() {
	// you cannot rotate the border
}
