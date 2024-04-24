## Slices Appending and Copying

* **To add an element to a slice, use the append() function.**

### append()

* append() function is a builtin function in Go and it returns a new slice after it appends a value to its end. Append() doesn't modify the initial slice; it returns a brand new one.

* To append a slice to a slice, use ellipsis operator.

### copy()

* **copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.** If the slices don't have the same number of elements it copies, the minimum of length of the slices.

* The copy() function doesn't always clone or duplicate a slice. If the destination slice has zero length nothing will be copied.