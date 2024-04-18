package main

import "fmt"

func main() {
	var x = 3   // int type
	var y = 3.1 // float type

	// x = x * y // Error

	res := x * int(y)
	// You can notice that Go didn't change the type of y
	fmt.Printf("Result: %v and Type of y is %T\n", res, y)
	// If you assign the converted value of y then the type will chage according to the function
	yy := int(y)
	fmt.Printf("Type of yy is %T\n", yy)

	// converted value of x to float64 and multiplied to y and value assigned to x
	x = int(float64(x) * y)
	fmt.Printf("Value of x: %v and Type: %T\n", x, x)

	var a = 5 // int type
	fmt.Printf("Type of a: %T\n",a)
	var b int64 = 2

	// int64 and int are not same type even though they may have the same size on a particular architecture
	// a = b // Error: uncomment to view the error
	a = int(b) // now the type of b int64 converted to int
}
