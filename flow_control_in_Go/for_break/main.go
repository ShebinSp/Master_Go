package main

import "fmt"

func main() {

	// break statement

	// Example one
	fmt.Println("Example 2")

	//  a program to find first 10 numbers which can divided by 13
	count := 0 // initialized variable count to count the numbers divicible by 13

	// its an infinite loop
	for i := 1; true; i++ { // here second statement acts as inifinite(always true, value of i will be kept increaseing in the loop)
		if i%13 == 0 { // if there is a number that can divicible by 13
			count++                                                               // increment the value of `count` by 1.
			fmt.Printf("%v is divisible by 13   |   %d * 13 = %d\n", i, count, i) // print that number
		}
		if count == 10 { // if the value of `count` reaches 10, then
			break // it breaks the current loop. The inner loop if there are more loops and continues execution-
			// at the statement following the end of the for loop.
		}
	}
	fmt.Println("for loop breaks and execute this statement.") // This is the statement that will be executed after-
	fmt.Println()                                              // the break statement.

	// Example 2
	fmt.Println("Example 2")

	// in this code we are iterating in a string to find a character.
	findMe := "ios"    // find the only first letters inside the string `findMe`,
	word := "continue" // from the string `str`.
	found := false

	for i := 0; i < len(findMe); i++ { // This is an outer loop, which has a length of the string `findMe` which is 2.
		for j := 0; j < len(word); j++ { // This is an inner loop, which has a length of the string `word` which is 2.
			if findMe[i] == word[j] { // Once we find the letter from the `findMe`, we break the inner loop and jumps to it's outer loop.
				fmt.Printf("%v is the %v letter in %s\n", string(findMe[i]), j+1, word)
				found = true // if we found a match for a letter in `findMe` in `word`, then set the value of boolean `found` to true.
				break
			}

		}
		// in the result you can see that, when the first letter "i" is found in "continues", the `break` statement breaks
		// the inner loop and exited the inner loop(j loop) and again the outer loop will start to iterate and again the-
		// control will flow to the inner loop.

		if !found { // if the value of found is false then printing the letter not in word.
			fmt.Printf("%s not found in %s\n", string(findMe[i]), word)
		}
		found = false // and again setting the value of `found` to false for next check.
	}
}
