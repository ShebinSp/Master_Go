package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
		c4 = iota
	)
	fmt.Println(c1, c2, c3)

	// we can simplify it 
	const (
		c11 = iota + 1 // here iota starting from 1
		c12
		c13
	)
	fmt.Println(c11,c12,c13)

	// If you want to increment 2 afrer each iota
	const (
		a = iota * 2
		b 
		c
		d 
		e
	)
	fmt.Println(a,b,c,d,e)


	// The value of iota begins default by zero, or by the number of position of the constant in the group constant
	// where you assigned the value as iota. so If you want to skip a value using blank identifier in iota
	 
	const (
		x = -(iota + 2) // -2
		_				// -3
		y				// -4
		z				// -5
	)
	fmt.Println(x,y,z)

}
