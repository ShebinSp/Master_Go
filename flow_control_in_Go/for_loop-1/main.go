package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// for loops are used to do repeat a task or a block of code multiple times

	// Example 1
	// uncomment below code to run the loop 10 times
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// Example 1
	// lets take an argument from the CLI.
	if args := os.Args; len(args) != 2 {
		fmt.Println("Need an interger argument to run the loop")
	} else if n, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("Please use only interger greater than zero! Error: ",err)
	} else {
		fmt.Printf("The loop will run %v times\n",n)
		for i := 1; i <= n; i++ {
			fmt.Println("Loop count: ",i)
		} // run code like this-> go run main.go 15
	}
}
