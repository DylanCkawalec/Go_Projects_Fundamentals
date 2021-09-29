package main

import "fmt"

func main() {

	courseName := "Learn Go For Beginners Crash Course"
	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}

	fmt.Println()

	fmt.Println("Lenghth of courseName is", len(courseName))

	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")

	fmt.Println("mySlise has", len(mySlice), " Elements")
	fmt.Println("the last element in my slice is", mySlice[len(mySlice)-1])

}
