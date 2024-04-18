Types and Zero Values
----------------------
Go is a statically typed programming language, and this means it does `type` checking at compile time. In Go, you must provide a type for each variable you declare or it should be able to infer, means to automatically detect the type of the variable at compile time, if not provided explicitly. In contrast a dynamically typed programming language like Python or PHP does type checking at runtime, so possible errors related to types are detected later when the program is running.
    Each variable stores a value with the same type, the type of a variable cannot be changed once declared. The main purpose of a statically typed language is to prevent as many bugs as possible, and it does this as early as possible.

**Zero Value**
  The zero value mechanism of Go ensures that the variable always holds a well defined value of its type. In Go there is no such thing as uninitialized variable. This simplifies the code and often avoids parts that are hard to find.

Go Zero Values:
--------------
* numeric types : 0
* bool type : false
* string type : "" ( empty string )
* pointer type: nil


## Go Data Types (Built-in)

* A type determines a set of values together with operations and methods specific to those values.

* There are predeclared types, introduced types with type declarations and composite types: array, slice, map, struct, pointer, function, interface, and channel types.

### Predeclared, Built-in Types

### Numeric types
    * int8, int16, int32, int64
    * uint8, uint16, uint64 : used to represent unsigned(positive) integers.
    * uint is an alias for uint32 or uint64 based on platform.
    * int is an alias for int32 ro int64 based on platform.
    * float32, float64: zero before the decimal point seperator can be omitted (-.5, -3,-0.1.4).
    * complex64, complex128
    * byte (alias for uint8).
    * rune (alias for int32).

### Byte & Rune Type

  In Go, the byte and rune data types are used to distinguish characters from integer values. Go doesn't a character data type, it uses byte and rune to represent character values.

  `var r rune = 'f'
    fmt.Printf("Type  of r: %T\n", r) // Type is 'int32' so, rune is an alias forint32.
    fmt.Println("r(ASCII): ",r) // 102 is the ASCII code for 'f'
    fmt.Printf("r(hexadecimal): %x\n",r) // 66 is the hexadecimal ASCII for 'f'`

### Bool Type

* Pre-defined constants `true` or `false`.
    var b bool = true
      fmt.Printf("Type  of b: %T\n", b)

### String Type
* Unicode chars written enclosed by double-quotes. Each character occupies b/w 1 & 4 bytes.

* A string value is a (possibly empty) sequence of bytes.
    var s string = "Hello Go!"
        fmt.Printf("Type of s: %T\n",s)


## Go Data Types (composite)

### Array & Slice Type

  * An array is a numbered sequence of elements of a single type, called the element type. The length of an array can be discovered using built-in function len(). The elements can be addressed by integer indexes starting from zero.

  * An array has a fixed length (we specify how many items are in the array when we declare it), but a slice has a dynamic length ( it can shrink or grow ).

  * Every element in an array or slice must be in same type.

### Map Type

  * A map is an unordered group of elements of one type, indexed by a set of unique keys of another type.

  * A map in Go is similar to dictionary in Python.

### Struct Type ( User Defined Type )

  * A struct is a sequence of named elements called fields, each of which has a name and a type.

  * A struct can be compared to class concept in Object Oriented Program.

### Pointer Type

  * A pointer is a variable that stores the memory address of another variable.

  * The value of an uninitialized pointer in nil.

      `var x int = 2
          ptr := &x
          fmt.Printf("Type of %q: %T with value(memory address(in hexadecimal)): %v\n","ptr", ptr,ptr)
          // To see the value a pointer points to use *
          fmt.Printf("%q points to the value: %v\n","ptr", *ptr)`
    #### output
          Type of "ptr": *int with value(memory address(in hexadecimal)): 0xc0000120f0
          "ptr" points to the value: 2

### Function & Interface Type

`func f(){} // empty function

fmt.Printf("Type of %q: %T\n","f",f)`
#### output
      Type of "f": func()

### Channel Type
  A channel provides a mechanism for concurrently executing function to communicate by sending and receiving values of a specific element type.