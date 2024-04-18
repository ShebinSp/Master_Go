### Simple Statements

* A simple or short statement is a statement that is allowed inside another statement, like a `if` or a `switch` statement.

* Variables declared in a simple statement are available only in the branches of that statements.

* Most often, simple statements are used to handle errors.

// Simple or Short Statement

`if i, err := strconv.Atoi("7777"); err == nil {
        fmt.Println("No error, i is: ",i)
    } else {
        fmt.Println("Error: ",err)
    }`
* The first part of `if` `(i, err := strconv.Atoi("7777"))`  is known as an initialization statement and it can be used to setup local variables. For functions that return an error the initialization statement is very useful. The error is scoped only the `if` block that handles it. This way, the code ends up being readable.

* The second part of the statement is a boolean expression that returns true or false. If it's true, the branch below will be executed. If there is no error, `if err == nil` it will print integer. If there is an error, the else case will be executed.