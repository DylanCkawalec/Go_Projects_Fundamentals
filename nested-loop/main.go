package main

import "fmt"

//debigging
func main() {
	for i := 1; i <= 10; i++ { //only executes 10 times
		fmt.Print("i is ", i)
		for j := 1; j <= 3; j++ {
			fmt.Print("   j: ", j) //executes 3 times every time the main loop occurs once
		}
		fmt.Println()
	}

}
