package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "rodent")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Eleement 0 is", animals[0])

	fmt.Println("First two elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "Elements long")

	fmt.Println("Is it sorted ? ", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Is it sorted now ? ", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)

	//deleting an string and the index and returning the string
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] //copy the last element to index i
	a[len(a)-1] = ""   //erase the last element
	a = a[:len(a)-1]   //delete the last element
	return a
}
