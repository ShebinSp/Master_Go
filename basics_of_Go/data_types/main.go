package main

import "fmt"

func main() {
	// Int Type
	var i1 int8 = 100  // can contain -128 to 127
	fmt.Printf("Type  of i1: %T\n", i1)

	var i2 uint16 = 65535 // can contain 0 - 65535
	fmt.Printf("Type  of i2: %T\n", i2)


	// Float Type
	var f1, f2, f3 float64 = 1.1, -0.2, 5.0
	fmt.Printf("Types  of f1, f2, f3 : %T, %T, %T\n", f1, f2, f3 )

	// Rune Type
	var r rune = 'f'
	fmt.Printf("Type  of r: %T\n", r) // Type is `int32` so, rune is an alias for int32.
	fmt.Println("r(ASCII): ",r) // 102 is the ASCII code for `f`
	fmt.Printf("r(hexadecimal): %x\n",r) // 66 is the hexadecimal ASCII for `f`


	// Bool Type
	var b bool = true
	fmt.Printf("Type  of b: %T\n", b)
	bb := 1 < 0
	fmt.Printf("Type  of bb: %T\n", bb)


	// String Type
	var s string = "Hello Go!"
	fmt.Printf("Type of s: %T\n",s)

	// Array Type
	var numbers = [7] int{1,2,3,-4,50}
	fmt.Printf("Type of numbers: %T\n",numbers)

	// Slice Type
	var cities = [] string {"London", "Cochi","Bangalore"}
	fmt.Printf("Type of cities: %T\n",cities)

	// Map Type
	balances := map[string]float64{
		"USD" : 88.8,
		"EUR" : 100.1,
	}
	fmt.Printf("Type of balances: %T\n",balances)

	// Struct Type
	type Person struct {
		name string
		age uint8
	}
	var p1 Person
	p1.name = "John Doe"
	p1.age = 18

	p2 := Person{
		name: "Doe John",
		age: 19,
	}
	fmt.Println("Details of p1 & p2: ",p1, p2)
	fmt.Printf("Name of %q: %s and %q: %d\n","person 1",p1.name,"age", p1.age)
	fmt.Printf("Name of %q: %s and %q: %d\n","person 2",p2.name, "age",p2.age)
	fmt.Printf("Type of %q & %q: %T, %T\n","p1","p2",p1,p2)
	
	// Pointer Type
	var x int = 2
	ptr := &x
	fmt.Printf("Type of %q: %T with value(memory address(in hexadecimal)): %v\n","ptr", ptr,ptr)
	// To see the value a pointer points to use *
	fmt.Printf("%q points to the value: %v\n","ptr", *ptr)

	// Function Type
	fmt.Printf("Type of %q: %T\n","f",f)
	
}

// function f()
func f(){
		
}