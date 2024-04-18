### Comments in Go

* To comment a line of code : // your code here

* To comment a block of code : /* your code here */



* **Comments are helpful to understand the code later. For documenting code, it is considered idiomatic to always use double slashes (//) syntax. You would only use the slash star syntax for debugging. Inline comments occur on the same line of a statement following the code itself. Inline comments should not be often, but can be effective for explaining tricky, non obvious parts of code.**

#### Example:

* // an integer variable declared

  var integer1 int

  interger1 = 7 // value 7 assigned <-this is an inline comment

  
### Naming Conventions in Go

* *Naming conventions are important for code readability and maintainability.*

#### **Naming Conventions:*

* Names start with a letter or an underscore( _ )

* Case matters: quickSort and QuickSort are different variables.

* Go keywords (25) can not be used as names.

* Use the first letters of the words

   var mv int   // mv -> max value

* Use fewer letters in smaller scopes and the complete word in larger scopes

* var packetsReceived int // NOT OK, name too long (to verbose)

   var n int // OK -> no. of packet received

   var taskDone boll // ok in larger scopes ( on package level )

* An another convension is to write acronyms in all caps.

    writeToDB := true // OK, idiomatic (DB is an acronym for database ).

    writeToDb := true // not OK.

* Recommended Go naming conventions:

   * Use Camel Casing

   maxValue := 100 // recommended

   max_value := 100 // not recommended