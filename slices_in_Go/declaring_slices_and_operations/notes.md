### Slices in Go

*   ***A slice is a composite type that stores a collection of elements but unlike an array a slice can shrink or grow.***

* We can create a keyed slice like a keyed array.

**DIFFERENCE B/W ARRAY & SLICE**
* **Array**

  * Has a fixed length defined at compile time.

  * The length of an arrays is part of its type, defined at compile time and cannot be changed.

  * By default an uninitialised array has all elements equal to zero.

**Slice**

  * Has a dynamic length (it can shrink or grow).

  * The length of a slice is not part of its type and it belongs to runtime.

  * An uninitialised slice is equal to nil ( its zero value is nil ).

  * Both slice and an array can contain only the same type of elements.

  * Slices with same element type can be assigned to each other, in an array the length of the array also should be same.