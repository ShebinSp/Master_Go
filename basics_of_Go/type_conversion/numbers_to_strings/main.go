package main

import (
	"fmt"
	"strconv"
)

func main() {

// convert numbers to string
	// If convert int to a string using the string function, it doesn't return the string representation for that number-
	// but the Unicode code point for that integer.
	s := string(97)
	fmt.Println("value of s: ", s)

	// we can't convert a float to string. To do so,
	var myStr = fmt.Sprintf("%f",44.2)
	fmt.Println("value of myStr: ",myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println("value of myStr1: ",myStr1)
	fmt.Println("value of myStr1: ",string(34234))
	
// Convert string to integer
	// Package Strconv implements conversions to and from string representations of basic datatypes which are int, float & bool
	// string to float
	s1 := "3.33" // type string
	fmt.Printf("Type of s1: %T\n",s1)

	// convert a string to float
	f1, err := strconv.ParseFloat(s1, 64) // it will return 64bit float & an error if occured
	// let us mute the error
	_ = err
	fmt.Printf("value of f1: %f and Type of f1: %T\n",f1,f1)

	// strconv.Atoi() -> Ascii to int or string to int
	i, _ := strconv.Atoi("11")
	// Strconv.Itoa() -> Int to Ascii
	s2 := strconv.Itoa(21)
	fmt.Printf("value of i: %d and Type of i: %T\n",i,i)
	fmt.Printf("value of s2: %s and Type of s2: %T\n",s2,s2)

}
