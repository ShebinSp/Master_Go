### Label Statements

* Labels are used in `break`, `continue`, and `goto` statements.

* It is illegal to define a label that is never used.

* In contrast to other identifiers, labels are not block scoped and do not conflict with identifiers that are not labels. They live in another space(You can name label with the same name you used for a variable).

* The scope of a label is the body of the function in which it is declared and excludes the body of any nested function.

* Most of the time labels are used to terminate outer enclosing loops.

**Labels are used to the compiler that we care about a specific place in the code and we are going to transfer execution to it at some point in the future. A label statement may be the target of a break, continue or goto statement.**