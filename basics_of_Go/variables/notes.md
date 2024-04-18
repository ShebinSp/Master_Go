Variables in Go
-----------------

* A variable is a memory location for holding a value.

* In Go a variable belongs and it's created at runtime.

* A declared variable must be used in Go.

* _ is the blank identified and mute the compile error returned by unused variables.

------------------------------------------

 Declaring variables:

* Using var keyword

    var x int = 7

    var x int = 7 ( it is inferred declaration, the data type should be omitted here ).

    var s1 string  // Declared to use later

    s1 = "Learning Go!"  // Assigning value to s1 later

    _ = s1 // s1 is blank identified to mute the compile error returned by unused variables.

* Using the Short Variable Declaration Operator ( := )

    age := 30 // Here go automatically understand the type of var age is an int by the value.

    name := "John Doe" // type will be inferred by Go runtime.

    
Multiple Declaration
-------------------

* using Short Variable Declaration Operator

    car, cost := "Audi", 60000 // type string & int declared

* using var keyword ( for different data types )

        var (

                salary float32

                firstName string

                permanent bool

            )

* using var keyword ( for same type )

    var num1, num2, num3 int

* When to use normal declaration( var ) and Short Declaration Operator( := )

    When we know the initial value then Short Declaration Operator can be used or var keyword otherwise.

    We use normal declaration for multiple declaration for better readability.


Multiple assignment:
------------------

* Tuple Assignment

    var i, j int

    i, j = 5, 7 ( tuple assignment, it allows multiples values assigned at once )

    // to swap the values

    i , j = j, i


* we can use expressions in Short Declaration Operator

    sum := 5 + 2.5

    fmt.Println("Sum = ", sum)