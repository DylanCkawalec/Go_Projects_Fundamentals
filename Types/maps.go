package main

import "fmt"

//reference type, keys and values
//maps are random, and are not sorted
func main() {
	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5
	intMap["six"] = 6
	for key, value := range intMap {
		fmt.Println(key, value)
	}

	delete(intMap, "five")

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")

	} else {
		fmt.Println(el, " is not in map")
	}
	//overwrite values in maps
	//maps are fast,easy and no pointers needed
	intMap["two"] = 4

}
