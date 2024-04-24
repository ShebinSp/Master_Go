package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	//A slice is formed by specifying a start or a low bound and a stop or high bound like-
	// a[start : stop] This selects a range of elements that includes the element at index start, but excludes the element at index stop.
	b := a[1:3] // it will return elements 2 and 3 which are at index 1 and index 2. The 3rd index is excluded
	// Slicing an array returns a slice, not an array. b is a slice
	fmt.Printf("Elements in b: %v  |  type of b: %T\n", b, b)

	// Example
	s1 := []int{1,2,3,4,5,6,7}
	s2 := s1[1:4] // from s1 3 elements included and 4 elements excluded
	fmt.Printf("Elements in s2: %v  |  type of s2: %T\n", s2, s2)

	// a missing low or start index defaults to 0 and a missing high or stop index defaults to the length of the slice operand
	fmt.Println("s1: from index 2: ", s1[2:]) // missing high index. Same as s1[2: len(s1)]
	fmt.Println("s1: upto index 2(index 3 is excluded): ", s1[:3]) // missing low index . Same as s1[0:3]
	fmt.Println("s1: ", s1[:]) // missing low and high index . Same as s1[0:len(s1)] -> the entire slice

	// Using append() function with slicing
	s1 = append(s1[:4], 100) // append(s1[index 0 to index 3], 100)
	fmt.Println("slice s1: ",s1) // slice s1:  [1 2 3 4 100]

	s1 = append(s1[:4],200) // the element at index 4 will be replaced with 200
	fmt.Println("slice s1: ",s1) // slice s1:  [1 2 3 4 200]



}
