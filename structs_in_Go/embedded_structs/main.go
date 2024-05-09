package main

import "fmt"

func main() {
	// NESTED or EMBEDDED STRUCTS
	// An embedded struct is just a struct that acts like a field inside another struct.

	type Contact struct {
		email, address string
		phone int
	}

	type Employee struct {
		name string
		salary int
		contactIfo Contact //  an embedded struct field
	}

	john := Employee{
		name: "John Keller",
		salary: 5000,
		contactIfo: Contact{
			email: "jK@company.com",
			address: "street 20, England",
			phone: 0012233344445,
		},
	}
	// Printing the details of employee john
	fmt.Printf("employee : %#v\n",john)

	// Accessing from embedded struct
	fmt.Println("Employee's email: ", john.contactIfo.email)

	// Update a field
	john.contactIfo.email = "joke@company.com"
	fmt.Println("Employee's new email: ", john.contactIfo.email)

	// We can also make a standalone struct using contact
	myContact := Contact{
		email: "myemail@gmail.com",
		address: "myaddress, India",
		phone: 002245135437,
	}
	fmt.Println("my contact: ",myContact)

}