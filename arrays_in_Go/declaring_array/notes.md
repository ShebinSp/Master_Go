### Arrays in Go

* An array is a composite, indexable type that stores a collection of elements of same type.

* An array has a fixed length.

* Every element is an array (or slice) must be of same type.

* Go stores the elements of the array in contiguous memory locations and this way it's very efficient.

* The length and the elements type determines the type of an array. The length belongs to array type and it's determined at compile time.

   accounts := [3]int{50,60,70}

   The array called accounts that consists of 3 integers has it's type [3]int