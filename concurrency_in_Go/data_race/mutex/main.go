package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Mutex

	// How to fix bug and write concurrently safe code

	const gr = 100

	wg := sync.WaitGroup{}
	wg.Add(gr * 2)

	// shared value
	var n int = 0
	fmt.Println("Initial value of n: ", n)
	// we have decided that the value of `n` should only be read after the write operation has finished. Here we don't care-
	// about the order of reads and writes. We only require that they do not occur simultaneously.

	// 1.
	var m sync.Mutex // declared an instance of mutex

	// starting 200 goroutines using a for loop
	for i := 0; i < gr; i++ {

		go func() {
			time.Sleep(time.Second / 10)
			// 2.
			// The `lock` method of the mutex variable blocks the access to the variable until the `unlock` method is called.
			m.Lock() // the code b/w lock and unlock will be executed only by one goroutine at a time.
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			// m.Unlock()
			wg.Done()
		}()
		// The instruction where n is incremented respectively decremented will be executed by only one goroutine at any 
		// point in time. This will prevent any race condition.
		// If one goroutine already holds the lock and if a new goroutine is trying to acquire the lock,
		// the new goroutine will be blocked until the mutex is unlocked.
	}
	wg.Wait()
	fmt.Println("Final value of n: ", n)
}
