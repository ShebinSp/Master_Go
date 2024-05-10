package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is bar")
}

func foobar() {
	fmt.Println("This is foobar")
}
func main() {
	// DEFER STATEMENT

	// A `defer` statement defers or postpones the execution of a function until the surrounding function returns
	// either normally or through a panic.
	defer foo() // it will execute just befor exiting the surronding function which is main.
	bar()

	fmt.Println("Just a string after defering foo() and calling bar()")

	defer foobar()

	// Why defer?
	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // this is the recommended way to close a file in Go.

	// code that works (read/write) with the file

}
