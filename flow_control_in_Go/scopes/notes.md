### Scopes in Go

* Scope means visibility, who can access certain names like functions or variables.

* The scope or the lifetime of a variable is the interval of time during which it exists as the program executes.

* A name cannot be declared again in the same scope (for example a function in the package scope), but it can be declared in other scope.

#### In Go there are 3 Scopes:

1. **File scope** -> import "fmt" // file scope

* Import statements are file scoped and only visible in the current file. If you use a function from this package, for example Println(), in another file of the same package you have to import the package there as well.

* Importing the package as another name, as an alias is permitted in Go.

2. **Package Scope** - > var n = 10 // outside a function

* Variables or constants declared outside any function are package scoped and are visible to all files of the package.

* Package scoped variables or constants can be used in other files in the same package, so no compile time error because of unused variable or const.

3. **Local (block) Scope** -> func main(){n := 11}

* The variables or constants declared inside a function called local or block scoped.

* They are visible until the end of the curly braces of the function it declared.

* The local scoped variables or constants  shadow the package scoped variables or constants.