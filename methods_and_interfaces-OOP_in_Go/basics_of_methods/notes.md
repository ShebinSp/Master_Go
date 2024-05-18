# Methods and Interfaces- OOP in Go

* Go doesn't have classes, but you can define methods on predefined types.

* A type may have a method set associated with it which enhances the type with extra behaviour. This way, a named type has both data and behaviour and represents better a real world concept.

* A method is called on a named  type value

* Creating a new type called names, which is a slice of strings

`type names []string`

* Creating a receiver function or method for this newly defined type.

* Any value of type names can get access to this function.

* A method is declared as a variant of the ordinary function declaration in which an extra parameter appears before the functions name. The parameter connects the function to the type or that parameter.

`func (n names) print() {`
    `for i, name := range n {`
        `fmt.Println(i+1, " ", name)`
   ` }`
`}`

* creating a variable called friends of type `names`.

`friends := names{"Linto", "Nidheesh", "Jibin"}`

* Calling the method of type `names`.

`friends.print()`

### It's idiomatic in Go to convert the type of an expression to access a specific method.

* Let's suppose we have an int64 value and want to convert it to time.duration.

* Remember that we can convert different types only if the share the same underlying type.

`var n int64 = 93475932747`
    `fmt.Println("n: ",n)`
    `fmt.Println("Converted n to time.Duration: ",time.Duration(n))`

* output: 
n:  93475932747
Converted n to time.Duration:  1m33.475932747s
