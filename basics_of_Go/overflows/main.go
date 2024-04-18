package main

import (
	"fmt"
	"math"
)

func main() {

// Overflow
	var x uint8 = 255
	x++ // overflow, value of x set to 0
	// Go resets the value to minimum value which is 0 for uint.
	fmt.Println("Value of x: ", x)
	
	// Compile time detected overflow
	/*var y uint8 = 256 // ERROR: range of uint is 0 - 255
	fmt.Println("Value of y: ",y)*/
	
	// Overflow happend on runtime
	var z int8 = 127
	fmt.Println("Value of z: ", z + 1) // z will be set to -127 on runtime

	// A float number overflows to infinite
	f := float32(math.MaxFloat32)
	fmt.Println("Value of f: ",f)
	f = f * 1.2
	fmt.Println("f: ",f)


// Underflow
	var b int8 = -128
	b-- // underflow, value of b wraps to maximum value of int8 which is 127
	fmt.Println("Value of b: ",b)


}
