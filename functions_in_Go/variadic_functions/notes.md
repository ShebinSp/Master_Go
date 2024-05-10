# Variadic  Functions

* Variadic functions are functions that take a variable number of arguments.

* Ellipsis prefix or three dots in front of the parameter type makes a function variadic and the function may be called with zero or more arguments for that parameter. So **variadic means zero or more**.

* If the function takes parameters of different types, only the last parameter of a function can be variadic.

* fmt.Println() is an example of a variadic function.

* A variadic function increases readability. It's easier to understand the function that is written with only one parameter instead of 4 or 5 or even six possible parametes.

* Another example of a well known variadic function is append built-in function. It appends items to an existing slice and returns back the same slice.
   nums := []int{1,2}
   it appends 3,4 to nums and returns it back

   It is a variadic function, which can call it with a variable number of arguments.
   nums = append(nums, 3,4,5)

* You can pass a slice to a variadic function by post fixing it with the **Variadic operator or redux**.
   In this case, the slice is directly passed to the function without a new slice being created.

* ### Good Practice to use variadic functions

1. When the number of parameters is unknown.

2. When you don't want to create a temporary slice just to pass it to a function. This could be an alternative to a variadic function to pass in a slice which can have a variadic number of elements.