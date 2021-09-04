package main

//everything has to have these imported.
import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
) //these are all the different pakcages that the program is importing into the file

func main() {
	//this is reading the user input from the screen
	reader := bufio.NewReader(os.Stdin)
	//this is declaring the variable as one of the functions from another package
	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		//these checks for different OS's
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {

			fmt.Println(doctor.Response(userInput))
		}

	}

}
