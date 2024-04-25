package main

import "fmt"

func main() {
	// Slices Backing Array (how a slice is represented in memeory)

	// When a slice is created by slicing or assigning a slice, then both slices shares the same backing array
	s1 := []int{10, 20, 30, 40, 50} // a slice called s1
	s3, s4 := s1[0:2], s1[1:3]      // slice s3 & s4 created from s1

	s3[1] = 600             // element at index 1 of slice s3 replaced by 600
	fmt.Println("s1: ", s1) // you can  see the element changed in all the slices s1,s3,s4
	fmt.Println("s4: ", s4) // because all the 3 slices shares the same backing array.

	// Example 2
	a := []int{1, 2, 3}                              // created a slice
	b := a                                           // slice b is created by variable assignment
	b[1] = 10                                        // modification of the slice `b` will alse affect the elements of slice `a`
	fmt.Println("slice a: ", a, "  |  slice b: ", b) // as they both point to the same underlying array.
	fmt.Println()

	// When a slice is created by slicing an array, the array becomes the backing array of the slice.
	arr := [4]int{10, 20, 30, 40}        // an array is initialised
	slice1, slice2 := arr[0:2], arr[1:3] // created 2 slices by slicing the array

	arr[1] = 2                                                                  // element in index 1 modified in the array
	fmt.Println("arr: ", arr, "  |  slice1: ", slice1, "  |  slice2: ", slice2) // the modification will affected in slices also
}
