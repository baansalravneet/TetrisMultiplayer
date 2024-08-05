package component

type GameOver struct {
	id int
}

func NewGameOver() *GameOver {
	return &GameOver{GAME_OVER_ID}
}

func (c *GameOver) Position() (int, int) {
	return 1, 1
}

func (c *GameOver) Pixels() []Pixel {
	pixels := []Pixel{}
	for i := 0; i < 14; i++ {
		for j, c := range "GAMEOVER" {
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

func (c *GameOver) Id() int {
	return c.id
}
