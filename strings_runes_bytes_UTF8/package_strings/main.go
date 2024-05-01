package main

import (
	"fmt"
	"strings"
)

func main() {
	// PACKAGE strings (section 1)

	p := fmt.Println // we can use `p` insted of fmt.Println
	// To check a string contains in a string
	str := "I love Go programming"
	find := "love"
	result := strings.Contains(str, find) // returns true whether any Unicode points are within the string and false otherwise
	p(result)                             // true

	// strings.ContainsAny()
	result = strings.ContainsAny(str, "io")
	p(result)

	// reports whether a Unicode  code point is within a string
	result = strings.ContainsRune(str, 'g') // runes are represented in ''
	p(result)

	// count() returns the number of instances of a substring
	n := strings.Count(str, "m")
	p(n)
	// if the substring is an empty string, count returns 1 + the number of runes in the string
	n = strings.Count("three", "")
	p(n) // 5 + 1 = 6

	// To UPPERCASE or to lowercase
	// strings are immutable so the returned string will be a new one.

	// to lowercase
	s := "GO PythoN Java"
	p(strings.ToLower(s))
	// to uppercase
	p(strings.ToUpper(s))

	// How to efficiently compare strings in Go
	p("go" == "go") // true
	// letter case matters here
	p("Go" == "go") // false

	// a standard solution is create a lowercase version of each string and compare them but not efficient
	p(strings.ToLower("Go") == strings.ToLower("go")) // traverses each string entirely before the comparison

	// strings.EqualFold() is used for case insensitivity.
	p(strings.EqualFold("Go", "go")) // recommended way to compare strings when we don't want to take into account the latter case.

	// SECTION 2

	// strings.Repeat() returns 'n' copies of the string that passed as first argument
	myStr := strings.Repeat("ab", 11) // ab repeates 11 times
	p(myStr)

	// strings.Replace() returns a copy of the string 's' with the first 'n' non-overlapping instances of old replaced by new
	myStr = strings.Replace("192.168.0.1", ".", ":", 2) // first 2 colons replaced by .
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1) //-1 sets all colons replaced by .
	p(myStr)

	// strings.ReplaceAll() returns a copy of the string s with all non-overlapping instances of old replaced by new.
	myStr = strings.ReplaceAll(myStr, ":", ",") // all colons replaced by ,
	p(myStr)

	// strings.Splite()

	ss := strings.Split("a,b,c", ",") // ["a", "b", "c"]
	fmt.Printf("Type of ss: %T\n", ss)
	fmt.Printf("Value of ss: %#v\n", ss) // each character is splited into
	fmt.Println(ss[1])

	ss = strings.Split("Go for Go", "")  // [G o   f o r   G o]
	fmt.Printf("Value of ss: %#v\n", ss) // each character is splited into []string{"G", "o", " ", "f", "o", "r", " ", "G", "o"}

	// strings.Join()
	ss = []string{"I", "learn", "Go"}
	myStr = strings.Join(ss, "--") // joins ss with --
	p(myStr)

	// strings.Fields()
	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr) // returns slice of string
	fmt.Printf("Type of fields: %T\n", fields)
	p(fields)
	p("element in index 1 is: ", fields[1])

	// remove trailing white spaces and tabs
	s1 := strings.TrimSpace("\t Goodbye windows, welcome Linux!\n")
	fmt.Printf("%q\n", s1)

	// To remove other leading and trailing characters
	s2 := strings.Trim("...Hello, @Gophers!??",".,@?")
	fmt.Printf("%q\n", s2)
	

}
