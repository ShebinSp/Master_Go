package main

import "fmt"

func main() {
	// BASICS OF STRUCTS

	// Let's suppose our application will store data about books. To do that we create some variables and initialize
	// those variables.
	title1, author1, year1 := "The Devine Comedy", "Davte Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	// we see the information about our books
	fmt.Println("Book1: ",title1, author1, year1)
	fmt.Println("Book2: ",title2, author2, year2)
	// The problem with an application like this is that there is no connection between the variables title, author and 
	// year. They dont act as a single unit and don't represent the book concept from the real world. and if there are
	// 1000 books we have to create 3000 different variables to store the information.

	// A much more better approach is to create a struct that groups together these different concepts into a single
	// concept called book.

	// Each field must be unique in the struct
	type book struct {
		title string
		author string
		year int
	}
	// creating an instance or object of type `book`.
	// you can omit the field name, but you have to keep the order and it is not a recomended way to create a struct.
	myBook := book{"The Devine Comedy", "Davte Aligheri", 1320}
	// When we are seeing struct, we are referring to the variable which holds the type of book. So `book` is struct-
	// type and `myBook` is struct. And struct is a built-in type in Go.
	fmt.Println("myBook: ",myBook)

	// with field name order doesn't matter and it is the best practice to initialize struct.
	bestBook := book{
		title: "Mackbeth",
		year: 1606, // year specified second
		author: "William Shakespeare",
	}
	fmt.Println("Best book: ",bestBook)

	// If we omit some field when creating struct, then omited fields will be zero valued according to their type.
	aBook := book{
		title: "Just a random book",
	}
	fmt.Printf("a book: %#v\n",aBook) // author and year will be zero valued whcih are "" and 0.


}
