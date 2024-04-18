**Untyped(typeless) Constant**

* const num = 6 // An untyped constant
* An untyped constant makes it possible for us to use constants in Go with great freedom.
* We can perform operations with different types of untyped constants without changing it's type.
    const x = 3
    const y = 2.2 * x // new const with float * int (possible to multiply float value with an int ).

* When an untyped constant is used in a context that requires a type, a type will be inferred and untyped constant has a default type, an implicit type that it transfers to a value if a type is needed but none is provided. An untyped constant is converted to default type only when needed in expression. OR Another way to think about untyped constant is that they exist in a kind of ideal space of values which is less restrictive than Go's full type system. To do anything with them we need to assign them to variables and when that happens the variable, not the constant itself, need a type and the constant can tell the variable what type it should have.

    const r = 7

    var rr float64 = r

    sr := r + rr

    fmt.Printf("Sum of r + rr : %v & Type of sr is: %T\n",sr,sr)