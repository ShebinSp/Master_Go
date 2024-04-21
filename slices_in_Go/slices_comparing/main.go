package main

import "fmt"

func main() {
	// Comparing Slices

	var n []int // declared an empty slice
	fmt.Println("is slice n nil?: ", n == nil)

	m := []int{} // initialized a slice
	fmt.Println("is slice m nil?: ", m == nil)

	// slices can only be compared to nil
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	// fmt.Println("a == b", a == b) // it is a compile time error
	// so to compare two slices, we can use a for loop to iterate over elements to compare
	var eq = true

	if len(a) == len(b) {  // remember to perform length comparison, its important when comparing slices
		for i, valueA := range a {
			if valueA != b[i] {
				eq = false
				break
			}
		}
	} else {
		eq = false
	}

	if eq {
		fmt.Println("Slice 'a' is equal to slice 'b'")
	} else {
		fmt.Println("Slice 'a' is not equal to Slice 'b'")

	}
}
