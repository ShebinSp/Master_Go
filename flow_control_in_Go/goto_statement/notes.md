### goto Statement

* **A goto statement transfers control to the statement with the corresponding label within the same function. It allows us to jump to any label inside same function.**

* Breaks and continues statements are restricted to be used only in for and switch statements. But goto statements don't have this restriction.

* It is not allowed to jump to a label after new variables have been introduced. Executing goto statements must not cause any variables to come into scope that were not already in the scope at the point of the goto. In other programming languages which has goto satatements can jump to anywhere.

* The use of goto statements is highly discouraged in any programming language because it becomes difficult to trace the control flow of a program making program difficult to understand and hard to modify.

* goto statements can be dangerous and should be used only with caution.