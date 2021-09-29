package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := []string{

		"Learn Go for Beginners Crash Course",
		"Learn Java for Beginners Crash Course",
		"Learn Solidity for Beginners Crash Course",
		"Learn Javascript for Beginners Crash Course",
		"Learn C for Beginners Crash Course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Solidity") {
			fmt.Println("'Solidity' is found in - '", x, "', and the index is - ", strings.Index(x, "Solidity"))
		}
	}

	fmt.Println(" - - - -- - -- - -- - - - ---- ---- -- - -")

	//                      1         2         3         4
	//            0123456789012345678901234567890123456789012345
	newString := "Go is a great programming language. Go for it!"
	fmt.Println(newString)
	fmt.Println("Does this line have 'Go' in it? : ", strings.HasPrefix(newString, "Go"))   // boolean
	fmt.Println("Does this line have Java in it? : ", strings.HasPrefix(newString, "Java")) //false
	fmt.Println("Does this line contain an exclamation mark? : ", strings.HasSuffix(newString, "!"))

	fmt.Println("How often does the word 'Go' occur in the sentence ? : ", strings.Count(newString, "Go"))
}
