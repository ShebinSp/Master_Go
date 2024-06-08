package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// How to use the built-in race detector of Go to find the race condition

	// To use race detector add the race flag `-race` to the normal go command line tool.`go run -race main.go`
	// When the race command line flag is set, the compiler instruments all memory access with code that
	// records when and how the memory was accessed. While the runtime library watches for unsynchronized
	// accesses to shared variables.

	const gr = 100

	wg := sync.WaitGroup{}
	wg.Add(gr * 2)

	// shared value
	var n int = 0
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
			wg.Done()                    // notifying the waitgroup that a goroutine is done.
		}()

		// Decrement
		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of n: ", n)
}

/*
$ go run -race main.go
Output: 'The out put is devided into 3 sections'

==================
* The FIRST SECTION tells that there was an attempt to read the value of `n` in order to decrement it
  inside a goroutine that we created. This is line 42 where the second goroutine tries to decrement n.

	WARNING: DATA RACE
	Read at 0x00c0000a4018 by goroutine 10:
	  main.main.func2()
	      /home/katana/go/src/Master_Go/concurrency_in_Go/data_race/race_detector-u2/main.go:42 +0x3d

* The SECOND SECTION tells that there was a simultaneous write by another goroutine. This is on line 35
  where the goroutine tries to write back the value of `n` it has just incremented.

	Previous write at 0x00c0000a4018 by goroutine 23:
	  main.main.func1()
 	     /home/katana/go/src/Master_Go/concurrency_in_Go/data_race/race_detector-u2/main.go:35 +0x4f

* The THIRD SECTION describes where the goroutine that caused the data race was created on lines 40 & 23.
	Goroutine 10 (running) created at:
	  main.main()
 	     /home/katana/go/src/Master_Go/concurrency_in_Go/data_race/race_detector-u2/main.go:40 +0x13d

	Goroutine 23 (finished) created at:
 	 main.main()
 	     /home/katana/go/src/Master_Go/concurrency_in_Go/data_race/race_detector-u2/main.go:32 +0x204
	==================
*/
