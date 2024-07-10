package input

import (
	"github.com/eiannone/keyboard"
)

func Start(buffer int) chan rune {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	inputChan := make(chan rune, buffer)
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
