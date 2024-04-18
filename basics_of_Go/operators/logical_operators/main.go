package main

import "fmt"

func main() {
	a, b := 5, 10

// && (conditional and)
	// The evaluated expression is TRUE only if both sub expressions are TRUE.
	// The condition is TRUE because '5 is greater than 1 and 10 is equal to 10
	fmt.Printf("The Logical condition '%v > 1 && %v == 10' is : %v\n",a, b, a > 1 && b == 10)

	// The condition is FALSE because '5 is not less than 1 so the first sub expression is-
	// evaluated to FALSE. So the second sub expression will not be evaluated evaluated at all. 
	fmt.Printf("The Logical condition '%v < 1 && %v == 10' is : %v\n",a, b, a < 1 && b == 10)

// || (conditional or)
	// If any of the expression is evaluated to TRUE, then the entire expression is evaluated to true.
	
	// first expression is TRUE
	fmt.Printf("The Logical condition '%v == 5 || %v == 100' is : %v\n",a, b, a == 5 || b == 100) // TRUE
	// both expressions are FALSE
	fmt.Printf("The Logical condition '%v != 5 || %v == 100' is : %v\n",a, b, a != 5 || b == 100) // FALSE 

// ! (not or logical negation)
	// !(not) operator evaluates the expression and assigns as not.
	// below 5 is equal to 5 and the ! opreator evaluates it as FALSE( 5 == 5 -> TRUE and (!)not TRUE equals FALSE).
	fmt.Printf("The Logical condition '!%v == 5' is: %v\n",a, !(a == 5))

	// here 5 == 1(FALSE) and !false(means TRUE).So 1 condition is TRUE so the statement is TRUE
	fmt.Printf("The Logical condition '!%v == 1 || %v == 100' is: %v\n",a,b, !(a == 1 || b == 100))
}
