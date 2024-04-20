package main

import "fmt"

func main() {
	// Declaring a slice
	var cities []string
	// lets check the values inside the slice cities
	fmt.Println("slice cities is equal to nil: ", cities == nil) // true
	// A nil slice is an empty slice with zero capacity
	fmt.Printf("slice cities: %#v\n", cities)
	fmt.Println("length of slice cities: ", len(cities))

	// we cannot assign a element to a nil slice
	// cities[0] = "London" // Error: index out of range

	// we can initialise elements when declaring a slice
	numbers := []int{1,2,3,4,5}
	fmt.Println("slice numbers: ",numbers)

	// creating a slice using make() function.
	// first argument of make() is the type and second argument is the length.
	nums := make([]int, 2)
	fmt.Printf("slice nums: %#v\n",nums) // a slice created with length 2 -> slice nums: []int{0, 0}


	// lets create a new type called names which is a slice of string
	type names []string
	friends := names{"Shelvin", "Jibin"} // a variable called friends of type names which is in fact a slice of strings
	fmt.Println("slice friends: ",friends)

	// the indexing is used to access a specific element from a slice
	fmt.Println("my best friend is: ",friends[0])
	myFriend := friends[1]
	fmt.Println("My best friend is : ",myFriend)

	// to modify a slice
	friends[0] = "Linto"
	fmt.Println("slice friends: ",friends)
	fmt.Println()

	// iteration over a slice is same as an array
	for index, value := range numbers {
		fmt.Printf("index: %v  -  value: %v\n",index, value)
	}
	fmt.Println()

	// Slices with same element type can be assigned to each other, in an array the length of the array also should be same.
	var n []int
	n = numbers
	fmt.Println("slice n: ",n)
}
