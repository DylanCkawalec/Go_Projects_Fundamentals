package main

/*
Description: this program will repeat everything you type,
and end when you press quit, or q.

Author bp.dk



*/
import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune //used to build a string

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Print any key or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}() //annonomous function has this at the end

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}

}
