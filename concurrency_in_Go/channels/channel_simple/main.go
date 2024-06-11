package main

import "fmt"

func f1(n int, ch chan int) { // receiving an interger value and a bidirectional channel
	ch <- n // sending the value of n into the channel
}

func main() {
	// Bidirectional and Unidirectional Channels

	c := make(chan int) // channel to send and receive (bidirectional channel)
	defer close(c)

	// channel only for receiving (unidirectional channel)
	c1 := make(<- chan string) // channel used only to receive string

	// channel only for sending (unidirectional channel)
	c2 := make(chan <- string) // channel used only to send string

	fmt.Printf("c: %T   c1: %T   c3: %T\n",c,c1,c2)


	// A very, very simple example that uses a channel to communicate between two goroutines.

	go f1(10, c) // launching f1 function as goroutine and sending a value and bidirectional channel to f1() as arguments.
	n := <- c // assigning the value received through channel to variable n.

	fmt.Println("Value received: ",n)
	fmt.Println("Exiting main...")
}