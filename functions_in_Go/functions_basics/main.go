package main

import (
	"fmt"
)

func f1() {
	fmt.Println("This is function f1()")
}

// The function f2 takes two parameter 'a' and 'b' which is only available in inside the function f2()
// Parameter is the reserved space for arguments or variables or values passed into the function.
// If you pass 2 and 3 into the f2 function, those 2 will be stored in 'a' and 3 in 'b'.
// The space reserved to store the variables is called parameters (a & b).
// The passed variables or values to a function is called arguments(2 & 3).

func f2(a, b int) {
	fmt.Println("Sum: ", a+b)
}

// Shorthand Parameters Notation
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// Return statement
func f4(f float64) int {
	return int(f)
}

// If a function returns multiple return values, we must specify return types of the values just like above inside
// parentheses just after the function parameter parentheses and as a convention is a function returns also an
// error value and it should be the last parameter.
func f5(a, b int) (int, int) {
	return a + b, a * b
}

// Named Return Values
//named return values are a great way to explicitly mention the return variables in the function definition itself.
func sum (a, b int) (s int) {
	fmt.Println("s: ",s) // s will be 0 at this  point
	s = a + b
	return // A return statement without arguments returns the named return values. This is known as a "naked" return. 
}

func main() {
	// FUNCTIONS IN GO

	// Lets create a function outside the main function which takes no argument and no return value.
	// Go executes automatically only the main and init functions which are predefined function names.
	// calling the function inside the application entry point which main() function.
	f1()

	// calling function f2(a, b int) - a & b are the parameters
	f2(2, 3) // The passed values are called arguments (2 & 3)

	f3(1, 2, 3, 4.4, 5.5, "s")

	// There should be a variable to receive the value returned from the function
	num := f4(5.5)
	fmt.Println("num : ", num)

	// The return values can be captured in variables using multiple assignments.
	a, b := f5(3, 5)
	fmt.Printf("a: %v  ,  b: %v\n", a, b)
	
	// In case of multiple return values, if you are only interested in one single value returned by the
	// function, you can assign the other values to blank identifier which stores the value to an empty
	// variable. This is needed because if a variable is defined but not used in Go, the compiler
	// complains about it. If you are not interested in 'b', write the blank identifier.
	a, _ = f5(3, 5)
	fmt.Println("a: ", a)

	mySum := sum(4, 8)
	fmt.Println("Sum: ",mySum)

}
