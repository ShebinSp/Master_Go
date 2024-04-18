package main

import "fmt"

type km float64 // defined type is km and underlying type is float64
type mile float64 // defined type is mile and underlying type is float64

func main() {

	// Defined Types

	// creatin a new type
	type age int // 'age' is the new type and 'int' is it's underlying type.
	// type 'oldAge' is the new defined type and 'int' is its underlying type not age.
	// remember that, underlying types are the predeclared types.
	type oldAge age

	// Example 1:

	type speed uint

	var s1 speed = 10
	var s2 speed = 20
	// we can perform operations on the same defined types
	// type of 'as' is main.speed, means ' type `speed` is defined within package named `main` '.
	as := (s2 + s1) / 2
	fmt.Printf("\nAverage Speed: %v  | Type of `as`: %T\n\n", as, as)
	// but can't on underlying type of an defined type.
	var x uint
	// x = s1 // Error: different types
	// we can assign the value by conveting the type
	x = uint(as)
	fmt.Printf("Average Speed: %v  | Type of `x`: %T\n\n", x, x)

	// Example 2:
	
	// var s3 speed = x // Error: different types. To solve this error:
	var s3 speed = speed(x)
	fmt.Printf("Speed: %v  | Type of `s3`: %T\n\n",s3,s3)

	// Example 3:
	// To make it clear I am  defining two new type `km` & `mile` outside the main()
	var parisToLondon km = 465
	var distanceInMiles mile

	// this is an error because, even if both km and mile have underlying type float64 still these are different types.
	// One type can be converted to another type if they share the same underlying type.

	// distanceInMiles = parisToLondon / 0.621 // Error: cannot use parisToLondon / 0.621 (value of type km) as mile value in assignment
	
	// The new type will inherite all the operations of the source type, means all operations that are available to floats-
	// are also available to these newly defined types.
	distanceInMiles = mile(parisToLondon) / 0.621 // correct assignment
	fmt.Println("Distance in miles from Paris to London: ",distanceInMiles)

}
