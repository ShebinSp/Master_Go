package main

import (
	"fmt"
	"math"
)

// Below there are two types rectangle and circle
type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Two methods to find the area and perimeter of circle
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Two methods to find the area and perimeter of rectangle
func (r rectangle) area() float64 {
	return r.height * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (r rectangle) method3() {
	fmt.Println("hello")
}

// Suppose we want to print the shape, area and perimeter, we should create a method with the same logic for
// each type like this,
func printCircle(c circle) {
	fmt.Println("Shape: ", c)
	fmt.Println("Area: ",c.area())
	fmt.Println("Perimeter: ",c.perimeter())
}
func printRectangle(r rectangle){
	fmt.Println("Shape: ", r)
	fmt.Println("Area: ",r.area())
	fmt.Println("Perimeter: ",r.perimeter())
}

//  we declare an interface like a struct or any other named type,
// using the type keyword, a name, and then the interface type.
// The interface contains only the signature of the methods, but not their implimentation.
// The signature of a method means its name, input, parameters and return values.

// Any type that impliments `area()` and `perimeter()` methods is said to be implenting the `shape` interface.
// Here, `rectangle` and `circle` types implement the interface, because both of them 
// implement `area()` and `perimeter()` methods.

type shape interface{
	// what behaviour of the rectangle and circles have in common, which is area and perimeter.
	area() float64
	perimeter() float64
}

// Now instead of `printRectangle()` and `printCircle()` functions, I want a single function that
// prints out information about any rectangle, circle so in fact, any shape, any type that implements
// the `shape` interface or implement methods `area()` and `perimeter()`.

// This function takes in a `shape` type, so in fact it can take in any type that implements the
// interface. If a variable has an interface type, then we can call methods that are in the
// named interface. This function works on any type that implements the `shape` interface.
// Any type that implements the interface is also of type of the interface. This is very important.
// Rectangles and circles are also of type `shape`. Also note that a type can implement one or-
// more interfaces.
// In this example, we'll call the function using a rectangle or a circle value because they 
// implement the interface.
// `s` is an abstract value, but I'll call the function using a rectangle or a circle value which-
// are concrete values.
func print(s shape){
	// I'll call the methods that define the behaviour or the interface.
	fmt.Printf("Shape: %#v\n",s)
	fmt.Printf("Area: %#v\n",s.area())
	fmt.Printf("Perimeter: %#v\n",s.perimeter())
}

func main() {
	// INTERFACES

	// An interface is a collection of method signatures that an object-
	// which is most of the time a named type, can implement.
	// Interfaces define the behaviour of an object and can achieve polymorphism.

	// Create a circle value and a rectangle value
	c1 := circle{radius: 5.}
	r1 := rectangle{width: 3., height: 2.1}
	// calling printCircle() and printRectangle() functions
	//printCircle(c1)
	//printRectangle(r1)

	// It is obvious that this approch is not good. We have repeated ourself and have created-
	// two functions(printCircle(), printRectangle()) with the same logic insted of one.
	// printCircle and printRectangle have the same logic, even though it's working it's not a
	// good practice to develop a program like this. So we use interfaces, which is like a contract
	// that defines behaviour that must be implimented for similar type of objects.
	// Interfaces are just a tool. Whether you can use them or not is upto you, but they can make
	// code clearer, shorter, more redable, and they can provide a nice API between
	//  packages or clients and servers.

	// After implementing the interface `shape`, instead of calling `printCricle()` and `printRectangle()`
	// on line 90 and 91, so two different function, I'll only call the `print()` function.
	// If the `print()` function was declared using an `s` parameter, we can pass in any value that
	// implements that implements that interface. Here the `shape` interface.
	print(c1)
	print(r1)

	// We also see here polymorphism in action. The types that implement the interface must implement all
	// methods or we get an error. They can freely implement also other methods for example `method3()` is
	// only for rectangle type on line 31.

	// If you comment out the `area()` method of `circle` type, you will get an error, `circle` doesn't
	// implement the `shape` interface, because of missing the interface signature `area()` which must
	// be implemented by type `circle`.

}
