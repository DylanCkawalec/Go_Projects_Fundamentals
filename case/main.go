package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "This is a clear EXAMPLE of why we search in one case only"

	searchString := strings.ToLower(myString)
	//wanna search this ^

	//we typically can convert characters to make this simplier.
	if strings.Contains(searchString, "this") {
		fmt.Println("Found it !")
	} else {
		fmt.Println("did not find it")
	}

	//other case functions

	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.Title(strings.ToLower(myString)))

}
