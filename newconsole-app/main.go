package main

//this programm will take a string, and then make a menu of options that the user can
//choose options from, and the program will repeat the options
//back to the screen.

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	//create a map
	coffees := make(map[int]string)
	coffees[1] = "Cappuchino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("____")
	fmt.Println("1 - Cappuchino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	char := ' '

	//listens to single key press
	for char != 'q' && char != 'Q' {
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		//rune data type
		i, _ := strconv.Atoi(string(char))

		_, ok := coffees[i]
		if ok {
			fmt.Println(fmt.Sprintf("you chose %s", coffees[i]))
		}
		//print information out
		//replacing the value using %d = int %s = string  %q = char

	}

	fmt.Println("Program Exiting...")
}
