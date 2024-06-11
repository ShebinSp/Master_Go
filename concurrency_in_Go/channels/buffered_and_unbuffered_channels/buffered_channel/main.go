package main

import (
	"fmt"
	"time"
)

func main() {

	//** BUFFERED CHANNEL**\\

	// c is a buffered channel because it has a capacity set at declaration.
	bufferedChannel := make(chan int, 3)
	// The sender of a buffered channel will block only when there is no empty slot in the channel, while the receiver
	// will block on the channel when it's empty.

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n",i)
			c <- 10 + i
			fmt.Printf("func goroutine #%d after sending data into the channel\n",i)
		}
		close(c) // even though it's closed, we can receive values that are waiting in the channel.
		// By the way, you don't need to close every channel when you are done with it. It's only necessary to close
		// a channel when it is important to tell the receiving goroutines that all data have been sent.

	}(bufferedChannel)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)
	// The goroutines start sending data into the channel before the main() goroutine had a chance to receive data from
	// the channel. The sender fo this buffered channel will block only when there is no empty slot in the channel. In
	// this case, after three writing attempts because the channel has a capacity of three.


	// The typical way to receive data from a channel where more goroutines are sending data to is to use a for loop
	// and iterate the channel. We can send up to three values on the channel without the goroutine blocking and the
	// receiver will block on the channel when it's empty. If the channel is neither full or empty, either a send or
	// receive operation could proceed without blocking. The buffered channel has a FIFO or the First In First Out
	// queuing discipline.
	for v := range bufferedChannel { // v := <- bufferedChannel
		fmt.Println("main goroutine received value from channel: ",v)
	}

	// What is the difference between a buffered and unbuffered channel?
	/*
	We say that communication over an unbuffered channel, the example from the previous lecture causes the sending and
	receiving goroutines to synchronize. Because of this unbuffered channels are sometimes called synchronous channels.
	By the way, the channel names are named as their type here. If you try a receive operation on a closed channel, 
	it will proceed without blocking and yield a zero value for the type that is sent through the channel. In this case
	the zero value for int, which is 0.
	*/
	fmt.Println(<- bufferedChannel) // 0
	// bufferedChannel <- 11 // panic: send on closed channel

}
