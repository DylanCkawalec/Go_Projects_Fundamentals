package main

import (
	"bufio"
	"fmt"
	"myapp/mylogger"
	"os"
	"time"
)

func main() {
	//read input form the user 5 times and write it to a log

	reader := bufio.NewReader(os.Stdin) //make a reader, much like a scanner
	ch := make(chan string)             //make a channel

	go mylogger.ListenForLog(ch) //use a go routine for infinate loop

	fmt.Println("Enter something") //instruction

	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n') //read the input, ignore the error
		ch <- input                         //send it to the channel named ch

		time.Sleep(time.Second) //wait because the loop does execute again

		//the go routine and the prompt would happen at the same time
		// so you need to add a time varient for the go routine
		//to finish executing and for the prompt to happen 1 second later

	}
}
