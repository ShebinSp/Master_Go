# String Slicing

* Unicode characters occupy between 1 and 4 bytes, ASCII letters occupy 1 byte and other unicode code points use more bytes in Go.

* A slice expression returns a part of a slice and both slices of the original and the returned one share the same backing array.

* Slicing a string is very efficient because it uses the same backing array.

* To return a rune, convert a slice of bytes to slice of runes and to print convert the rune to string.

* Converting between different types, like a string and a rune slice is not efficient. Because a backing(underlying) array is created each time.