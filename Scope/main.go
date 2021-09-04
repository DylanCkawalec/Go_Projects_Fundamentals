package main

import (
	"myapp/packageone"
)

//decalre a package level variable for the main
//package name myVar
var myVar = "this is a package level variable"

func main() {
	//variables
	//declare a block variable for the main function
	//called blockVar
	var blockVar string = " This is the block level variable "
	secondVar := "some Value"
	
	//create an exported function in packageon called PrintMe

	packageone.PrintMe(myVar, blockVar)
	//in the main function, print out the values of myVar.
	//blockVar, and PackageVar on one line
	//function in packageon

}
