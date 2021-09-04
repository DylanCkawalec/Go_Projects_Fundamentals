package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

//reciever
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println() //carriage return
}

func (a *Animal) howManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

//variadict function
func main() {
	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.Says()

	cat := Animal{
		Name:         "cat",
		Sound:        "Meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.howManyLegs()
}

func sumMany(nums ...int) int {
	total := 0 //initialize the total value to a set 0

	for _, x := range nums { //without an iteration key, just go over all the values of what's passing into nums
		total = total + x //assign total + x
	}
	return total
}
