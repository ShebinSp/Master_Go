package main

import "fmt"

func main() {
	// continue statement

	// This for loop prints even numbers from 1 - 10.
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {		// if the value of `i` is an even number, then value of `i` prints
			fmt.Println(i)
		} else {   // if not, the iteration of loop is skipped and the control will go to post statement(i++).
			continue
		}
	}
}
