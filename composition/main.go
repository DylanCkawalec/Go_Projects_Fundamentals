package main

import (
	"fmt"
)

//types
type Vehical struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehical    Vehical //calls composition, has access too the fields in vehical
}

//tied to vehical
func (v Vehical) showDetails() {
	fmt.Println("Number of passangers: ", v.NumberOfPassengers)
	fmt.Println("Number of wheels:", v.NumberOfWheels)
}

//tied to car
func (c Car) show() {
	fmt.Println("Make: ", c.Make)
	fmt.Println("Model: ", c.Model)
	fmt.Println("year: ", c.Year)
	fmt.Println("Electric: ", c.IsElectric)
	fmt.Println("Hybrid: ", c.IsHybrid)
	c.Vehical.showDetails()

}

func main() {
	//vehical types
	suv := Vehical{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}
	Motorcycle := Vehical{
		NumberOfWheels:     2,
		NumberOfPassengers: 1,
	}
	//car types
	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   true,

		Vehical: suv,
	}

	teslaModelX := Car{
		Make:       "tesla",
		Model:      "Model X",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehical:    suv,
	}

	Kawasaki := Car{
		Make:       "Susuki",
		Model:      "787",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehical:    Motorcycle,
	}

	volvoXC90.show()
	fmt.Println()
	//make another one
	teslaModelX.Vehical.NumberOfPassengers = 7
	Kawasaki.show()
	fmt.Println()
	teslaModelX.show()
}
