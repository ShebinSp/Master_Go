package main

import (
	"fmt"
	"strings"
)

func main() {
	// PACKAGE strings
	pr := fmt.Println // we can use `pr` insted of fmt.Println
	// To check a string contains in a string
	str := "I love Go programming"
	find := "love"
	result := strings.Contains(str, find) // returns true whether any Unicode points are within the string and false otherwise
	pr(result)                            // true

	// strings.ContainsAny()
	result = strings.ContainsAny(str, "io")
	pr(result)

	// reports whether a Unicode  code point is within a string
	result = strings.ContainsRune(str, 'g') // runes are represented in ''
	pr(result)

	// count() returns the number of instances of a substring
	n := strings.Count(str, "m")
	pr(n)
	// if the substring is an empty string, count returns 1 + the number of runes in the string
	n = strings.Count("three","")
	pr(n) // 5 + 1 = 6

	// To UPPERCASE or to lowercase
	// strings are immutable so the returned string will be a new one.
	
	// to lowercase
	s := "GO PythoN Java"
	pr(strings.ToLower(s))
	// to uppercase
	pr(strings.ToUpper(s))


	// How to efficiently compare strings in Go
	pr("go" == "go") // true
	// letter case matters here
	pr("Go" == "go") // false
	
	// a standard solution is create a lowercase version of each string and compare them but not efficient
	pr(strings.ToLower("Go") == strings.ToLower("go")) // traverses each string entirely before the comparison

	// strings.EqualFold() is used for case insensitivity.
	pr(strings.EqualFold("Go", "go")) // recommended way to compare strings when we don't want to take into account the latter case.
}
