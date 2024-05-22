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
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type shape interface {
	area() float64
	perimeter() float64
}

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %#v\n", s.area())
	fmt.Printf("Perimeter: %#v\n", s.perimeter())
}
func main() {
	// INTERFACE TYPE ASSERTION
	// A type asserion provides access to an interface concrete value.

	var s shape = circle{radius: 2.5} // declared an interface value that holds a circle type value. The circle type-
	// implements the interface.

	// Adding a new method to the circle type, volume of a sphere with the circle's radius.
	// As we know, an interface has an underlying dynamic type and a concrete value.
	fmt.Printf("Dinamic type of interface s: %T\n", s) // circle
	// Even though the type is circle, call a method of the underlying type on the interface value, an error returns
	// s.volume() // its an error
	// Why? even s is of type circle and circle has volume() method defined..
	// Because, even though the dynamic type os 's' is circle, we can access only the methods that are defined-
	// inside the interface `shape` which is 'area()' and 'perimeter()'.

	// An interface value hikes its dynamic value. Actually, this is one of the reasons we use interfaces. They-
	// isolate and decouple one part of the code from another.

	// How can we access the dynamic value in order to be able to invoke a method of the underlying dynamic type-
	// on the interface value?
	// To do that, we need to use something called type assertion. It will extract and return the dynamic value-
	// of the interface value.
	// We use the interface value 's' and the type of the dynamic value we want to extract between-
	// parentheses. This is how to do type assertion. Now calling volume() method now is not an error.
	boll := s.(circle).volume() // not a good practice.
	fmt.Println("Ball Volume: ", boll)

	// But, sometimes the type assertion can fail, and it's a good practice to check whether the asserton succeeded-
	// or not. They type assertion can return two values. The underlying value and a boolean value that reports
	// whether the assertion succeeded.
	// So lets code
	ball, ok := s.(circle) // ok equals 'true' if the assertion was successful and 'false' otherwise.
	// If the assertion was successful, boll holds the dynamic value and if not, it holds the zero value for that type.
	// This is similar to the way we read from a map value.
	if ok {
		fmt.Printf("Ball Volume: %v\n", ball.volume())
	}

	//** TYPE SWITCHES **//
	// A type switch is a construct that permits several type assertions in series and runs the first case with-
	// a matching type. The declaration in a type switch has the same syntax as a type assertion, but the specific-
	// type is replaced with the keyword `type`.
	
	// s = rectangle{width: 2, height: 3}
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n",value) // executed if the type is circle.
	case rectangle:
		fmt.Printf("%#v has rectangle type\n",value) // executed if the type is rectangle.
	}
	// Type switch is like a regular switch statement, but the cases in a type switch specify types, not values, and
	// those are compared against the type of the value held by the given interface value.
	// If you change the concrete value of 's' to rectangle, then s.type will equals to the rectangle in cases.

}
