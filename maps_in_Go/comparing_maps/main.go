package main

import "fmt"

func main() {
	// COMPARING MAPS

	// A map can be compared only to nil.
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}
	// fmt.Println("a equal to b?",a == b) // invalid operation: a == b (map can only be compared to nil)

	// To compare maps in Go, fmt.Sprintf() can be used only if both key-value pairs are string.
	s1 := fmt.Sprintf("%s",a)
	s2 := fmt.Sprintf("%s",b)
	fmt.Println("s1 & s2: ", s1,s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
