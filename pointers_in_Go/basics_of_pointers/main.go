package main

import "fmt"

func main() {
	// POINTERS IN GO

	// To access memory address of a variable use '&' before the variable
	name := "Linto"
	fmt.Println("Memory address of variable 'name' is: ", &name)

	// How to declare and initialize pointers
	// The expression `&` means the address of `x` and creates a pointer to an integer variable.
	// The new variable `ptr` is of type `*int` which is pronounced pointer to int.
	var x int = 2
	ptr := &x // address of x is assigned to ptr
	fmt.Printf("variable ptr is of type '%T' with value of '%v'\n", ptr, ptr)

	// `ptr` holds the address of `x`
	fmt.Printf("Address of `x`: %p\n", &x)
	fmt.Printf("Value in `ptr`: %v\n", ptr) // address of x

	// A pointer itself also a variable stored somewhere in memory. If you want to print the memory address of a
	// variable use %v and & variable as the argument.
	fmt.Printf("`ptr` is of type '%T' with a value of '%v' and address of `ptr` is '%p'\n", ptr, ptr, &ptr)

	// if you want to declare a pointer without initializing
	var ptr1 *float64 // not initialized and the value is 'nil'
	_ = ptr1
	// another way is to use the built in function `new()`
	// The new() function takes a type as argument and allocates enough memory to accommodate a value of that type-
	// and returns a pointer to it.
	p := new(int) // created a pointer called `p` to an int type
	x = 100
	p = &x
	fmt.Printf("`p` is of type '%T' with a value of '%v'\n", p, p)
	fmt.Printf("Address of `x`: %p\n", &x) // same of the value of p & ptr

	// Dereferencing Operator
	// To find the value of pointer points to use the * operator which is also called the dereferencing operator.
	// It is placed before a pointer variable and returns the value in that memory address.

	// above `p` points to `x` or equivalently `p` contains the address of `x`
	// The variable to which `p` is written `*p` so (*p == x)
	// The * operator denotes the pointers underlying value.
	*p = 90 // equivalent to x = 90
	fmt.Println("x and p: ", x, *p)
	
	// To make it clear
	fmt.Println("*p == x: ", *p == x) // true

	// The * Operator in `var ptr1 *float64 and `*p`
	// * in front of a type, as in *float64 `var n *float64` means type description-
	// we are declaring a pointer to that type, in this case float64.
	// When we see * in front of a pointer, it means the dereferencing operator or-
	// that we want the value that the pointer is pointing to.`

	*p = 10 // x = 10 because p is address of x and *p is value of x.
	// we can divide x through the pointer
	*p = *p / 2 // x / 2
	fmt.Println("x after *p / 2: ", x) // x changed by pointer

	// Remember
	// &value => pointer
	// *pointer => value

}
