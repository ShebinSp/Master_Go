# PASSING POINTERS TO FUNCTIONS

* Pointers are also copied when using them as arguments, the function creates a new pointer variable, a copy that points to the same memory address.  All these variables that are copied when passed to functions are local to the function and available only inside the function body.

* Functions work on copies, not on originals in Go. This is what pass by value means and in Go everything is passed by value.