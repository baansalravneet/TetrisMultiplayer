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
