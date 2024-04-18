// There are 3 Scopes:
//  - File Scope
//  - Package Scope
//  - Block (local)  Scope
package main
// file scoped
import "fmt" // import statements are file scoped and only visible in this file.  If you use a function from this package,
			 //for example Println(), in another file of the same package you have to import the package there as well.

// import "fmt" // Error: within the same scope we must have unique names
import f "fmt" // importing the package as another name, as an alias, is permitted in Go

// package scoped
var n = 10 // Variables or constants declared outside any function are package scoped and are visible to all files of the package.
const done = true

func main() {
	// local (block) scoped
	var task = "Running" // the var task is visible until the end of the curly brace of main() function. 

	// The var `n` is completely new variable under the local scope. Above is the package scoped var `n` and
	var n = 11 // this is local scoped

	fmt.Println(task, done)
	f.Println("n in local scope: ",n) // The local scope varaible `n` has shadowed variable `n` from package scope.

	f1()
}
func f1() {
	fmt.Printf("n in f1(): %v\n",n) // variable `n` from package scoped'
}