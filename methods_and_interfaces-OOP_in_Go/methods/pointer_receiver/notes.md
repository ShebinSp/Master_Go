# METHOD - Pointer Receiver

* Calling a function with arguments makes a copy of each argument value. If a function needs to update a variable passed as argument, or if an argument is so large that we wish to avoid copying it, we must pass the address of the variable, not the variable. The same goes for methods that need to update the receiver.

* If a method takes a pointer receiver, it's good to convert all the methods to take pointer receivers. It doesn't matter if they have to change the original value or not.

* If you don't know what to use a value receiver ot a pointer receiver, you can apply these simple rules.

* Use a pointer receiver when you want to modify the receiver so the value type or whe you want to avoid copying large amount of data that is automatically copied when passing values.

* The method declarations are not permitted on named types that are themselves pointer types.

` type distance *int // named type which is a pointer to int.`
`     func (d *distance) m1(){ // error: distance is already a pointer type.`
`        fmt.Println("Just a message")`
`     }`

* We create  methods only for value types, not for pointer types.