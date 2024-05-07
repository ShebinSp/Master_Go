package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// WORKING WITH FILES IN GO

	// The common way to work with files in Go is using `package os`
	// Other Go packages used are `io` and `bufio`

	// How to create a new file

	// created a new file which is pointer to os.File
	var newFile *os.File
	fmt.Printf("Type of newFile := %T\n", newFile)

	// create a variable to handle the error which may occure, Go doesn't have exceptions like in other languages.
	var err error

	// Create() function creates a file if it doesn't already exist. If it exists, the file is truncated.
	// This function returns a file descriptor, which is in fact a pointer to os.File and an error value.
	newFile, err = os.Create("a.txt") // argument is the name of the file with extension.
	if err != nil {
		// fmt.Println("Error on os.Create(): ", err)
		// os.Exit(1) // os.Exit(1) exist the program without running defferred functions()
		log.Fatal("Error: ", err) // idiomatic way to handle errors in Go
	}

	// os.Truncate() is used to truncate or empty a file.
	err = os.Truncate("a.txt", 0) // the second argument 0 will completely empty the file. If 50 used insted of 0-
	// the file will be truncate to 50 bytes. If the file is over 50 bytes, everything past 50 bytes will be lost.
	// Input some text in a.txt to check.
	if err != nil {
		log.Fatal("Error: ", err)
	}
	// When we done with the file we must close it
	newFile.Close() // idiomatic way - defer newFile.Close()


	// How to OPEN and CLOSE an existing file
	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	file.Close() // defer file.Close()

	// Open a file using `os.OpenFile()`
	file, err = os.OpenFile("a.txt", os.O_CREATE | os.O_APPEND, 0644) // | is or
	if err != nil {
		log.Fatal("Error: ", err)
	}
	file.Close() // defer file.Close()


	// How to get File Information
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt") // returns value of fs.FileInfo and an error value
	p := fmt.Println // for lesser code
	p("File Name: ",fileInfo.Name()) // name is a method of FileInfo
	p("File size in bytes: ",fileInfo.Size())
	p("Last modified: ",fileInfo.ModTime()) // file last modified time
	p("Is Directory? : ",fileInfo.IsDir())
	p("Permissions: ",fileInfo.Mode())


	// How to Check a File Exist or Not
	// os.Stat() returns file info and an error if file doesn't exist
	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err){
			// log.Fatal("The file does not exist")
			log.Println("The file does not exist")
		}
	}


	// How to Rename of Move a file
	oldPath := "a.txt" // use a valid path that can be specific to the operating system
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	


	// How to Remove a File
	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
