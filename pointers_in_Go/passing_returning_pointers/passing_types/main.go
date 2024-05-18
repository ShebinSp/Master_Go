package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile phone"
	sold = false
}
func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300 // equivalent to p.price = 300
	p.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}
func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changeValues():", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	// invoking function had not effect on our variables
	// The function worked and modified copies, not originals
	fmt.Println("After calling changeValues():", quantity, price, name, sold)
	fmt.Println()

	// What should we do to change the originals?: we pass pointers to the function
	fmt.Println("Before calling changeValuesByPointer():", quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling changeValuesByPointer():", quantity, price, name, sold)

	// What happens when passed a struct value to the function
	gift := Product{
		price:       100,
		productName: "Watch",
	}
	changeProduct(gift) // no changes in originals
	fmt.Println("gift: ", gift)
	fmt.Println("Before calling changeProductByPointer: ", gift)
	changeProductByPointer(&gift) // original changed
	fmt.Println("After calling changeProductByPointer: ", gift)

	// Things work differently with slices and maps. They don't store the actual data, but a reference to another
	// memory address where the data is stored. This means that even though they are passed by value so copies
	// are created, the copies points to the same data as the originals.
	// When a function changes a slice or map, the actual data is changed.
	prices := []int{1,2,3}
	// The slice already contains a pointer to its backing array, so we don't need to use a pointer to modify the
	// slice inside the function.
	changeSlice(prices)
	fmt.Println("Prices slice after calling the changeSlice(): ",prices)

	myMap := map[string]int{
		"a":1,"b":2,
	}
	changeMap(myMap)
	fmt.Println("myMap after calling the changeMap(): ",myMap)

	// Arrays behave just like ints, floats or strings.
	// Which means the functions don't modify the arrays that are passed.
	// However it is not good practice to pass array to a function. Pass slices insted.

}
