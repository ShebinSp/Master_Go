# Strings in Go

* Go has two additional integer types called byte and rune that are aliases for uint8 and int32 data types. In Go, the byte and rune data types are used to distinguish characters from integer values.

* Go doesn’t have a char data type. It uses byte and rune to represent character values.

* Characters or rune literals are expressed in Go by enclosing them in single quotes, as in 'x' or '\n' . Rune literals such as ‘a’ , ‘b’, ‘c’, ‘x’ or ‘\n’ are represented using Unicode Code Points. A code point is a numeric value that represents a rune literal.

* The character encoding scheme ASCII which is a Unicode subset, comprises 128 code points.

* A string is a series of bytes values. A string is a slice of bytes and any byte slice can be encoded in a string value.

* The Go terminology for code points is runes . A rune represent a single unicode character. Rune 0x61 in hexadecimal represents the rune literal ‘a’.


-------------------------------------

* Strings are defined between double quotes and not in single quotes like in JavaScript or even Python.

* Strings in Go UTF-8 encoded by default.

* If you want to use double quotes inside the double quotes you must use a backslash in front of each double quotes to escape it or using backticks `He says: "Hello"` .

* A string literal enclosed in backticks is called a raw string and it is interpreted literally. Back slashes or other special characters have no special meaning.

* `fmt.Println("C: \\Users\\User1")`  each backslash will escape the next one.

#### **String Concatenation**

* To concatenate strings use the addition operator or the plus sign.

* Each time you can concatenate to a string value, Go creates a new string because the strings are immutable in Go and this is not efficient.