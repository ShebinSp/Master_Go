package main

import (
	"fmt"
)

func main() {
	// STRINGS IN GO

	// Strings are defined between double quotes
	s1 := "Hi there Go!" // what is between double quotes called string literals
	fmt.Println("s1: ", s1)

	// If you want to use double quotes inside the double quotes you must use a backslash in front of each
	// double quotes to escape it or using backticks `He says: "Hello"`
	fmt.Printf("He says: \"%s\"\n", s1)
	fmt.Println("He says: \"Hello there Go!\"")

	// A string literal enclosed in backticks is called a raw string and it is interpreted literally.
	// Back slashes or other special characters have no special meaning.
	fmt.Println(`He says: "Hii hello!"`) // this is a valid Go code.

	// lets create a raw string where escape sequences are not interpreted
	s2 := `I like Go!` // raw string, which means if there is \n it has no effect.
	fmt.Println("s2: ", s2)
	// to make clear
	fmt.Printf(`Here comes the \n next line`) // \n will not start a new line
	fmt.Println()                             // seperate Println() to start a new line

	// write strings on multiple lines
	fmt.Println("Price: 100\nBrand: Puma")
	// this is same as
	fmt.Println(`Price: 100
Brand: Nike`) // you should avoid tabs or spaces in between backticks for a better format.

	// if you want to use back slashes in string use a pair of backticks
	fmt.Println(`C: \Users\User1`)
	// or you want double quotes
	fmt.Println("C: \\Users\\User1") // each backslash will escape the next one

	// To concatenate strings use the addition operator or the plus sign
	s3 := "I love " + "Go " + "Programming" // this is a string concatenation
	fmt.Println("String concatenation: ", s3 + "!")

	// lets get an element from a string using indexing
	fmt.Println("Element at index 0: ",s3[0]) // what you will get is the ASCII value of I. because a string is a slice ofbyte in Go
	// To print an index use the string() function
	fmt.Printf("Element at index 0: %s\n",string(s3[0])) // byte converted into string
	
	// strings are immutable
	// s3[5] = x // compile time error: cannot assign to s3. A string is immutable in Go.

	// If you want to print a string use %s or %q for quoted string
	fmt.Printf("normal: %s\n",s3)
	fmt.Printf("quoted: %q\n",s3)

}
