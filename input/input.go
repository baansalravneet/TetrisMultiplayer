package input

import (
	"github.com/eiannone/keyboard"
)

func Start() chan rune {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	inputChan := make(chan rune)
	go func() {
		for {
			char, _, _ := keyboard.GetKey()
			inputChan <- char
		}
	}()
	return inputChan
}

func Stop() {
	keyboard.Close()
}
