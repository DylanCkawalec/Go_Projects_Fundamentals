package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()

	name := "Hello World"
	fmt.Println("String: ", name)

	fmt.Println()

	fmt.Println("Bytes -----------")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("INDEX\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	//using PLus (+) concatination not very efficient
	// using fmt is more efficient
	//using string builder is very efficient

	fmt.Println("------------------------------")
	h := "hello,"
	w := "world"

	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)

	fmt.Println(sb.String())

	fmt.Println("------------------------------")

	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("substring first 13 characters")
	fmt.Println(name[0:13])

	fmt.Println("Last 13 Characters of substring ")
	fmt.Println(name[13:26]) //slices

}
