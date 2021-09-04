package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

//type for a variable
type Cat struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user Cat

	user.UserName = readString("what is your name?")

	user.Age = readInt("How old are you?")

	user.FavoriteNumber = readFloat("What is your favorite number ? ")

	user.OwnsADog = readBool("Do you own a dog (y/n) ? ")

	fmt.Println(fmt.Sprintf("your name is %s. You are %d years old. Your favorite number is %.4f. owns a dog: %t .",
		user.UserName,
		user.Age,
		user.FavoriteNumber, user.OwnsADog))
}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	for {

		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("please enter a whole number")
		} else {
			return num
		}

	}

}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)

		if err != nil {
			fmt.Println("please enter a number")
		} else {
			return num
		}

	}

}

func readBool(s string) bool {
	//listen for a single keypress
	err := keyboard.Open()
	//if there's an error , then log the error out
	if err != nil {
		log.Fatal(err)
	}
	//close the keyboard immediatly after opening it
	defer func() {
		_ = keyboard.Close() // _ means no error

	}()
	//print the question
	for { //
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}

	}

}
