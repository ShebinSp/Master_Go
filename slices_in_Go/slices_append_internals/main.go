package main

import "fmt"

func main() {
	// The Internals of append() function

	// lets start a slice which is == nil
	var nums []int
	fmt.Printf("nums: %#v\n", nums)
	fmt.Println("Length of nums: ",len(nums), "  |  Capacity of nums: ",cap(nums)) // Length : 0   |  Capacity : 0

	// using append() function
	nums = append(nums, 1,2)
	fmt.Println("Length of nums: ",len(nums), "  |  Capacity of nums: ",cap(nums)) // Length :  2   |  Capacity :  2

	// The capacity of the new backing array is now larger than the length to avoid creating a new backing array when the-
	// next append() is called.
	nums = append(nums, 3)
	fmt.Println("Length of nums: ",len(nums), "  |  Capacity of nums: ",cap(nums)) // Length :  3   |  Capacity :4  (2+2)

	nums = append(nums, 4)
	fmt.Println("Length of nums: ",len(nums), "  |  Capacity of nums: ",cap(nums)) // Length :  4   |  Capacity :  4
	
	// Creating a new backing array each time an append() function is called not efficient at all. So each time-
	// an append() function is called the the capacity will increase 2,4,8,16,32,64 and so on.
	nums = append(nums, 55)
	fmt.Println("Length of nums: ",len(nums), "  |  Capacity of nums: ",cap(nums)) // Length :  5   |  Capacity : 8 (4+4)

	// this line of code overwrite the last element of the slice and add 5 other element
	nums = append(nums[0:4], 100,200,300,400,500)
	fmt.Println("Length of nums: ",len(nums), "  |  Capacity of nums: ",cap(nums)) // Length :  9   |  Capacity : 16 (8+8)
	fmt.Println("nums: ",nums)

	// lets check a more advanced example
	letters := []string{"A","B","C","D","E"}
	letters = append(letters[0:1], "X","Y")
	fmt.Println("Length of letters: ",len(nums), "  |  Capacity of letters: ",cap(nums)) // Length :  9   |  Capacity : 16 (8+8)
	fmt.Println("letters: ",letters)

	// we cannot access a slice after its length.
	// fmt.Println(letters[4]) // runtime error: trying to access 4th index of slice -> index out of range

	// but we can access using a slice expression because slices see the whole backing array.
	fmt.Println("letters: ",letters[3:5])

}
