package main

import "fmt"

// A function that takes a pointer as argument
func change(a *int) *float64 {
	// Pointes are copied too, Functions work on copies, not on originals in Go.
	*a = 100
	// In Go, it's perfectly legal for a function to return a pointer to local variable.
	b := 5.
	return &b
}

func changVar(a int){
	a = 77
}

func main() {
	// PASSING POINTERS TO FUNCTIONS

	x := 8
	p := &x
	fmt.Println("Value of x before calling change(): ", x)
	change(p)
	fmt.Println("Value of x after calling change(): ",x)
	fmt.Println("Value of x before calling changeVar(): ",x)
	changVar(x) // not changed the value of x, the function would have worked on copies, not on originals.
	fmt.Println("Value of x after calling changeVar(): ",x)



}
