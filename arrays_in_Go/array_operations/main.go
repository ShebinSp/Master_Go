package main

import (
	"fmt"
)

func main() {
	// Array Operations

	// Remember that we cannot add or remove elements from the array because of its fixed length, but can be modified
	// Modify an array
	fmt.Println("Modify an array")
	numbers := [5]int{1, 222, 3, 4, 555}
	fmt.Printf("Values of array - numbers: %v\n", numbers)
	numbers[0] = 111 // here i accessed the first index of array numbers and modified the value to 111.
	numbers[2],numbers[3] = 333, 444
	fmt.Printf("Values of array - numbers after: %v\n\n", numbers) // 1 changed to 111
	// you can not access an index out of an array
	// numbers[5] = 1 // error- array numbers has only 5 index which are 0,1,2,3,4.

	// iterate over an array
	fmt.Println("iterate over an array using the `range` keyword")
	// using the `range` keyword
	for index, value := range numbers { // range is a keyword not a function
		fmt.Println("index: ",index, "value: ",value)
	}

	// usnig for loop
	fmt.Println("using for loop")
	n := len(numbers) // it is a good practice, you can save some coputing power.
	for i := 0; i < n /*or len(numbers)*/; i++ { // you dont have to call len() funtion over each iteration.
		// print index from one add 1 to the index: index + 1. but remember index of an array starts from 0.
		fmt.Println("index: ",i + 1, "value: ",numbers[i]) // here we are accessing indexes in `i`th position. If you want to
	}

	// Multi dimention array (array in an array)
	fmt.Printf("\n\nMulti dimention array\n")
	balances := [2][3]int{ // an array which has 2 rows and 3 columns like a matrix
		{5,6,7},
		{1,2,3}, // dont forget the coma, it is mandatory when creating array on multiple lines.
	}
	fmt.Println("Multi dimentional array balances- ",balances)
	fmt.Println("for detailed index access")
	for i := 0; i < 2; i++ { // outer loop is for iterating over rows.
		for j := 0; j < 3; j++ { // inner loop is for iterating over columns
			fmt.Printf("value in row %v column %v: %v",i,j,balances[i][j])
		}
	}
	fmt.Println()

	// Arrays Equality
	fmt.Println("Arrays Equality")
	// 2 arrays are equal if they have the same length and the same elements in the same order
	ar := [3]int{1,2,3}
	ar1 := ar
	fmt.Println("ar is equal to ar1: ",ar == ar1) // true
	// then lets modify an element in ar
	ar[0] = 3
	fmt.Println("ar is equal to ar1: ",ar == ar1) // false

}
