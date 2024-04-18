package main

import "fmt"

func main() {
	a, b := 2, 3

	// increment assignment
	a += b // value of 'b' added to value of a and assigned to 'a'
	fmt.Println("value of a (increment assignment): ", a)

	// decrement assignment
	b -= 2 // from the value of a 2 will be decrement and assigned to 'a'
	fmt.Println("value of b (decrement assignment): ",b)

	// multiplication assignment
	b *= 10
	fmt.Println("value of b (multiplication assignment): ",b)

	// division assignment
	b /= 5
	fmt.Println("value of b (division assignment): ",b)

	// modulus assignment
	a %= 3 // computes modulus of 'a', which is 5 and 3 and assign the result back to 'a'
	fmt.Println("value of a (modulus assignment): ",a)

	// Increment and Decrement statement
	// The increment and decrement statements increment or decrement their operands by 1.
	//According Go specification increment and decrement are statements,not operators-
	// like in C, C++ or Java.

	// increment statement
	x := 0
	x += 2 // added 2 using increment assignment
	x++ // 1 will added to x using increment statement
	fmt.Println("value of x (increment desk statement): ",x) // value of 'x' is 3

	// decrement statement
	y := 2
	y--
	fmt.Println("value ofyx (decrement desk statement): ",y) // value of 'y' is 1


}
 