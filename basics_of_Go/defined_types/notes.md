### Defined Types

* **A defined type is also called a named type, is a new type created by the programmer(using keyword `type`) from another existing type which is called the underlying or source type(int, float, string, bool etc).**

* A new defined type must have a new name and can have its new methods.

* The underlying type provides the representation, operations and size of the newly defined type.

* Even though the defined type and the source type share the same type representation, operations and size they are different types. A new type it's not just an alias for an existing type, it's a completely new type.

* **If we want to perform operations b/w source and defined types, we must convert one type into the other type. A type can be converted to another type if they share the same underlying type.**

* There is no type-hierarchy in Go.

### Why do we need Defined Types

* We can attach methods to newly defined types.

* Type safety. We must convert one type into another to perform operations with them.

* Readability. When we defined a new type let's say *type usd float64* we know that new type represents the US Dollar not only floats.

### Points to Note:

* type age int 'age' is the new type and 'int' is it's underlying type. underlying types are the predeclared types.

* `type km float64 // float64 = underlying type
type mile float64`
Even if both km and mile have underlying type float64 still these are different types.

* One type can be converted to another type if they share the same underlying type.

* we can perform operations on the same defined types. The new type will inherite all the operations of the source type, means all operations that are available to floats are also available to these newly defined types.

* type of `a` is main.speed, means - type `speed` is defined within package named `main`.