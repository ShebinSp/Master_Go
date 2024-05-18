# Passing Types as Pointers
* Go allows to pass parameters to functions both by pointers and values.

* Sometimes when a pointer is passed to a function is called passed by reference, strictly speaking, this is not correct.

* There is no passed by reference in Go, but only passed by value. Each time you pass an argument to a function, even if it's a pointer, that argument will be copied to a different location.

* Slices and maps are not meant to be used with pointers. Because they don't store the actual data, but reference to another memory address where the data is stored. This means that even though they are passed by values so copies are created, the copies points to the same data as the originals. So when a function changes a slice or map, the actual data is changed.

* Arrays behave just like ints, floats or strings. Which means the functions don't modify the arrays that are passed. However it is not good practice to pass array to a function. Pass slices insted.

### When we should pass by value and When we should pass by pointer?

* The answer depends on variable type, and if we want that variable to be changed inside the function or not.

* **NOTE** that passing by value is cheaper that passing by pointer. Pointer also put pressure on garbage collector.

* Pass pointers to function only when it's really necessary.