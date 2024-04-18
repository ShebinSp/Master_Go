package main

import "fmt"

func main() {
	// --------> Arithmetic Operators <--------
	// Arithmetic operators apply to numeric values and are used to perform common mathematical operations.
	// There is no operator for power in Go. Use function math.Pow() to calculate power.

	// Addition
	a, b := 2, 1
	res := a + b
	fmt.Printf("Result of %v + %v : %v\n", a, b, res)

	// Subtraction
	res = a - b
	fmt.Printf("Result of %v - %v : %v\n", a, b, res)

	// Multiplication
	res = a * b
	fmt.Printf("Result of %v * %v : %v\n", a, b, res)

	// Divition
	res = a / b
	fmt.Printf("Result of %v / %v : %v\n", a, b, res)

	// Modulus
	res = a % b
	fmt.Printf("Result of %v %q %v : %v\n", a, "%", b, res)

								// OR
	// I have used parentheses to change the default priority of arithmetic operators. Without
	// parentheses multiplication and division have a higher priority than addition and subtracton.
	// That's because in mathematics & computer programming the order of operations or
	// operator precedence is a collection of rules that reflect the conventions about which
	// operations will be first ecaluated.
	res = (a + b) / (a - b) * 2
	fmt.Printf("Result of (%v + %v) / (%v - %v) * %v : %v\n", a, b, a, b, 2, res)

}
