package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// How to read a file line by line or using any other delimiter
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// The file value returned by os.Open() can be wrapped in bufio Scanner just like a buffered reader.
	// Then we'll call the Scan() to read up to the next delimiter and then use text for string or bytes
	// for byte slice to get the data that was read.

	// The default scanner is bufio.ScanLines() means it will scan a file line by line. There are also ScanWords()
	// and ScanRunes(). We can impliment a custorm split function if we want.
	scanner := bufio.NewScanner(file)

	// To read word by word
//	scanner.Split(bufio.ScanWords) // this will read words by words
//	scanner.Split(bufio.ScanRunes) // this will read rune by rune, here character by character

	success := scanner.Scan()
	if success {
		err := scanner.Err()
		if err == nil {
			log.Println("Scan was complete and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}
	log.Println("First line found: ", scanner.Bytes()) // Printed as byte

	// to get all the lines we use a for loop
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // printed as slice
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
