package main

import "fmt"

func main() {

	// Example 1
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too Expensive")
	}

	// buy only if price is 100 or below 100 and in stock.
	if price <= 100 && inStock { // inStock means: inStock == true
		fmt.Println("In stock, Buy now!")
	}
	

	// Only one of the following branch will be executed.
	// The else branch will be executed only if other branches are false.
	if price < 99 { // parenthesis are no required to enclose the testing condition
		fmt.Println("Too cheap")
	} else if price == 100 {  // this statement will be executed.
		fmt.Println("Buy now")
	} else { //executed only once if all the if branches are false (it's optional)
		fmt.Println("It's expensive!")
	}

	// Example 2
	age := 55

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return in %d years!\n",18 - age)
	} else if age == 18 {
		fmt.Println("It's your first vote!")
	} else if age > 18 && age <= 100 {
		fmt.Println("You are elegible to vote")
	} else {
		fmt.Println("Invalid age!")
	}

}
