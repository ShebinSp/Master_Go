package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	i, err := strconv.Atoi("20x") // we will get an error here
	// pointing out the error if there is any
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("i is: ",i)
	}

	fmt.Println("Example 1")

	// lets try above statement as simple statement
	// first part of if(i, err := strconv.Atoi("7777") is known as an initialization statement and it can be used to setup local variables.
	// for functions that return an error the initialization statement is very useful. The error is scoped only the `if` block that-
	// handles it. This way, the code ends up being readable.
	// The second part of the statement is a boolean expression that returns true or false. If it's true, the branch below will be executed.
	// If there is no error, `if err == nil` it will print integer. If there is an error, the else case will be executed.
	if i, err := strconv.Atoi("777"); err == nil {
		fmt.Println("No error, i is: ",i)
	} else {
		fmt.Println("Error: ",err)
		// handle the error
	}
	
	fmt.Println("Example 2")
	if args := os.Args; len(args) != 2 { // args is type string
		fmt.Println("One argrment is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil { // converting string to int
		fmt.Println("The argument must be in integer! Error: ",err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km) * 621)
		fmt.Printf("Type of args[1]: %T\n",args[1])
		fmt.Printf("Type of km: %T\n",km)

	} // try arguments: 1. go run main.go, 2. go run main.go hello, 3. go run main.go 3

}
