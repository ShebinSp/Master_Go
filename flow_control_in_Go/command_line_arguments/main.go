package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args",os.Args) // access all the arguments
	fmt.Println("Path: ",os.Args[0]) // access the path
	fmt.Println("First argument: ", os.Args[1]) // access the first argument
	fmt.Println("No of items inside os.Args: ",len(os.Args)) // total number of arguments
	fmt.Println("-----------------------------")
	
	// Example
	// converting string argument from os.Args[1] to float64
	result, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("\nType of os.Args[1]: %T\n",os.Args[1])
	fmt.Printf("Type of result: %T\n",result)
	fmt.Printf("Value of result: %v\n",result)
}
