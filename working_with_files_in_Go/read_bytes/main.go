package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// HOW TO READ BYTES
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 2) // length of byteSlice = 2
	// io.ReadFull() reads bytes exactly equal to the len(byteSlice)
	numberBytesRead, err := io.ReadFull(file, byteSlice) // only 2 bytes will read from the file
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	// Read the entire contents
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data: \n%s\n", data)
	log.Println("Size of data: ", len(data), " bytes")

	// Another way to read the entire file
	// io.ReadFile() is useful if you need a quick way to read the entire contents of the file into a byte slice.
	fileData, err := os.ReadFile("test.txt")
	log.Printf("Data read: \n%s\n", fileData)
	log.Println("Size of data: ", len(fileData), " bytes")

}
