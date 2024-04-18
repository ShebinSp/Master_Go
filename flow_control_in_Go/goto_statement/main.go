package main

import "fmt"

func main() {
	// goto statement

	// using a label and goto statement the following piece of code creates a loop, like a flase statement does.
	i := 0
loop: // label, this is the *label `loop` point.
	if i < 5 {
		fmt.Println("value of i in loop: ", i)
		i++
		goto loop // It will transfer the execution of the program to the *label `loop` point(on line 10)
	}


	// It is not allowed to jump to a label after new variables have been introduced.

	//	goto todo // uncomment to view the error
	// x := 0 // uncomment to view the error
	// x := 0 // uncomment to view the error

todo:
	if i <= 5 {
		fmt.Println("value of i in todo: ", i)
		i++
		goto todo

	}
}
