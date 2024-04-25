package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// To Create a complete new slice from an existing slice or a copy of an existing slice-
	// we should use the append() function.

	// Example 1
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...) // the slice newCars is a brand new one. It doesn't-
	// shares the backing array of the cars.
	cars[0] = "BMW" // element modified in the `cars` but not in the `newCars`.
	fmt.Println("slice cars: ", cars, "  | slice new cars: ", newCars)

	// Example 2
	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3] // the length(3) is the number of elements in the slice-
	// the capacity(5) is the number of the elements in its backing array
	fmt.Println("length and capacity of newSlice: ", len(newSlice), cap(newSlice))

	// Example 3
	// In the slice header the address of the backing array, which is a pointer-
	// is the address of the first element of its backing array.
	newSlice = s1[2:5] // {30,40,50}
	fmt.Println("length and capacity of newSlice: ", len(newSlice), cap(newSlice))

	// Slicing Operations are cheaper than array operations
	fmt.Printf("Address of s1: %p\n", &s1) // print the address of the slice header

	fmt.Printf("Address of s1 & newSlice: %p  %p\n", &s1, &newSlice) // address may be different but they share same underlying array
	newSlice[0] = 1000                                               // lets modify newSlice to prove it

	fmt.Println("s1: ", s1) // value in index 2 changed in s1.

	// How to check the memory size of any variable including arrays and slices
	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5, 6}
	
	// The slice occupies less memory than an array.
	fmt.Printf("array's size in bytes: %d\n", unsafe.Sizeof(a)) // 8 bytes per item
	fmt.Printf("slice's size in bytes: %d\n", unsafe.Sizeof(s)) // 3 items in the slice header * 8
	
}
