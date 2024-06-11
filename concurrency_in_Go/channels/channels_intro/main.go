package main

import "fmt"

func main() {
	// Channel

	// declaring a channel
	var ch chan int // the value of uninitialized channel is nil.
	fmt.Println(ch) // nil

	// initializing a channel
	ch = make(chan int)
	// A channel is like a pointer, and passing channels to functions has the same effect as passing pointers to functions.
	fmt.Println(ch) // 0xc000018120. It is a memory address

	// we can declare and initialize the channel at the same time using the short variable declaration operator
	ch1 := make(chan int) // channel declared and initialized.
	_ = ch1

	// <- channel operator
	// SEND
	ch <- 10 // sending 10 to the channel ch.

	// RECEIVE expression in an assignment statement
	num := <- ch // we wait value to be send into the channel, and when we get one we assign it to the num.
	_ = num

	// we can also wait for a value to be sent into the channel and print out that value when we get
	fmt.Println(<-ch)

	close(ch) // closes the channel

	// If we try to run the program at this time we will get a Fatal error called deadlock, because channels are using in
	// conjunction with goroutines. And in this program there is no goroutine started.



	
}
