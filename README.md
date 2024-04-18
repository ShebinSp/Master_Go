# KNOW BEFORE YOU GO TO MASTER GO
--------------

  Go work space is directory where the go code resides. By default the go work space resides in an environment variable called GOPATH.

* type 'go env' in terminal to check GOPATH and other variables (on linux).

--------------
In the go directory there will be 3 subdirectories at its root, which are **bin, pkg, src**

* src contains the source files for your own or other downloaded packages.

* pkg contains go package objects aka package archives and they are non executable files or shared libraries ending in .a. The files from the pkg directory contain the compiled package binary code, which is customized for a machine architecture and the go version. You cannot run these packages directly as they are not binary files. They are typically imported and used inside other executable packages.

* bin contains compiled and executable Go programs. When you run go install on a program directory Go will put the executable file under this folder.

Package
-------

  Every Go program must include a package. In Go programming, a package is a collection of Go source files that are organized together in a single directory. Each package typically contains related code that serves a particular purpose or functionality. Packages facilitate code organization, reuse, and modularity in Go programs.

  In Go programming, the package main declaration at the beginning of a Go source file indicates that this file belongs to the main package. The main package is a special package in Go that serves as the entry point for a Go program. When you compile a Go program, the compiler looks for a main package and expects it to contain a main function.

**Key characteristics of packages in Go include:**

1. Encapsulation: Packages encapsulate code and provide a way to control access to its components through the use of identifiers. This helps in building modular and maintainable codebases.

2. Visibility: Go uses a mechanism to control the visibility of identifiers within a package. Identifiers that start with a capital letter are exported from the package and can be accessed by code outside of the package, while those starting with a lowercase letter are unexported and can only be accessed within the package.

3. Imports: Go programs can import packages using the import keyword followed by the package path. Imported packages provide access to their exported identifiers, enabling code reuse and composition.

4. Dependency Management: Go's package management system, including tools like go mod, facilitates dependency management by allowing developers to specify dependencies and manage versioning.

5. Standard Library: Go comes with a rich standard library comprising packages for various functionalities, such as I/O operations, networking, concurrency, cryptography, and more. These packages are readily available for use in Go programs without the need for additional installation.
----------------
* To run a .go file there are several ways,

1. From terminal `go run main.go`

2. From vs code integrated terminal  `go run main.go`' or `go run ./`.

3. From vs code integrated terminal `ctrl + F5` to run in debug mode.
----------------

Go Command
----------

The `go` command is the prime tool we have to interact with the Go language. It gives us the possibility to compile and execute Go applications. `go` is a tool for managing Go source code.

* `go run`: it compiles and runs the application. It doesn't produce an executable

* `go run file.go` : compiles and immediately runs the Go program.

* `go build` : It just compiles the application. It produces an executable.

* `go build file.go` : compiles a bunch of Go source files. It compiles packages and dependencies.

* If you run `go build`  it will compile the files in current directory and will produce an executable file with the name of the current working directory.

* `go build -o app`  will produce an executable file called app.

* Compiling for Windows: `GOOS=windows GOARCH=amd64 go build -o winapp.exe`

* Compiling for Linux : `GOOS=linux GOARCH=amd64 go build -o linuxapp`

* Compile for Mac : `GOOS=darwin GOARCH=amd64 go build -a macapp`

* `go install` : Both `go install` and `go build` will compile the package in the current directory.

* If the package is main, `go build` will place the resulting executable in the current directory. and `go install` will move the executable to `GOPATH/bin`.

* When running `go install` you use paths relative to `GOPATH/src`.
example : go install katana/go/src/file_name (running from home terminal).

Formatting Go Source Code (gofmt)
---------------------------------

* Go stongly suggests certain styles.

* `gofmt` which comes from golang formatter will format a program's source code in an idiomatic way that is easy to read and understand.

* This tool, `gofmt` , is built into the language runtime, and it formats Go code   according to a set of stable, well-understood language rules.

* We can run it manually at the command line or we can configure out IDE, for example VS Code, to run gofmt each time a file saved.


  Example 1:
    `gofmt -w main.go` ( -w rewrite the source code or overwrites the file with suggested go format ).

  Example 2 :
    `gofmt -w -l master_go` ( -l flag tells  which file is modified, master_go is the name of the directory, if you want to certain files in directory).

  Example 3: 
    `gofmt` (from the file directory, which will format all the files in current directory ).

  Example 4:
    In vs code, right click -> format document will format the go code or `ctrl + shift + I`.


Comments in Go
-------------
* To comment a line of code : //

* To comment a block of code : /*  */



    Comments are helpful to understand the code later. For documenting code, it is considered idiomatic to always use double slashes (//) syntax. You would only use the slash star syntax for debugging. Inline comments occur on the same line of a statement following the code itself. Inline comments should not be often, but can be effective for explaining tricky, non obvious parts of code.

Example:

  // an integer variable declared

  var integer1 int

  interger1 = 7 // value 7 assigned <-this is an inline comment


Naming Conventions in Go
------------------------
  
* Naming conventions are important for code readability and maintainability.

**Naming Conventions**

  * Names start with a letter or an underscore( _ )

  * Case matters: quickSort and QuickSort are different variables.

  * Go keywords (25) can not be used as names.

  * Use the first letters of the words
      var mv int   // mv -> max value

  * Use fewer letters in smaller scopes and the complete word in larger scopes
      var packetsReceived int // NOT OK, name too long (to verbose)
      var n int // OK -> no. of packet received
      var taskDone boll // ok in larger scopes ( on package level )

  * An another convension is to write acronyms in all caps.
      writeToDB := true // OK, idiomatic (DB is an acronym for database ).
      writeToDb := true // not OK.

  **Recommended Go naming conventions:**

    * Use Camel Casing

    * maxValue := 100 // recommended
      max_value := 100 // not recommended