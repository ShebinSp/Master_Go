package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Function recevies a WaitGroup value as an input parameter. The parameter must be passed as a pointer.
func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) started")
	for i := 1; i <= 3; i++ {
		fmt.Println("f1, i = ", i)
		time.Sleep(time.Second * 1)
	}
	fmt.Println("f1 (goroutine) finished")
	wg.Done()
}
func f2() {
	fmt.Println("f2 started")
	for i := 1; i <= 5; i++ {
		fmt.Println("f2, i = ", i)
	}
	fmt.Println("f2 finished")
}

func main() {
	// WAITGROUPS

	// To create a WaitGroup
	var wg sync.WaitGroup // this WaitGroup is used to wait for all the goroutines that have been launched to finish
	// After creating the weight group, we call the `Add()` method before attempting to execute the goroutine.
	wg.Add(1) // 1 is the number of goroutines to wait for.
	fmt.Println("main execution started")

	// The WaitGroup must be passed as a pointer.
	go f1(&wg)

	fmt.Println("Number of goroutines after go f1(): ", runtime.NumGoroutine())

	f2()

	// we call `wg.Wait` to block the execution of the main() function until the goroutines in that WaitGroup
	// successfully completed.
	wg.Wait()

	// notice that main() function wait to f1() function to finish executing.
	fmt.Println("main execution stopped")

}
