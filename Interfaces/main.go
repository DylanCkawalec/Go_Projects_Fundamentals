package main

//run these commands newbie
//go mod init myapp
//go run main.go

import "fmt"

//generate a riddle
func riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says %s and has %d legs. what animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}

//interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

//dog is the type for dogs

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

//Cat is the type for cats
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	hasTail      bool
}

//these are the interfaces for dogs and cats
//you need them to have the same values of their respective interface "Animal"
func (d *Dog) Says() string {
	return d.Sound
}
func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (d *Cat) Says() string {
	return d.Sound
}
func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}

func main() {

	//Make a dog like this
	dog := Dog{
		Name:         "Dog",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}

	riddle(&dog)
	//another way to make a variable
	var cat Cat
	cat.Name = "Cat"
	cat.NumberOfLegs = 4
	cat.Sound = "Meow"
	cat.hasTail = true

	riddle(&cat) //the & will interface these to point at their respective interface

}
