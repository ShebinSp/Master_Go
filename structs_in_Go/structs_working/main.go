package main

import "fmt"

func main() {
	// WORKING WITH STRUCTS

	type book struct {
		title  string
		author string
		year   int
	}
	// decalring and initializing a variable of type book
	lastBook := book{title: "Anna Karenina"}
	// Retriving the field title
	fmt.Println("Last book title: ", lastBook.title)
	// if you want the value to a variable
	lastBookTitle := lastBook.title // lastBookTitle type is string
	_ = lastBookTitle

	// Structs in Go fixed at compile time
	// You can't add or remove fields from a struct at runtime.

	// Uptate the field
	fmt.Printf("last book: %#v\n", lastBook)
	// updating the struct field `author` and `year`
	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("last book: %+v\n", lastBook) // z%+v prints both the field names and their values

	// How to compare the struct values?
	// We can compare struct values using double equals sign and not equals.
	// Two structs are equal if their corresponding fields are equal.

	someBook := book{title: "Anna Karenina"}
	fmt.Println("someBook equal to aBook?: ", someBook == lastBook)
	aBook := book{title: "Anna Karenina", author: "Leo Tolstoy", year: 1878}
	fmt.Println("lastBook equal to aBook?: ", aBook == lastBook)

	// How to copy a struct
	// When we want to make a copy or a clone of a struct, we can use the equal sign or the assignment operator.
	// Go will create a brand new struct value in memory.
	myBook := aBook
	myBook.year = 2022
	fmt.Println("myBook: ", myBook, "  		   aBook: ", aBook)
}
