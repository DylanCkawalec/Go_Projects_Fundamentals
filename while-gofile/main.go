package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano()) //#1

	i := 1000 //#2

	//execute a loop while i > 100
	for i > 100 {
		// get a random number between 1 and 1001
		i = rand.Intn(1000) + 1 // you add 1 because go panics if the number is 0
		fmt.Println("i is : ", i)
		if i > 100 {
			fmt.Println("i is ", i, " so the loop keep's going")
		}
	}

	fmt.Println("Got ", i, "and broke out of loop. ")
}
