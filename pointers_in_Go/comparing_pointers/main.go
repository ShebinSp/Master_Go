package main

import "fmt"

func main() {
	// COMPARING POINTERS

	// A pointer can point to a variable of any type, it can point to another pointer as well.
	// How to create a pointer to another pointer?
	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of p1: %v, address of p1: %v\n",p1,&p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n",pp1,&pp1)

	// How to dereference a pointer to a pointer?
	fmt.Printf("*p1 is %v\n",*p1)
	fmt.Printf("*pp1 is %v\n",*pp1)
	// Use the double dereferencing to get the value
	fmt.Printf("**pp1 is %v\n",**pp1)

	// If p1 is a pointer that points to int, can we increment that int value using p1? : Yes.
	**pp1++ // increments a
	fmt.Printf("value of a: %v\n\n",a) // 6.5

	//** COMPARING POINTERS**//
	fmt.Println("** COMPARING POINTERS**")
	// The zero value for a pointer of any type is nil

	// The test `p` not  equals to nil is true if `p` points to any variable, then pointers are comparable to pointers
	// are equal if and only if they point to the same variable or both are nil. 
	
	var p2 *int
	fmt.Printf("%#v\n",p2) // p2 which is a pointer to int is nil.

	y := 5
	p2 = &y

	z := 5
	p3 := &z
	// Are these 2 pointers p2 and p3 equal?
	// no because they are pointing to different variables, even those variables are equal.
	fmt.Println("Is p2 == p3?",p2 == p3) // false
	// Two pointers are equal if they are pointing to the same variable or both are nil.

	p4 := &y
	fmt.Println("Is p2 == p4?",p2 == p4) // true

}
