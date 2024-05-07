package main

import (
	"log"
	"os"
)

func main() {
	// HOW TO WRITE A BYTE SLICE TO A FILE

	// os.OpenFile() open the file in write only mode if file exists, and then it trucates the fie.
	// If the file doesn't exist, it creates the file with 0655 permissions.
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY | os.O_TRUNC | os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// byte slice is just like string but mutable, means you can modify each byte or rune.
	byteSlice := []byte("I learn Go")

	byteWritten, err := file.Write(byteSlice) // returns the number of bytes written and an error value
	if err != nil {
		log.Println(err)
	}
	log.Printf("Bytes Written: %d\n",byteWritten)

	// Another way to write to a file is to use `io.util` package, which has a function called `WriteFile()`
	// This function will handle creating, opening, writing a slice of bytes and closing the file. It's very useful
	// if you just need a quick way to dump a slice of bytes to a file.
	// ioutil.WriteFile() is deprecated. Use os.WriteFile() instead.
	bs := []byte("Go Programing is cool")
	err = os.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Println(err)
	}

}