package main

import (
	"fmt"
)

func main() {
	// MAPS in GO

	// declaring a map
	// example 1
	var employees map[string]string // map employees decalred (nil)
	fmt.Printf("map employees: %#v\n", employees)
	fmt.Printf("No of pairs: %d\n", len(employees))

	// we can get the corresponding value of a key even if the map is not initialized
	// If an element doesn't exist in a map, it returns a zero value of that type.
	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"]) // ""

	// example 2
	var accounts map[string]float64
	fmt.Printf("map accounts: %#v\n", accounts["0x231"])

	// When declaring a map, we can use only comparable types as keys. ie, we can't use a slice-
	// because a slice is not comparable.
	// var m1 map[[]int]string // invalid key type

	// we can declare arrays as key, since arrays are comparable.
	var m2 map[[3]int]string // valid
	_ = m2

	// ADD elements to the map employee

	// It is illegal to populate an initialized map in Go

	// employees["Doe"] = "Programmer" // nil dereference in map update


	// initialize the map using the make() function or a map literal before add any element.
	// INITIALIZE a map - short declaration
	people := map[string]float64{} // map is initialized but empty and not nil with length 0
	// maps can only compare to nil
	fmt.Printf("Is map people nil?: %v | Length : %v\n", people == nil, len(people))

	// lets add some elements to the people
	people["John"] = 21.1
	people["Marry"] = 23
	people["Doe"] = 11.11
	fmt.Println("people: ", people)

	// INITIALIZE using make() function
	map1 := make(map[int]int) // map map1 is declared and initialized but empty
	fmt.Printf("Is map map1 nil?: %v  |  Length : %v\n", map1 == nil, len(map1))
	map1[3] = 5

	// INITIALIZE when declaring with key-value pairs
	balances := map[string]float64{
		"USD": 303.1,
		"EUR": 430.3,
		//	55 : 312.1 // Error: all keys must be the same type
	}
	fmt.Println("Map balances: ", balances)

	// IF THE KEY EXISTS ITS UPDATES IT'S VALUE but IF THE KEY DOES'T EXIST, IT ADDS THE KEY VALUE PAIR.
	balances["USD"] = 500 // Updates the value 303.1
	balances["GBP"] = 100 // Adds a new key GBP with corresponding value 100
	fmt.Println("Map balances: ", balances)

	// Sometimes you need to distinguish a missing entry from a zero value-
	// to do that the coma ok idiom is used.

	// Missing entry
	fmt.Println("RON: ",balances["RON"]) // key RON is not in the map the defaul zero value of float result
	// So, to check the missing entry, use the coma ok idiom syntax
	v, ok := balances["RON"] // ok if the RON key exists, false otherwise
	if ok {
		fmt.Println("The RON balance is: ",v)
	} else {
		fmt.Println("The RON key does not exist in the map")
	}

	// We can iterate over a map but not indicated, because:
	// 1. The order of key-value pair changes.
	// 2. The maps in Go have been designed for a fast lookup time, not for fast looping.

	for k, v := range balances {
		fmt.Printf("key : %#v  |  value : %#v\n",k,v)
	}

	// To delete a key-value pair from the map, use the built-in function delete()
	delete(balances, "USD")
	fmt.Println("balances after deletion: ",balances)

}
