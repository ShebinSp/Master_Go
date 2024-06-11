package main

import "fmt"

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	c <- f
}

func main() {
	//** GOROUTINES AND CHANNELS **\\

	// An example of many goroutines that run concurrently and communicate in between using a channel. In this example
	// each goroutine to compute the factorial of `n`. By the way, the factorial of `n` is the product of all positive
	// integers leass than or equal to `n`.

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
		
		go factorial(n, ch)

		// After the factorial goroutine is spawned, main() should wait for a message to come through the channel, this is
		// called a blocking call.
		// The main goroutine is put to sleep and waits for factorial goroutine to send a message through the channel.
		f := <-ch // this is a blocking operation.

		// After the main goroutine receives the message, it wakes up, prints out the message and exits entirely.
		if f == 0 {
			fmt.Printf("The factorial of %d is a complex number\n", n)
		} else {
			fmt.Printf("The factorial of %d is %d\n", n, f)
		}
	}

	// If there are more goroutines, the main routine should wait for a message through the channel from each goroutine.
	// lets start 20 goroutines that calculate the factorial.
	
	for i := 1; i <= n; i++ {
		go factorial(i, ch)
		f := <- ch

		if f == 0 {
			fmt.Printf("The factorial of %d is a complex number\n", n)
		} else {
			fmt.Printf("The factorial of %d is %d\n", i, f)
		}
	}
}
