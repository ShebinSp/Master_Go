package main

import "fmt"

func main() {
	// declaring an array
	var numbers [4]int // this array declares with 4 elements which has 0 values.
	accounts := [3]int{50,60,70} // using short variable declaration with elements 50,60,70
	
	// numbers
	fmt.Println("Elements in array 'numbers': ", numbers)
	// for all value with details of an array
	fmt.Printf("%#v \n",numbers)
	// accounts
	fmt.Println("Elements in array 'accounts': ", accounts)
	fmt.Printf("%#v \n",accounts)
}
