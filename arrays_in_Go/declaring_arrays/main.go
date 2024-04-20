package main

import "fmt"

func main() {
// Different ways to declare an array
	// declaring an array (the initial value of elements are zero)
	fmt.Println("Declare arrays")

	var numbers [4]int // this array declares with 4 elements which has 0 values.
	fmt.Println("Elements in array - numbers: ", numbers)
	// for all value with details of an array
	fmt.Printf("%#v \n",numbers)

	accounts := [3]int{} // using short variable declaration.
	fmt.Println("Elements in array -accounts: ", accounts)
	fmt.Printf("%#v \n",accounts)

	var a1 = [4]float64{} // this array is initialized with type float64, by default with 4 elements which has it's value 0.
	fmt.Printf("type and values of array - a1: %#v\n",a1)
	fmt.Println()

	// declare an array and initialize values
	fmt.Println("declare an array and initialize values")

	var a2 = [5]int{1,2,3-4,5} // integer array
	fmt.Printf("type and values of array - a2: %#v\n",a2)

	var a3 = [3]bool{true, false,true} // boolean array
	fmt.Printf("type and values of array - a3: %#v\n",a3)

	a4 := [4]string{"John", "Doe","John Doe", "Doe John"} // string array
	fmt.Printf("type and values of array - a4: %#v\n",a4)

	// It is not mandatory to initialize all the values in an array
	a5 := [5]int{1,2} // only 2 values initialized and we can initialize 3 values later
	fmt.Printf("type and values of array - a5: %#v\n",a5)
	// we can access the index of an array by index number which starts from zero
	a5[2] = 3
	a5[3] = 4
	a5[4] = 5
	fmt.Printf("type and values of array - a5: %#v\n\n",a5)

	fmt.Println("ellipsis operator")
	// ellipsis operator (...)
	// an ellipsis operator finds out the length of an array automatically, but here we cannot initialize values later.
	a6 := [...]string{"a","b","c","d","e"} // you can initialize as many as elements you want
	fmt.Printf("The length of array - a6: %d\n",len(a6))

	// sometimes we want to declare the elements of an array on a multiple lines for better readability
	a7 := [...]int{ // i am using ellipsis operator because i dont know how many elements will be in the array at this moment.
		1,
		2,
		3,
		4,
		5,
		6,
		7, // a coma in the last line is mandatory otherwise compile time error will occur.
	}
	fmt.Printf("The length of a7: %d |  values:  %v\n",len(a7),a7)

}
