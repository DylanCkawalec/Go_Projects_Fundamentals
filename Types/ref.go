package main

import "fmt"

func main() {
	x := 10
	//x is now 10, and we will point now.
	myFirstPointer := &x

	fmt.Println("x is : ", x) // this is only pointing to x where it is in memory
	fmt.Println("My first pointer is", myFirstPointer)
	//when you point to something in memory, it will change the value of the original address.
	//because the pointer is allocating 15 to x, x then is 15,

	//even though we didn't change the value of x
	*myFirstPointer = 15

	fmt.Println("x is now: ", x)
	changeValueOfPointer(&x)

	fmt.Println("after function call is now: ", x)

}

//value named num with a pointer that is a integer
func changeValueOfPointer(num *int) {
	*num = 25
}
