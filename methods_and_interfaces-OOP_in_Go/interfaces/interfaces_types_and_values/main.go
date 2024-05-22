package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type shape interface{
	area() float64
	perimeter() float64
}

func print(s shape){
	fmt.Printf("Shape: %#v\n",s)
	fmt.Printf("Area: %#v\n",s.area())
	fmt.Printf("Perimeter: %#v\n",s.perimeter())
}

func main() {
	// INTERFACES TYPES AND VALUES
	// The zero value of an interface type is `nil`. The `nil` neither value nor concrete type,
	//  and calling a method on a `nil` interface is a runtime error. However,
	//it is common to write methods that successfully handle being called with a nil receiver.

	var s shape
	fmt.Printf("value of s: %T\n",s)

	ball := circle{radius: 2.5} // a concrete value of type circle
	
	// A variable of interface type can hold any type that implements that interface
	s = ball // under the hood, interface values can be thought of as a pair of a concrete value and a dynamic type.
	// An interface value holds a value of a specific underlying concrete type.

	// Calling a method on an interface value or a function that takes in an interface value executes the method-
	// of the same name on its underlying type.
	print(s) // it executes the print() function by passing in a circle value because this is the underlying concrete-
	// type of s.

	// what is the concrete type of s now?
	fmt.Printf("value of s: %T\n",s) // circle. This is, by the way, what polymorphism means. As can take 'poly' or many-
	// dynamic forms at runtime or the dynamic type of an interface value may vary during execution.

	// Example
	// Creating a concrete value of type rectangle
	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("value of s: %T\n",s) // rectangle, which is the concrete type of the value assigned to the variable at
	// runtime. This is always true unless the value is equal to nil, and in this case it has no type. This is somehow-
	// similar to what we have seen in slices.
	// A slice holds a reference to a backing array and we an say that an interface also works in a similar way by-
	// dynamically holding a reference to the underlying type.
	// In conclusion we say that in Go variables have types known at compilation and never change, but interfaces-
	// have dynamic types that change during runtime.




}
