# EMPTY INTERFACE

* As it's name, suggest, it is an interface with no method. It's completely empty.

* Any type satisfies this interface, and that means it can represent any value. This is very important and makes the empty interface very useful sometimes.

* An empty interface may hold values of any type, even though it seems very simple, The empty interface is a key concept in Go.

* We cannot use an empty interface in operations. An interface stores in fact two values; a dynamic type and a dynamic concrete value. To access the dynamic value, we have to do an assertion. If we want to use a function that accepts a slice as argument  or a method that works on a slice value, we must retrieve the dynamic value using an assertion.

`fmt.Println("Slice in empty: ", len(empty.([]int)))`
`// This is how we get the dynamic value stored in the interface.`

### Why Empty Interface useful?

* Empty interfaces are used by code that handles values of unknown type, and you can pass an empty interface type as a function parameter of any type.

### Problems of Empty Interface

* Empty interfaces can cause programs to become hard to maintain.

* The empty interface type is increasingly being misused as a convenient way to bypass go compilers type checking and one of the principle of Go it that us allows us to write type safe code.

* There are places where empty interface is used in place where explicit interface could have been used. The problem with runtime type checking is that you will never know if there is a problem until it is run.

* Use empty interfaces only if it's really necessary.