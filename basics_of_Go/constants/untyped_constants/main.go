package main

import "fmt"

func main() {
	const a float64 = 5.1 // typed constant

	const b = 6.7 // untyped constant

	// we can initialize constant with expressions
	const c float64 = a * b
	const str = "Hello " + "Go!"
	const boo = 5 > 10

	const x = 5
	const y = 2.2 * x
	fmt.Printf("Value of \"x\": %v, and the type of \"x\"is %T\n", x, x)
	fmt.Printf("Value of \"y\": %v, and the type of \"y\"is %T\n", y, y)
	// an untyped constant makes it possible for us to use constants in Go with great freedom
	// we can cross the different type untyped(typeless) constants without changing its type.

	var i int = x // x changes into int
	var j float64 = x // behind the scene:-> var j float 64 = float64(x)
	_, _ = i,j  // muted compile time error

	// When an untyped constant is used in a context that requires a type, a type will be inferred
	// and untyped constant has a default type, an implicit type that it transfers to a value
	// if a type is needed but none is provided. An untyped constant is converted to default 
	// type only when only when needed in an expression.

	const r = 7
	var rr float32 = r
	sr := r+rr
	fmt.Printf("Type of r: %T & Type of rr: %T\n",r,rr)
	fmt.Printf("Sum of r + rr : %v & Type of sr is: %T\n",sr,sr)

	// Another way to think about untyped constants is that they exist in a kind of ideal space
	// of values which is less restrictive than Go's full type system. To do anything with them
	// we need to assign them to variables and when that happens the variable, not the constant
	// itself, needs a type and the constant can tell the variable what type it should have.
}
