package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func init() { // the init() function will be invoked before the main() function.
	time := time.Now().Format("Jan 2 15:04:05 2006 MST")
	fmt.Println("The time is: ", time)
	fmt.Println()
}
func main() {
	// Switch statement

	// Example 1
	language := "Go"

	switch language {
	case "Python":
		fmt.Println("You are learning Python")

	case "Go", "golang":
		fmt.Println("Set, ready to Go")
	default:
		fmt.Println("C may be good for a start")
	}

	// Example 2
	n := 0                                               // try command-> go run main.go 3
	if args := os.Args; len(args) > 1 && len(args) < 3 { // we are taking a command line argument
		if av, err := strconv.Atoi(args[1]); err == nil {
			n = av
		} else {
			fmt.Println("Please enter an integer, Error:", err)
		}
	}

	switch n {
	case 0: // if there is any error in type conversion this statement will work as `break` statement here.
	case 1:
		fmt.Println("You have entered ", n)
	case 2:
		fmt.Println("You have entered ", n)
	case 3:
		fmt.Println("You have entered ", n)
	case 4:
		fmt.Println("You have entered ", n)
	case 5:
		fmt.Println("You have entered ", n)
	default: // not mandatory
		fmt.Println("Please a number between 1 and 5 ")
	}

	// Example 3
	n = 5
	switch true {
	case n%2 == 0:
		fmt.Printf("%d is even number\n", n)
	case n%2 != 0:
		fmt.Printf("%d is odd number\n", n)
	}

	// Example 4
	// A missing expression means true
	hour := time.Now().Hour() // assigns the current hour to variable `hour`

	switch { // missing expression means boolean `true`
	case hour < 12:
		fmt.Println("Its AM")
	case hour >= 12:
		fmt.Println("Its PM")
	}

}
