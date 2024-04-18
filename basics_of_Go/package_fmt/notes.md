Functions In fmt
---------------

* Package fmt implements formatted I/O with functions analogous to C's printf and scanf.

* It's used mainly to print out to stdout

* fmt.Println() writes to standard output.

* fmt.Printf() prints out to stdout according to a format specifier called verb.

* fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()

* s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)

* fmt.Println(s) // => a is 10, b is 15.500000, c is Gophers



**The format specifiers(verbs) in fmt.Printf()**

* %d - int, %+d - (%+5) = +5

* %f - float, %.3f (it means only show 3 decimal points, change the number to adjust the decimals ).

* %s - string

* %q - quoted string

* %v - for all type of values

* %#v - a Go-syntax representation of the value

* %t - for boolean expressions

* %T - To print out the type of a variable

* %p - pointer (address in base 16, with leading 0x)

* %b - To print the Base -2 (binary) of an integer

* %8b - To print out in 8 bits ( 1 byte for IP addresses )

* %x - Convert a decimal number into Hexadecimal

* To quote a line in Println() or Printf(): ("he said "\hello\"")