package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// HOW TO GET INPUT FROM COMMAND LINE INTERFACE
	scanner := bufio.NewScanner(os.Stdin) // os.Stdin reads the output from the command line
	fmt.Printf("%T\n", scanner)
	scanner.Scan()           // expects to get input from command line
	text := scanner.Text()   // input from command line stored to `text` as string
	bytes := scanner.Bytes() // input from command line stored to `bytes` as bytes

	fmt.Println("Input text: ", text)
	fmt.Println("Input bytes: ", bytes)
	// To read from a file use `<` symbol in command line when running the program
	// go run main.go < main.go // it will print the first line in main.go
	//---------------------------------------------------------------------------------

	// hOW TO READ THE INPUT CONTINUESLY OR UNTIL A SPECIFIC STRING IS SCANNED

	// until the user types 'exit'
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered: ", text)
		if text == "exit" {
			fmt.Println("Exiting the scanning...")
			break
		}
	}
	fmt.Println("Just a message after the for loop")
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

}
