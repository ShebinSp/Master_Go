package main

import (
	"fmt"
	"math"
)

// The first interface `shape` with a single method `area()` which returns a `float64` value
type shape interface {
	area() float64
}

// The second interface `object` with a single method `volume()` which returns a `float64` value
type object interface {
	volume() float64
}

// Declaration of an interface means specifying methods belonging to it. Besides methods, an interface can also include
// other interfaces defined in the current or in another package and import it.
// When we include an interface in another interface, we add all the methods from the embedded interface.

// `geometry` interface embeddd `shape` and `object` interfaces and has also another method called `getColor()`
// here order doesn't matter. We can write the embedded interfaces in any order.
type geometry interface {
	shape
	object
	getColor() string
}

// A new named type cube satisfy the interface. Satisfying an interface means implementing all methods of the interface.
// and cube type implements both area() and volume().
type cube struct {
	edge  float64
	color string
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}
func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}
func (c cube) getColor() string {
	return c.color
}

// Function `meassure` works on a geometry interface type. It returns the area and the volume of any geometry shape.
func meassure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func main() {
	// EMBEDDED INTERFACES

	// In Go an interface cannot implement other interfaces or extend them. Inheritance is not supported.What we can do is
	//  create a new interface by merging two or more interfaces. This is known as interface embedding.

	// Since `cube` implements the methods area() and volume(), it implements interfaces, shape and object and since the
	// interface `geometry` is embedding interfaces, shape and object, cube type must also implement getColor method
	// in order to satisfy the geometry interface. At this moment we cannot call the measure function with a cube value
	// as an argument. If do, it returns an error.
	// c := cube{edge: 3}
	// a, v := meassure(c) // Error: cube doesn't implement geometry (missing method getColor method). So lets add-
	// field color to cube type and implement getColor method. Now `cube` satisfies the `geometry` interface.
	// Means measure function can be also called on cube values.
	c := cube{edge: 3}
	a, v := meassure(c)
	fmt.Println("Area of the cube: ", a, " Volume of the cube: ", v)

	// NOTE: Circular embedding of interfaces is disallowed and will be detected at compile time.

}
