### Overflows

**In Go, overflow occurs when a numeric value exceeds the maximum or falls below the minimum representable value for its data type. Go has defined behaviors for overflow depending on the type of integer being manipulated.**

* In Go, when you increment a variable of type and it reaches its maximum value, an overflow occurs. This means the value wraps around to 0 instead of throwing an error or raising an exception.

`var x uint8 = 255
    x++ // overflow, value of x set to 0
    // Go resets the value to minimum value which is 0 for uint.
    fmt.Println("Value of x: ", x)`


* When you are attempting to assign a value that is outside the valid range for that type. As a result, the Go compiler will raise an error.

`var y uint8 = 256 // ERROR: range of uint is 0 - 255
    fmt.Println("Value of y: ",y)`


* Runtime overflow

`var z int8 = 127
    fmt.Println("Value of z: ", z + 1) // z will be set to -127 on runtime`

* A float number overflows to infinite(+Inf).


* Constant will not overflow. Because the value of a constant will not change. If you declare a typed constant with a value larger than type maximum value, Go will catch that error in compile time.

### Underflow
**If the value of an integer falls under the minimum value then underflow happens, and value of the variable sets to it's maximum value.**


* **if your application needs to work with very, very big numbers, that could overflow, you should use a package called 'big' that implements arbitrary precision arithmetic for big numbers.**