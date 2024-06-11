package main

import (
	"fmt"
	"time"
)

func main() {
	//** UNBUFFERED CHANNELS **\\

	// A channel created with a simple call to make() function is called an unbuffered channel.

	var unbufferedChannel = make(chan int) // unbuffered channel
	_ = unbufferedChannel

	// Buffered Channel
	// The `make()` function accepts an optional second argument of type `int` called the channel capacity.
	// If the capacity is non-zero, `make()` creates a buffered channel.
	var bufferedChannel = make(chan int, 3) // buffered channel
	_ = bufferedChannel

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 7
		fmt.Println("func goroutine after sending data into the channel")

	}(unbufferedChannel)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <- unbufferedChannel
	fmt.Println("main goroutine received data: ",d)

	// For an unbuffered channel, the sender the func goroutine blocks on the channel until the receiver the
	// main() goroutine receives the data from the channel. The message func goroutine sending data into the channel is
	// printed out only after the main goroutine has already received the value.
	// That's because, the line 24 blocks the goroutine until main wakes up. Its sleeping and cannot read from the channel.
	
	time.Sleep(time.Second)
}
