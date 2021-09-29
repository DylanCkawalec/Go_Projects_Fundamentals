package main

import "fmt"

func main() {
	//starting at Zero
	courseName := "Learn Go For beginners Crash Course"
	//must cast using string method
	fmt.Println(string(courseName[0]))
	fmt.Println(string(courseName[6]))

	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}

	fmt.Println()
	for i := 13; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}

}
