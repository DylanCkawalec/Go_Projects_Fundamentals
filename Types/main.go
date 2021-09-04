package main

import "fmt"

// aggregate types (array, struct)

//reference types( pointers, slices, maps, functions, channels)

//interface type
type Car struct {
	Tire       int
	Lux        bool
	Bucketseat bool
	Make       string
	Model      string
	Year       int
}

func main() {
	var myCar = Car{
		Tire:       4,
		Lux:        true,
		Bucketseat: true,
		Make:       "Volvo",
		Model:      "XC90",
		Year:       2019,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
}
