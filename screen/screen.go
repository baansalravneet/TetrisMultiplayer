package screen

import (
	"fmt"
	"time"
)

type Screen [][]rune

func (s *Screen) render() {
	for i := 0; i < len(*s); i++ {
		fmt.Printf("\033[F\033[K")
	}
	for _, row := range *s {
		for _, v := range row {
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

func (s *Screen) Update(newChar rune, x, y int) {
	(*s)[x][y] = newChar
}

const screenHeight = 5
const screenWidth = 5

func Start() (chan int, *Screen) {
	s := newScreen()
	ticker := time.NewTicker(100 * time.Millisecond)
	quit := make(chan int)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.render()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return quit, &s
}

func newScreen() Screen {
	s := make([][]rune, screenHeight)
	for i := range s {
		s[i] = make([]rune, screenWidth)
		for j := range s[i] {
			s[i][j] = ' '
		}
	}
	return s
}
