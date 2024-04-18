### For Loop

**A `for` statement is used to execute a block o code repeatedly. One of the reasons a computer is so useful is  that it can repeat operations multiple times very quickly. In programming a `for loop` represents one way to execute repetitive tasks.**

*In contrast to other programming languages where there is also a `while` loop, in Go there is only one loop statement, the `for` loop statement.*
----------------
### Structure of for loop in Go

`for i := 0; i < 10; i++ {

  fmt.Println(i)

}`

* for -> the keyword which represents a for loop statement

* i := 0 -> The initialization statement.

    It runs only once, before the loop begins and it declares a temporary variable called `i` that exists only inside the for loop.

* i < 10 -> Testion boolean expression.

    This is a testing boolean expression  that returns `true` or `false`. The loop block of code continues executing as long as this condition is true. The loop will start only if this condition is true.

* i++ -> Post statement.

    This is the post statement of the for loop and is executed after each step of the loop, after each iteration. In this case `i` is incremented by 1 after each iteration. Between the pair of curly braces we have the for block of code, in this case there is only one statement which is `fmt.Println(i).