package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 (goroutine) started")
	for i := 1; i <= 3; i++ {
		fmt.Println("f1, i = ", i)
	}
	fmt.Println("f1 (goroutine) finished")
}
func f2() {
	fmt.Println("f2 started")
	for i := 1; i <= 5; i++ {
		fmt.Println("f2, i = ", i)
	}
	fmt.Println("f2 finished")
}

func main() {
	// INTRODUCTION TO GOROUTINES

	// SPAWNING GOROUTINES, THE go KEYWORD.
	fmt.Println("main execution started")

	// Important values in Concurrency and Parallelism

	// NumCPU() returns the number of logical CPUs or cores usable by the current process.
	fmt.Println("Number of CPUs: ", runtime.NumCPU())

	// NumGoroutine() returns the number of goroutines that currently exists.
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())

	// OS
	fmt.Println("OS: ", runtime.GOOS) // GOOS is environment variable
	// Architecture
	fmt.Println("Architecture: ", runtime.GOARCH) // GOARCH is environment variable

	// How many operating system threads may be executed in parallel on this Operating System?

	// The GhostScheduler uses a parameter called GOMAXPROCS to determine how many OS threads may be actively
	// executing go code simultaneously. Its default value is the number of CPUs on the machine. So on a
	// machine with four CPU cores, the scheduler will schedule Go code on upto four OS threads at once.
	// GOMAXPROCS is the `n` in `n:m` scheduling. This parameter can be explicitly controlled by using
	// GOMAXPROCS environment variable or the `runtime.GOMAXPROCS()` function.
	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0)) // If n is less than 1 it doesnot change the current setting.

	// Creating first goroutine
	// A goroutine is in fact a function called using a go keyword before its name. lets create 2 functions-
	// f1() and f2().

	// Launch goroutine for f1() function
	go f1()
	// number of goroutines are running
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine())

	// By default, every go standalone application creates one goroutine. This is the main goroutine or main() function.
	// When calling f1() using go prefix, it spawns a new goroutine that runs concurrently with main(). However the
	// main() function doesn't wait for the goroutine to execute or to finish. The main() function actually terminates
	// before the goroutine gets a chance to execute. When main() function executes, the entire program exists.
	// We need some sort of synchronization between the main goroutine and f1 goroutine. Or in other words, we want to
	// wait for the goroutine to finish its execution before continuing. So if we follow the logic and introduce some
	// wainting time in the main function before exiting, we give a chance to f1 goroutine to start and execute.

	// we make the main goroutine sleep for 2 seconds, the f1 goroutine will execute at this time.
	time.Sleep(time.Second * 2)

	// However, this was just a proof of concept. In real concurrent applications, we don't use time.Sleep() to
	// wait for goroutines. because if the goroutine needs more than 2 seconds to complete its job, The main 
	// goroutine would exit without waiting for the other goroutine.


	// Calling f2() function
	f2()

	fmt.Println("main execution stopped")

}
