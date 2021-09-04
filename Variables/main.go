package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = " and Don't type you number in, just press ENTER when ready. "

func main() {

	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2

	var secondNumber = rand.Intn(8) + 2

	var subtraction = rand.Intn(8) + 2
	//value 0 is initiated
	var answer = firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)

}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)
	// display a welcome / instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("_____________________")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')
	//then I'll take the through the games

	fmt.Println("Multiply your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now, Multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the numebr you Originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract,", subtraction, prompt)
	reader.ReadString('\n')
	//give them the answer

	fmt.Println("the answer is", answer)
}
