package main

import (
	"fmt"
	"time"
)

func main() {
	//** select STATEMENT **\\

	c1 := make(chan string)
	c2 := make(chan string)

	start := time.Now().Second()

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Hii!"
	}()

	time.Sleep(time.Second * 2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Receiceved: ", msg1)
		case msg2 := <-c2:
			fmt.Println("Received: ", msg2)
		default:
			fmt.Println("no activity")
		}
	}
	/*
	In this case, main receives a message from each goroutine. Basic sends and receves on channels are blocking.
	However, we can use select with a default clause to impliment non-blocking sends, receives and even non-blocking
	multi ways selects.
	*/

	end := time.Now().Second()
	fmt.Println("Total time taken: ",end - start, "seconds") // 2 seconds

	// This is the total execution time since both the first and the second sleeps execute concurrently.
	// We say that the select statement is used when we want to wait on multiple goroutines simultaneously.



}
