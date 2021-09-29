package main

import (
	"fmt"
	"strings"
)

func main() {
	//          1         2         3         4
	// 0123456789012345678901234567890123456789012345
	newString := "Go is a great programming language. Go for it!"

	fmt.Println()

	if strings.Contains(newString, "Go") {
		//newString = strings.Replace(newString, "Go", "GoLang", 1) //change this value
		newString = strings.ReplaceAll(newString, "Go", "GoLang") //replaces all instances

	}

	fmt.Println(newString)

	//string comparison

	if "Alpha" > "Absolute" {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("A is not greater than B") //will print this one // these are runes.
	}

	badEmail := " me@here.com  "
	//how to trim whitespace
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("=%s=", badEmail)
	fmt.Println()
}
