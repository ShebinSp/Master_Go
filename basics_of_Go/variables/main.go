package main

import "fmt"

func main() {
	// using var keyword
	var x = 7  // declared and assigned value
	var xx int // declared to use later
	y := 9     // short declaration
	xx = 8     // value assigned to 'xx'

	//_ = y // variable y is muted inorder to avoid compile time error, (unused variable)

	// adding a conditional statement to mute the 'y'
	if (x+xx)%2 == 1 {
		_ = y // if the addition value of x & xx is is odd then var y is muted
		fmt.Printf("X and XX is : %d , %d\n", x, xx)
	} else {
		y = x + xx // If the addition value of x & xx is even then value assigned to y
		fmt.Println("Y is : ", y)
	}

	car, cost := "Audi", 50000 // Declared multiple variables using short declaration operator
	fmt.Println("Car price: ", car, cost)

	car, year := "BMW", 2022 // car is declared but error muted becoz year is a new declaration

	_, _ = car, year // Both car & year is muted

	// An another way of multiple declaration
	// used to declaration variables with different types
	var (
		salary float32
		firstName string
		permanent bool
	)
	salary = 3000
	firstName = "abc"
	permanent = true
	fmt.Println("Employ Details: ", firstName, salary,permanent)

	// to declare variable in same type
	var a, b, c int
	fmt.Println("int type variables values are :", a,b,c)

	fmt.Println()
	// MULTIPLE ASSIGNMENTS & SWAPPING VALUES
	var i, j int

	i, j = 5, 7 // multiple assignments
	fmt.Printf("Normal values of I = %d , J = %d\n", i, j)

	i , j = j , i

	fmt.Printf("Swapped values: I = %d , J = %d\n",i, j)

	fmt.Println()
	// we can also use expressions in Short Declaration Operator
	sum := 5 + 2.5
	fmt.Println("Sum = ",sum)

}
