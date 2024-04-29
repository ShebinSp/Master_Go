package main

import "fmt"

func main() {
	// HOW TO SLICE A STRINGS

	s1 := "I love Go"
	// the slice expression returns a new string that contains the bytes from index 2 to index 6, 6 being excluded.
	fmt.Println("Slicing s1: ", s1[2:6])

	// A slice that contains non ASCII characters
	s2 := "കമ്പ്യൂട്ടർ"
	fmt.Println("Slicing non ASCII s2: ",s2[0:1])// unicode representation of ക

	// How to return a rune, not a slice of bytes
	// convert a slice of bytes to slice of runes
	rs := []rune(s2)
	fmt.Printf("Type of rs: %T\n",rs) // []int32
	// slice the rune slice and convert it into string
	fmt.Println("s2: ",string(rs[0:1]))
}
