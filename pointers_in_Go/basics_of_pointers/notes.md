# Pointers

* The Computer Memory(RAM) can be thought of as a sequence of boxes or cells, placed one after another in a line. Each cell is labeled with a unique number, which increments sequentially; this number is the address of the cell or its memory location.

* Each cell holds a single value. Everything the CPU does is fetching and storing values into memory cells.

### What is a Variable?

* A variable is just a convenient, aphanumeric nickname or label for a memory location.

* When we write var a int = 7, we create a label called a for a memory location where the value 7 of type int will be stored.

* In a nutshell: memory is just a series of numbered cells, and variables are just nicknames for memory locations assigned by the compiler.

### What is a Pointer?

* A pointer is a variable that stores the memory address of another variable.

* The pointer points to memory address of a variable, just as a variable represents the memory address of a value.

* A pointer value is the address of a variable or nil if it hasn't been initialized yet.

* Pointers have the power to mutate or change the data they are pointing to.

* **Unlike C, Go has no pointer arithemetic. In C, pointer can be incremented or decremented. In Go it is a compile time error.**

* To access the memory address of a variable, Go provides the ampersand operator, also known as address of operator.

* To find out the value of pointer points to use the * operator which is also called the dereferencing operator. It is placed before a pointer variable and returns the value in that memory address.

* The * operator denotes the pointers underlying value. Or in other words, * means take this address and give me the value that exists at the address.

* **The * Operator**

* When you see a star(*) in front of a type, as in *float64 `var n *float64`, it means something completely different than you see a star in front of a pointer, as in `*p`.

* in front of a type, as in *float64 `var n *float64` means type description. We are declaring a pointer to that type, in this case float64.

* When we see * in front of a pointer, it means the dereferencing operator or that we want the value that the pointer is pointing to.

**If you have a value, you turn it into an address or into a pointer by using the ampersand operator. so &value returns pointer( ptr := &x ).**

**If you have a pointer, you turn it into a value by using the * operator (n := *ptr).**