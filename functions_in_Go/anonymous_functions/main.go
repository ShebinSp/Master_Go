package main

import "fmt"

// The function takes in an int and returns an anonymous function that returns an int.
// The function that will returned will be defined inline.
func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	// ANNONYMOUS FUNCTION

	// Being an anonymous function, so function without a name cannot invoke it later because there is no name to call, so
	// invoke it now to invoke the function.
	func(msg string) {
		fmt.Println("Msg: ", msg)
	}("I am an annonymous function!") // To invoke use parantheses and an argument which is a string value.

	a := increment(10)
	fmt.Printf("Type of a: %T\n",a)
	a() // 11
	fmt.Println(a()) // 12
	a() // 13
	a() // 14
	fmt.Println(a()) // 15
}
