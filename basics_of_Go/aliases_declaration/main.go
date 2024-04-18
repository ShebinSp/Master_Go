package main

import "fmt"

func main() {
	// alises

	var a uint8 = 1
	var b byte = 3
	// the type of var 'a' and 'b'
	fmt.Printf("a: %T    |     b: %T\n",a,b) // both are type uint8 or byte is an alias for uint8

	// Because a, b are aliases, we can perform operations without converting the types.
	a = b
	fmt.Printf("Type of 'a': %T    |    Value of 'a': %v\n", a, a)

	// Declaring a new aliase named second for uint
	type second = uint // declared using the equal sign instead of a whitespace after the new name
	// no need to convert operations (same type). If i have to define a declared type insted of alias,
	// I have to convert the value to performarithmetic operation.
	var hour second = 3600
	fmt.Printf("Minutes in an hour: %d \n",hour/60)
}
