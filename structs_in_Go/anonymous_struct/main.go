package main

import "fmt"

func main() {
	// ANONYMOUS STRUCT and ANNONYMOUS FIELDS

	// An ANONYMOUS STRUCT is a struct with no explicitly defined struct type alias.
	// Anonymous structs are useful when you don't want to reuse a struct type.
	diana := struct {
		firstName, lastName string
		age int
	}{
		firstName: "Diana",
		lastName: "Muller",
		age: 30,
	}

	fmt.Printf("Anonymous Struct: %#v\n",diana)
	fmt.Printf("%s's age: %d\n",diana.firstName, diana.age) // retriving first name and age of diana


	// ANONYMOUS FIELDS
	// We can define struct type without any field name.
	// We can just mention the field types and Go will use type as field name.
	// It is not allowed to have more than one type as anonymous fields.
	type book struct {
		string
		float64
		// string // repeating a type is an error
		bool
	}

	b1 := book{"1984 by George Orwell", 10.3, false}
	fmt.Printf("book b1: %#v\n",b1)
	// to retrive a value we use the type we want
	fmt.Println("Book: ",b1.string)

	// It is allowed to mix anonymous fields with named fields.
	type employee struct {
		name string // named field
		salary int // named field
		bool // anonymous field
	}

	e := employee{"John", 3000, false}
	fmt.Printf("e: %#v\n",e)
	// If you want change the anonymous field
	e.bool = true
	fmt.Printf("e: %#v\n",e)
}