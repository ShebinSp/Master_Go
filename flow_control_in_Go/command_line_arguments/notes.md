### Command Line Argument: os.Args 

  **The user input mostly comes from the keyboard. For the purpose of user input Go provides a standard library package called OS. This package communicates with the operation system and allows you to get OS functionalities; OS package contains a variable of type slice called Args.**

* `fmt.Println("os.Args",os.Args)`
   `go run main.go linux
   os.Args [/tmp/go-build2312086626/b001/exe/main linux]`

   * Here linux is the command argument and ^ this is the value of os.Args.
   
   * All values in os.Args are in string.

   * The first element is the path to the program which is executed.

   * Second argument is the program argument so linux.

* we can access the command line argument using indexes which start from zero.

    `fmt.Println("Path: ",os.Args[0])
    fmt.Println("First argument: ", os.Args[1])
    fmt.Println("No of items inside os.Args: ",len(os.Args))`