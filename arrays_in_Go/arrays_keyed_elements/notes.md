### Arrays with Keyed Elements (advanced feature of Go)

* *In composite literals, like arrays and slices, the index can be optionally provided* and the following rules apply:

* Each element has an associated integer index, making its position in the array or slice; the index starts from 0, not 1.

* An element with a key uses the key as its index; the key must be a non negative constant represented by a value of type int.

* An element without a key uses the previous elements index +1. If the first element has no key then its index is 0.

* The keyed element can in be any order.

* You can use this as a more compact way to initialise arrays or slices if the area of the slice has many zero values and just few non-zero values.

* **One advantage of keyed elements arrays is that you can skip contiguous parts when enumerating elements and the skipped elements will be initialised with zero values. You specify the first couple of elements and still specify the length you want the array or slice to have.**