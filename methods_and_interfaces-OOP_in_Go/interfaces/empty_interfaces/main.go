package main

import "fmt"

// empty interface
type emptyInterface interface{}

type person struct{
	info interface{}
}

func main() {
	// EMPTY INTERFACE'

	// An interface which is empty or no method called empty interface. It;s completly empty.
	// Any type satisfies this interface, and that means it can represent any value.
	// The empty interface is a key concept in Go.
	// In Go, empty interface is everywhere, for example the Println() function,
	// func Println(a ...interface{})(n int, err error) // This is a variadic parameter and it means that the function
	// can receiver any number of empty interface values or any number of arguments of any type.

	// Empty interface can hold any value
	// The interface is a type, so create a value of type `emptyInterface`
	var empty emptyInterface

	// store any value in `empty`
	// int
	empty = 5
	fmt.Println("Int in empty: ", empty)

	// string
	empty = "Go"
	fmt.Println("String in empty: ", empty)

	// slice
	empty = []int{1, 2, 3}
	fmt.Println("Slice in empty: ", empty)

	// How to see the length of the empty?
	// We cannot use an empty interface in operations.
	// fmt.Println("Length of empty",len(empty)) // its an error
	// We cannot use an empty interface in operations. An interface stores in fact two values;
	// a dynamic type and a dynamic concrete value. To access the dynamic value, we have to do an assertion.
	// If we want to use a function that accepts a slice as argument  or a method that works on a slice value,
	// we must retrieve the dynamic value using an assertion.
	fmt.Println("Slice in empty: ", len(empty.([]int))) // This is how we get the dynamic value stored in the interface.

	// Empty interfaces are used by code that handles values of unknown type, and you can pass an-
	// empty interface type as a function parameter of any type.
	// Lets create a struct type `person` that has one field called `info` which is of type `empty interface`
	p1 := person{}
	// store a string value
	p1.info = "John Doe"
	fmt.Println("Person 1: ", p1.info)
	// store an int value
	p1.info = 30
	fmt.Println("Person 1: ", p1.info)
	// store a slice
	p1.info = []float64{3.,4.4,5.,6.,7.777}
	fmt.Println("Person 1: ", p1.info)
	// length of the p1.info
	concreteValue := p1.info.([]float64) // extracting concrete value from p1.info empty interface
	fmt.Println("Length of the p1.info:",len(concreteValue))

	// This is the power of empty interface, and everything that is so powerful comes with a cost.
	// Empty interfaces can cause programs to become hard to maintain.
	// The empty interface type is increasingly being misused as a convenient way to bypass go compilers
	// type checking and one of the principle of Go it that us allows us to write type safe code.
	// There are places where empty interface is used in place where explicit interface could have
	// been used. The problem with runtime type checking is that you will never know if there is a
	// problem until it is run.
	// Use empty interfaces only if it's really necessary.

}
