# Functions in Go

* A function is a small piece of code that is dedicated to perform a particular task based on some input values.

* Go recommends writing function name in simple word or cammelCase. Even though underscore or snakecase function names are valid, they are not idiomatic in Go.

* Within the same package function names must be unique.

* One of Go's features is that functions and methods can return multiple values.

* Go doesn't support function overloading.

* main() and int() are predefined function names.

* Syntax:

`func (receiver) name(parametes) (returns) {`

`        // code -> function body here`

`}`

* In Go a function is defined using `func` keyword.

* A function performs its tasks based on some input values. Those input values are called functions arguments.

`func f2(a, b int) {`
`    fmt.Println("Sum: ", a+b)`
`}`
   * The function f2 takes two parameter 'a' and 'b' which is only available in inside the function f2()

   * Parameter is the reserved space for arguments or variables or values passed into the function.

   * If you pass 2 and 3 into the f2 function, those 2 will be stored in 'a' and 3 in 'b'.

   * The space reserved to store the variables is called parameters (a & b).

   * The passed variables or values to a function is called arguments(2 & 3).

   * Calling function f2(a, b int) from main(). ( a & b are the parameters)

   `f2(2, 3) // The passed values are called arguments (2 & 3)`

### Shorthand Parameters Notation

* **When you have multiple consecutive parameters of the same type, you may omit the type name for the parameters up to the final parameter that declares that type. This is called Shorthand Parameters Notation.**

`func f3(a, b, c int, d, e float64, s string) {`
`    fmt.Println(a, b, c, d, e, s)`
`}`

### Return Statement

* A function can also return one or more values which can be printed or assigned  to other variables known as return values.

* When a function returns explicitly using a return statement, you must specify the data type of the return value just after the function parameter parentheses. However, in a function, the return statement is optional.

* Any statements below the return statement will never be executed.

`func f4(f float64) int {`
`    return int(f)`
`}`

* There should be a variable to receive the value returned from the function.

`num := f4(5.5)`
`fmt.Println("num : ", num)`

### Named Return Values

* named return values are a great way to explicitly mention the return variables in the function definition itself.

**Naked Return**

* A return statement without arguments returns the named return values. This is known as a "naked" return.

* Naked return statements should be used only in short functions. They harm readability in longer functions.