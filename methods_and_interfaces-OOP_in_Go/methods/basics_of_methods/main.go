package main

import (
	"fmt"
	"time"
)
// Creating a custom type and method

// creating a new type called names, which is a slice of strings
type names []string

// Creating a receiver function or method for this newly defined type.
// Any value of type names can get access to this function.
// A method is declared as a variant of the ordinary function declaration in which an extra parameter-
// appears before the functions name. The parameter connects the function to the type or that parameter.
func (n names) print() {
	for i, name := range n {
		fmt.Println(i+1, " ", name)
	}
}

// In this example `n` is called a method's receiver and is like an input parameter, just that it's written
// before a function's name and not between parentheses that come after the function's name.
// `n` is the actual copy of the names we are working with and is available in the function and is like-
// `this` or `self` from Object Oriented Programming.
// Any variable of type `names` can call this function on itself like `variable_name.print`.
// In Go, we don't use special name like `this` of `self` for the receiver.
// We choose Receiver's name just as we would for any other parameter. It is a good idea to choose something-
// short an to be consistent across all methods. A common choice is the first letter of the named type,
// like `n` for the `names` and `f` for the `fahrenheit`.

func main() {
	/**BASICS OF MEHTODS**/

	const day = 24 * time.Hour
	fmt.Printf("Type of constant day: %T\n", day)
	fmt.Println("d: ", day)

	// Let's use a method from the standard library to see how it works.
	// we'll use the `seconds` method of time.duration type
	seconds := day.Seconds() // this method is also known as a receiver function.
	// A method is called on a named  type value, in this case, on a day which is of type time.duration.
	// Seconds() returns the duration which is a floating point number of seconds.
	fmt.Printf("Type of seconds: %T\n", seconds)
	fmt.Printf("Seconds in a day: %v\n", seconds)
	// In above example we've seen that named types can have functions also called methods or receiver functions
	// attached to them.

	// Creating a variable called friends of type `names`.
	friends := names{"Linto", "Nidheesh", "Jibin"}
	// To call the method we use the names type `friends.method_name`. That is the selection operator and it selects
	// a name from a namespace like from a struct, a package or a type.
	// In method called the receiver argument in our case, `friends` appears before the method name.
	friends.print()
	fmt.Println()

	// Another way to call the method is
	names.print(friends) // this is not very common to call a method but it is possible.
	// In fact, a method is a function that takes a receiver as argument.
	// The major difference between a method and a function is that a method belongs to a type, and the function-
	//belongs to a package.

	/* It's idiomatic in Go to convert the type of an expression to access a specific method.
	   Let's suppose we have an int64 value and want to convert it to time.duration.
	   Remember that we can convert different types only if the share the same underlying type.
	*/
	var n int64 = 93475932747
	fmt.Println("n: ",n)
	fmt.Println("Converted n to time.Duration: ",time.Duration(n))

}
