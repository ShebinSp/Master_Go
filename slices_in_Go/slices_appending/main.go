package main

import (
	"fmt"
)

func main() {
	// how to append and copy slices

	numbers := []int{1, 2}
	fmt.Println("slice numbers before appending: ", numbers)
	// To add an element to a slice, use the append() function. It is a bultin function in Go and it returns a new slice
	// after it appends a value to its end.
	numbers = append(numbers, 3, 4)
	fmt.Println("slice numbers after appending: ", numbers)

	// to copy a slice into other slice
	n := []int{100, 200}
	fmt.Println("slice n: ", n)
	numbers = append(numbers, n...) // use ellipsis operator with slice n
	fmt.Println("slice n numbers: ", numbers)
	fmt.Println()

	// copy using copy() function
	src := []int{10, 20, 30}
	fmt.Println("slice src: ", src)
	dst := make([]int, len(src))

	nn := copy(dst, src) //copy(destination, source), nn is the number of elements copied.
	fmt.Println("slice dst: ", dst)
	fmt.Println("number of elements copied: ", nn)

	// what if destination slice is created with smaller number of elements
	//If the destination slice has zero length nothing will be copied.
	dst1 := make([]int, 2) // dst1 has the length 2, only two elements can be copied. we can append as many as we want
	no := copy(dst1, src) // copied only first 2 elements from src(10,20)
	fmt.Println("Slice dst1: ", dst1)
	fmt.Println("number of elements copied: ", no) // 2
	dst1 = append(dst1, numbers...) // appended slice numbers to dst1
	fmt.Println("ds1: ",dst1)
}
