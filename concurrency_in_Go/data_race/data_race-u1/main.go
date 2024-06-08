package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// DATA RACE

	// Launching 200 goroutines that work on the same shared value

	// 100 goroutines will increment the shared value and 100 goroutines will decrement the same shared value.
	// and they will do this at the same time.
	const gr = 100

	wg := sync.WaitGroup{}
	wg.Add(gr * 2)

	// shared value
	var n int = 0
	ctr := 0
	fmt.Println("Initial value of n: ", n)

	// starting 200 goroutines using a for loop
	for i := 0; i < gr; i++ {
		// It's not necessary to create a standalone function for the goroutines, we can use an anonymous function also
		// called a function literal.

		// Increment
		go func() {
			// we introduce some artificial time to simulate a more complex operation.
			time.Sleep(time.Second / 10) // it will wait for a 10th of a second or for 100 milliseconds.
			n++                          // incrementing the value of n
			ctr++

			fmt.Println("+++++++")
			fmt.Println("n: ",n)
			fmt.Println("count: ",ctr)
			wg.Done()                    // notifying the waitgroup that a goroutine is done.
		}() // the pair of parathensis is for function invocation

		// Decrement
		go func() {
			time.Sleep(time.Second / 10)
			n-- // decrementing the value of n
			ctr++

			fmt.Println("-------")
			fmt.Println("n: ",n)
			fmt.Println("count: ",ctr)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of n: ", n)
	fmt.Println("Total count: ",ctr)

	/*
	The value of `n` expected to be 0 but is 0, 1 or 2 which is not stable. Because all 200 goroutines work on the same
	shared value, which is `n`. The value of `n` is incremented, respectively decremented in seperated goroutines.
	Behind the scene is, Imagine that a goroutine reads the value of `n`, which is 1, it has just incremented and
	tries to decrement it, so make it 0, at the same time, another goroutine had already read n in order to increment
	it from 0 to 1. Now, depending on which goroutine finishes first, the final value fo `n` will be 0 or 1. We don't
	know which goroutine will finish first, so we don't know the final value of `n`. So depending on the order in which
	these 200 goroutines complete, the final value of `n` will be unpredictable. This is why it's called a data race.
	The value of `n` changes depending on the order in which the goroutines finish.
		Notice the `+`, `-` and `count` to get more clarity.
	*/
}
