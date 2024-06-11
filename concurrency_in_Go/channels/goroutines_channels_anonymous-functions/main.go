package main

import (
	"fmt"
)

func main() {
	//** GOROUTINES, CHANNELS AND ANONYMOUS FUNCTIONS**\\

	// It is typical in Go to use a function literal also called anonymous function to spawn a new goroutine.
	// Here is an example of a goroutine that uses an anonymous function using the same factorial example.

	// Instead of creating a standalone function for calculating the factorial lets use an anonymous one

	fmt.Println("Enter the number to find the factorial")
	var n int

	ch := make(chan int)
	defer close(ch)

	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		fmt.Println("Input an integer number")
	} else if n < 0 {
		fmt.Println("Factorial does not exist for negative numbers")
	} else if n == 0 {
		fmt.Println("The factorial of 0 is 1")
	} else {

		for i := 1; i <= n; i++ {
			go func(n int, c chan int) {
				f := 1
				for i := 2; i <= n; i++ {
					f *= i
				}
				c <- f
			}(i, ch)
			fmt.Printf("Factorial of %d is %d\n", i, <-ch)
		}
	}
}
