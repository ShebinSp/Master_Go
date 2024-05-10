package main

import (
	"fmt"
	"strings"
)


func f1(a...int){
	fmt.Printf("f1(): %T\n",a)
	fmt.Printf("f1(): %#v\n",a)
}

// A variadic function that modifies one of the arguments that are passed
func f2(a ...int){
	a[0] = 30
}

// A variadic function to return sum and product of float arguments passed
func sumAndProduct(a ...float64)(float64, float64){
	sum := 0.
	product := 1.
	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

// Mix variadic and non variadic parameters
// put the non variadic parameter before the variadic parameter
func personInfo(age int, names ...string)string{
	fullName := strings.Join(names, " ")
	pInfo := fmt.Sprintf("Age: %d, Full Name: %s",age,fullName)
	return pInfo
}



func main() {
	// VARIADIC FUNCTION

	f1(1,2,3,4,5)
	// variadic means zero or more. So f1() with no agruments
	f1()

	// Another example of a well known variadic function is append built-in function.
	// It appends items to an existing slice and returns back the same slice.
	nums := []int{1,2}
	// it appends 3,4,5 to nums and returns it back.
	// It is a variadic function, which can call it with a variable number of arguments.
	nums = append(nums, 3,4,5)

	// pass a slice to a variadic function by post fixing it with the Variadic operator or redux.
	// In this case, the slice is directly passed to the function without a new slice being created.
	f1(nums...) // slice passed into variadic function f1()

	f2(nums...)
	fmt.Println("nums: ",nums)

	s, p := sumAndProduct(2.0, 5., 10.)
	fmt.Println("s , P: ",s,p)

	info := personInfo(30, "wolfgang","Amadeus", "Mozart")
	fmt.Println("Person Informations: ",info)
}