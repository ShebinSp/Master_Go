package main

import (
	"fmt"
)

func main() {
	const seconds = 60
	const munites = 60
	const hours = 24
	const days = 7
	fmt.Println("Seconds in a day: ", (seconds*munites)*hours)
	fmt.Println("Seconds in 1 week: ", (seconds*munites*hours)*days)

	// To declare multiple constants. Constants can be declared with or without their type
	const m, n, o int = 1, 2, 3 // typed constants 
	const m1, n1, o1 = 4, 5, 6 // untyped constants

	// This method can be used for multiple different type constants(group constants)
	const (
		nn  float32 = 1.1
		mm          = 1
		s           = "s"
		boo         = true
		ss  string  = "ss"
	)

	// In group constants if only first value is assigned then, next constants will be
	// assigned with the previous value as their value.
	const (
		num1 = 500
		num2
		num3
	)
	fmt.Println("Constants are: ", num1, num2, num3)

	// Constant rules
	// 1. You can't change a constant on runtime.
		// const num = 7
		// num = 5 -> ERROR,.

	// 2. You can't initilize a constant at runtime. It belongs to compile time.
		// const power = math.Pow(2, 3) -> ERROR, constants belongs to compile time & functions
		// belongs to runtime.
	
	// 3. You can't use a variable to initialize a constant
		// Because variables belong to runtime and constats to compile time. So you can't
		// initialize a constant to runtime values.
		// t := 5
		// cont tc = t // ERROR.

	// 4. We can use the length of a string as constant using len() function. and it should
		// be a string literal because a string literal is an unnamed constant,
		// not a string variable. 
		// Here we can initilize a const using len() funtion, because len() is a built-in 
		// function which is available at compile time and math.Pow() is a runtime funtion.
		const l1 = len("hello")
		fmt.Println("Value of l1: ",l1)
}
