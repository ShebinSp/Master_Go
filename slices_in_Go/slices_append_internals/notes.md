# How to work with Slices

* Slices are a key data type in Go

## The Internals of append() function

 * append() is not always a very efficient slice operation.

 * `var nums []int` here the capacity of the slice, which is a 0, is the length of the backing array and the arrays have a fixed length in Go.

 * `nums = append(nums, 1, 2)` -> How could i add 2 elements to the backing array?

    * The elements of the slice are stored in the backing array, not in the slice and the answer is that it creates a new backing array if the slice capacity is full.

    * The append() function creates a new backing array with a larger capacity to avoid creating a new backing array when the next append() is called.

     * `Length :  3   |  Capacity :  4`

    * Creating a new backing array each time an append() function is called not efficient at all. So each time an append() function is called the the capacity will increase 2,4,8,16,32,64 and so on. It avoids creating new backing arrays.
    
    * We cannot access a slice after its length. But we can access using a slice expression because slices see the whole backing array.