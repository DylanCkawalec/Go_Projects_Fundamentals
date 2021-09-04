package main

import "fmt"

func main() {
	apples := 18
	oranges := 9

	//bool expression

	fmt.Println(apples == oranges) //condition will equate to false
	fmt.Println(apples != oranges) //evaluates to true

	// > < >= <=
	fmt.Printf("%d > %d: %t", apples, oranges, apples > oranges)
	fmt.Println()
	fmt.Printf("%d < %d: %t", apples, oranges, apples < oranges)
	fmt.Println()
	fmt.Printf("%d >= %d: %t", apples, oranges, apples >= oranges)
	fmt.Println()
	fmt.Printf("%d <= %d: %t", apples, oranges, apples <= oranges)
	fmt.Println()

}
