## Slice Expressions

* Arrays, slices and strings are sliceable. Slicing doesn't modify the array or slice, it returns a new one.

* A slice is formed by specifying a start or a low bound and a stop or high bound like `a[start : stop]`.

* This selects a range of elements that includes the element at index start, but excludes the element at index stop.

* Slicing an array returns a slice, not an array.

* For convenience any of the indexes may be omitted, a missing low index defaults to 0 and a missing high index defaults to the length of the slice operand.