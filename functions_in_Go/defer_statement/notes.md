# Defer Statement

* **A `defer`statement defers or postpones the execution of a function until the surrounding function returns either normally or through a panic.**

* Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in a last in first out or LIFO order.

* The deferred calls arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

### Why difer Statement?

* `defer` is used to ensure that a function call is performed later in the program's execution, usually for purpose cleanup.

* Let's suppose we want tho create a file, write to it, and then close the file when we are done with it. Immediately after creating the file variable, we `defer` the closing of that file. The function that closes the file will be executed at the end of the enclosing function main(). So after the writing to the file operation has finished.