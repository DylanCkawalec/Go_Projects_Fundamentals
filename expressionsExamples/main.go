package main

import "fmt"

func main() {
	age := 10           //integer literall of 10
	name := "jack"      //string literal
	rightHanded := true //bool literal

	fmt.Printf("%s is %d years old. Right Handed: %t", name, age, rightHanded) //expressions
	fmt.Println()

	ageInTenYears := age + 10 //expression

	fmt.Printf("In ten years , %s will be %d years old", name, ageInTenYears) //expressions
	fmt.Println()
	isATeenager := age >= 13                         //expression
	fmt.Printf(name, "is a teenager: ", isATeenager) //expressions

	//expressions are just evaluated to a single value
}
