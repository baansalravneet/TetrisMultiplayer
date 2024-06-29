package gameloop

import (
	"tetris/screen"
	"time"
)

func Loop(quitChan chan int, s *screen.Screen) {
	for i := 0; i < 26; i++ {
		s.Update(rune(i+'A'), 0, 0)
		time.Sleep(200 * time.Millisecond)
	}
	quitChan <- 1
}
