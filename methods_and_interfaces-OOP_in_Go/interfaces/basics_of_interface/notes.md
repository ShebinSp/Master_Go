# Interface

* Interfaces is a core concept in Go and interfaces are everywhere.

* An interface is a collection of method signatures that an object which is most of the time a named type, can implement.

* Interfaces define the behaviour of an object and can achieve polymorphism.

* Interfaces are not generic types like in other languages like Java.

* In Go, we declare an interface like a struct or any other named type, using the type keyword, a name, and then the `interface` type.

* The interface contains only the signature of the methods, but not their implementation. The signature of a method means its name, input, parameters and return values.

* To implement an interface in Go, we just need to implement all the methods in the interface. Go interfaces are implemented implicitly. If a type implements all the methods in the interface, then that type is said to implement that interface. This is very important and it makes Go special. It's like saying if it walks like a duck, swim like a duck and quacks like a duck, then it's a duck.

* Any type that implements the interface is also of type of the interface, this is very important.

* A variable of an interface type can hold any value of any type that implements the interface.