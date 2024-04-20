package main

import "fmt"

func main() {

	// Arrays with Keyed Elements (advanced feature of Go)

	// Example 1
	grades := [3]int{ // keyed elements can in be
		2: 10,
		0: 5,
		1: 7,
	}
	fmt.Println("array grades: ", grades) // [5 7 10] check the element's order.

	// example 2
	accounts := [3]int{2: 50}                 // this array has 3 elements but specified only the one.
	fmt.Println("array accounts: ", accounts) // [0 0 50]

	// example 3
	names := [...]string{
		5: "John",
	}
	fmt.Println("names: ", names, " length: ", len(names)) // names:  [     John]  length:  6

	// what if we don't specify key of an element?
	cities := [...]string{ // an unkeeds element gets its index from the last keyed element.
		5:        "Paris",
		"London", // so London is in index 6
		1:        "New York",
	}
	fmt.Println("cities: ", cities) // cities:  [ New York    Paris London]

	// example 4
	// one advantage of array with keyed elements
	weekend := [7]bool{5:true, 6:true}
	fmt.Println("weekend: ",weekend)

}
