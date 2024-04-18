package main

import "fmt"

func main() {

	// Label Statement
	people := [5]string{ "Mark", "Brenda", "Helen", "Antonio", "Michael"}
	friends := [3]string{"Marry", "Mark", "Helen"}

outer: // Label statement : Most of the time labels are used to terminate outer enclosing loops.
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d,\n", friend, index+1)
				break outer // usually a break statement will exit the inner or where the break is stated. But Label statement-
				// used with break to exit or terminate the outer loop. Means if there is a match the whole loop will be terminated.
			}
		}
	}

	outer := 1 // its a variable called outer. Labels don't conflict with other names or identifiers.
	_ = outer

}
