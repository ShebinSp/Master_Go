package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// STRING LENGTH IN BYTES AND RUNES

	// Unicode characters occupy between 1 and 4 bytes, ASCII letters occupy 1 byte
	//  and other unicode code points use more bytes in Go.
	s1 := "Go"                          // a string used only ASCII
	fmt.Println("len of s1: ", len(s1)) // 2 each ascii character used 1 byte

	word := "देश"
	fmt.Println("len of word: ", len(word)) // 9, each character is represented using multiple bytes

	fmt.Println(len("ആ")) // 3 bytes

	// returns number of runes or unicode points, not bytes
	n := utf8.RuneCountInString(word) // 3 runes or characters in the string 'word'. 3 runes * 3 bytes = 9 (len(word))
	m := utf8.RuneCountInString("ആ")  // 1 rune & 3 bytes
	fmt.Println("n & m: ", n, m)

}
