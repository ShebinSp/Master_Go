package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b' // a & b are rune literals or characters are represented in int32 Unicode code points.
	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	str := "țară" // țară means country in romanian & and each rune occupies 1 and 4 bytes.
	// UTF8, which is the encoding scheme in Go is a variable length in coding scheme. It uses b/w 1 & 4 bytes for each one.
	// The built in function len() returns the number of bites, not tokens or characters.
	fmt.Println("rune str: ", len(str))

	// when accessing the index we get the byte at the position and not the character of rune.
	fmt.Println("Byte (not rune) at position 1: ", str[1])

	// if we want to see what's inside the string, byte by byte, we loop over the string
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // but it will show some strange characters
	}

	fmt.Printf("\nutf8.DecodeRuneInString()\n")
	// decode a string rune by rune manually
	// to get exactly as rune by rune we use uft8 package
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size // size of str is 6 bytes, and we increment the i by size of the rune in size
	}
	fmt.Println()

	// To decode a runes in a string automaticaly, we use a four and the range loop.
	for _, r := range str {
		fmt.Printf("%c", r)
	}
	fmt.Println()
}
