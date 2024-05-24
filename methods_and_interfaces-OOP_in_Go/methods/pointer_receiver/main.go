package main

import "fmt"

type car struct {
	brand string
	price int
}

// a normal function to change the car properties, pass by value mechanism
func changeCar(c car, newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// a method to change the car properties, pass by value mechanism
func (c car) changeCarM(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// a method with a pointer receiver to chage the car properties
func (c *car) changeCarMP(newBrand string, newPrice int) {
	// You don't have to use the dereferencing operator(*) here. Go do it automatically
	// (*c.price)= newPrice
	c.brand = newBrand
	c.price = newPrice
}

func main() {
	// METHODS POINTER RECEIVER

	// We know, calling a function with arguments makes a copy of each argument value.
	// If a function needs to update a variable passed as argument, or if an argument is so large that we wish to avoid-
	// copying it, we must pass the address of the variable, not the variable.
	// The same goes for methods that need to update the receiver.

	// Changeing the properties of a type
	myCar := car{"Audi", 50000}
	changeCar(myCar, "BMW", 55000)
	// The change of properties of car will reflect only inside the function changeCar()
	fmt.Println("My car: ", myCar) // expected `BMW 55000` but no changes made.

	// Lets try a method
	myCar.changeCarM("Volvo", 70000)
	fmt.Println("My car: ", myCar) // expected `Volvo 70000`, but no changes made.

	// Lets try a method with a pointer receiver.

	// You don't have to create a seperate pointer like-
	(&myCar).changeCarMP("test", 000) // this is one way to create pointer for a type or object. OR-
	fmt.Println("My car: ", myCar)
	// This is an another way
	var myCarP *car                  // declared a variable as pointer of type car
	myCarP = &myCar                  // assigned the address of myCar to the pointer myCarP
	myCarP.changeCarMP("test1", 000) // calling method by myCarP
	fmt.Println("My car: ", myCar)

	// But Go will automatically detect variable or pointer for method by the method receiver.
	myCar.changeCarMP("Porshe", 100000)
	fmt.Println("My car: ", myCar) // myCar changed to Porshe, 100000. Because we used the pointer to change-
	// the object myCar.

	// If a method takes a pointer receiver, it's good to convert all the methods to take pointer receivers. It
	// doesn't matter if they have to change the original value or not.
	// If you don't know what to use a value receiver ot a pointer receiver, you can apply these simple rules.
	// Use a pointer receiver when you want to modify the receiver so the value type or whe you want to avoid-
	// copying large amount of data that is automatically copied when passing values.

	// The method declarations are not permitted on named types that are themselves pointer types.
	// type distance *int // named type which is a pointer to int.
	// func (d *distance) m1(){ // error: distance is already a pointer type.
	//	 fmt.Println("Just a message")
	// }
	// We create methods only for value types, not for pointer types.
	

}
