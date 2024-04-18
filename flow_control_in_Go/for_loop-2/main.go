package main

import "fmt"

func main() {
	// while loop in Go

	j := 10 // initialization statement
	for j > 0 { // conditional expression
		fmt.Println("Value of j: ",j)
		j-- // post decrement statement
	}

	// infinite loop
	// Below code is neverending
	// click on the terminal and press `ctrl + c` to stop the program.
	sum := 0
	for {
		sum++
		fmt.Println("Sum: ",sum)
	}
	fmt.Println("Value of sum is: ",sum) // This statement will not be reached because The infinite loop is stil running
	 
}
