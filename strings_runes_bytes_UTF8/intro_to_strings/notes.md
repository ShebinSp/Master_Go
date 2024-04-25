# Introduction to Strings

* Strings are defined between double quotes and not in single quotes like in JavaScript or even Python.

* Strings in Go UTF-8 encoded by default.

* If you want to use double quotes inside the double quotes you must use a backslash in front of each double quotes to escape it or using backticks `He says: "Hello"` .

* A string literal enclosed in backticks is called a raw string and it is interpreted literally. Back slashes or other special characters have no special meaning.

* `fmt.Println("C: \\Users\\User1")`  each backslash will escape the next one.

#### **String Concatenation**

* To concatenate strings use the addition operator or the plus sign.

* Each time you can concatenate to a string value, Go creates a new string because the strings are immutable in Go and this is not efficient.