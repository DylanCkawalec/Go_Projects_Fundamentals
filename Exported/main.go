package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 80000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 10000, FullTime: false},
	{FirstName: "Rob", LastName: "Packard", Salary: 35000, FullTime: false},
	{FirstName: "Sahid", LastName: "Bromada", Salary: 60000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}
	//myStaff.All()

	//log.Println(myStaff.All())
	staff.OverPaidLimit = 10000
	log.Println("Overpaid Staff; ", myStaff.Overpaid())
	log.Println("Underpaid Staff; ", myStaff.Underpaid())

	//myStaff.notvisible()
}
