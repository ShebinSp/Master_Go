# Struct in Go

* A `struct` is a sequence of **named elements**, called **fields**. Each of them has a name and a type. We can also think of a struct as a collection of properties that are related together. They are useful for grouping data together to form records. Structs are useful to represent concepts from the real world like cars, people, or books.

* If you familiar with OOP from other language **you can think of a struct as of a class**. The struct fields are like instance attributes we define in OOP.

* Unlike traditional Object-Oriented-Programming, Go does not have a class-object architecture. Rather we have structs which hold complex data structures.

* ***A structs is nothing more that a schema containing a blueprint of data a structure will hold. This blueprint is fixed at compile time***. It's not allowed to change the name or the type of the fields at runtime. You can't add or remove fields from a struct at runtime.


# Basics of Structs

* Let's suppose our application will store data about books.

   * To do that we create some variables and initialize those variables.

    * title1, author1, year1 := "The Devine Comedy", "Davte Aligheri", 1320
    * title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

   * We see the information about our books

    * fmt.Println("Book1: ",title1, author1, year1)
    * fmt.Println("Book2: ",title2, author2, year2)

* *The problem with an application like this is that there is no connection between the variables title, author and year. They dont act as a single unit and don't represent the book concept from the real world. and if there are 1000 books we have to create 3000 different variables to store the information.*

* **A much more better approach is to create a struct that groups together these different concepts into a single concept called book.**

    type book struct {
        title string
        author string
        year int
        }

* Creating an instance or object or struct for struct type `book` with fields. It is a good practice structs with field names.

    bestBook := book{
        title: "Mackbeth",
        year: 1606, // year specified second
        author: "William Shakespeare",
    }
* You can omit the field names, but you have to keep the order and it is not a good practice

    * myBook := book{"The Devine Comedy", "Davte Aligheri", 1320}

* If we omit some fields when creating struct, then omited fields fields will be zero valued according to their type.

    aBook := book{
        title: "Just a random book",
    }
    fmt.Printf("a book: %#v\n",aBook) // author and year will be zero valued whcih are "" and 0.