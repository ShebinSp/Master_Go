package main

import (
	"fmt"
	"log"
)

func devide(a, b int) int {
	if b == 0 {
		log.Fatalf("\nError: Division by \"%d\"",b)	 
	}
	return a/b
}

func main() {
	// Divition a number by zero is an error(its an error in Go also).
	// Lets code a custom error which happens when a number divided by zero.
	sum := devide(5, 0)
	fmt.Println("Sum :",sum)
}
