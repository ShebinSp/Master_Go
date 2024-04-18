package main

import "fmt"

// Comparison operators compare two operands and yield an boolean value(true or false).

func main() {

	// == (equal) operator
	a, b := 3, 2
	c := a == b                                    // 'a' compared to 'b' using ==, and the value will be assigned to 'c' as boolean.
	fmt.Printf("%v is equal to %v: %v\n", a, b, c) // value of 'c' is false

	b++ // value of 'b' incremented by 1.

	c = a == b                                     // again 'a' compared to 'b' using ==, and the value will be assigned to 'c' as boolean.
	fmt.Printf("%v is equal to %v: %v\n", a, b, c) // value of 'c' is true
	fmt.Println()

// != (not equal) operator
	c = a != b
	fmt.Printf("%v is not equal to %v: %v\n", a, b, c) // value of 'c' is false

	b-- // value of 'b' incremented by 1.
	c = a != b
	fmt.Printf("%v is not equal to %v: %v\n", a, b, c) // value of 'c' is true
	fmt.Println()

// < (less) operator
	c = a < b // checking 'a' less than 'b'
	fmt.Printf("%v is less than %v: %v\n", a, b, c)

	b += 2 // value of 'b' is incremented by 2
	c = a < b
	fmt.Printf("%v is less than %v: %v\n", a, b, c)
	fmt.Println()

// <= (less or equal) operator
	c = a <= b // checking 'a' is less than or equl to 'b' and value assigned to c
	fmt.Printf("%v is less than or equal to %v: %v\n", a, b, c)

	a += 2        // value of 'a' is incremented by 2
	c = a <= b // checking 'a' is less than or equl to 'b'
	fmt.Printf("%v is less than or equal to %v: %v\n", a, b, c)
	fmt.Println()

// > (greater) operator
	c = a > b // checking 'a' is greater than to 'b' and value assigned to c
	fmt.Printf("%v is greater than %v: %v\n", a, b, c)

	a -= 2 // value of a decremented by 2
	c = a > b // checking 'a' is greater than to 'b' and value assigned to c
	fmt.Printf("%v is greater than %v: %v\n", a, b, c)
	fmt.Println()


	// >= (greater or equal) operator
	c = a >= b // checking 'a' is greater than or equl to 'b' and value assigned to c
	fmt.Printf("%v is greater than %v: %v\n", a, b, c)

	a += 2 // value of a incremented by 2
	c = a >= b // checking 'a' is greater than or equl to 'b' and value assigned to c
	fmt.Printf("%v is greater than %v: %v\n", a, b, c)



}
