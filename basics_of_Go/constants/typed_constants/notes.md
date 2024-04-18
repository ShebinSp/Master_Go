Constants in Go
---------------

* In Go, er use the term constant to represent fixed (unchanging) values.

* We use constants to avoid possible errors ( variables that change when they shouldn't ) or to replace a value only in one place instead of in many places.

* All basic literals (1,3.4,"hello", true) are in fact unnamed constants.

* A constant belongs to the compile time and it's created at compile time. It's value can not be changed while the program is running.

* Another advantage of using constants is that Go can not detect runtime errors at compile-time but constants belong to compile time so errors can be detected earlier.

* You can declare constants that store numbers, strings or booleans.

* Value should be assigned to a constant when it's declared.


**Constant rules**

* You can't change a constant.

    const num = 7

    num = 5 // ERROR

* You can't initilize a constant at runtime. Constant belongs to compile time.

    const power = math.Pow(2,3) // ERROR, functions belongs to runtime.

* You can't use a variable to initialize a constant. Because variables belongs to runtime & constants belongs to compile time.

    t := 5

    const tc = t // ERROR

* We can use the length of a string as constant value using len() function, and it should be a string literal because a string literal is an unnamed constant, not a string variable. Here we can initialize a constant using len() function, because len() is a built-in function which is available at compile time and math.Pow() is a runtime funtion.

    const strLen = len("hello")
    
    str := "hello"
    const strLen = len(str) // ERROR
