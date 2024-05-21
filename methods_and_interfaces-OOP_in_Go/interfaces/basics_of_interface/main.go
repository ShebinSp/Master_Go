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

func main() {
	// INTERFACES

	// An interface is a collection of method signatures that an object-
	// which is most of the time a named type, can implement.
	// Interfaces define the behaviour of an object and can achieve polymorphism.

	// Create a circle value and a rectangle value
	c1 := circle{radius: 5.}
	r1 := rectangle{width: 3., height: 2.1}
	// calling printCircle() and printRectangle() functions
	printCircle(c1)
	printRectangle(r1)

	// It is obvious that this approch is not good. We have repeated ourself and have created-
	// two functions(printCircle(), printRectangle()) with the same logic insted of one.
	// printCircle and printRectangle have the same logic, even though it's working it's not a
	// good practice to develop a program like this. So we use interfaces.

}
