# Strings in Go

* Go has two additional integer types called *byte and rune* that are aliases for uint8 and int32 data types. **In Go, the byte and rune data types are used to distinguish characters from integer values.**

* Go doesn’t have a char data type. **It uses byte and rune to represent character values.**

* *Characters or rune literals* are expressed in Go by enclosing them in single quotes, as in 'x' or '\n' . Rune literals such as ‘a’ , ‘b’, ‘c’, ‘x’ or ‘\n’ are represented using Unicode Code Points. **A code point is a numeric value that represents a rune literal.**

* The character encoding scheme ASCII which is a Unicode subset, comprises 128 code points.

* A *string* is a series of bytes values. **A string is a slice of bytes** and any byte slice can be encoded in a string value.

* The Go terminology for code points is runes . A rune represent a single unicode character. **Rune 0x61 in hexadecimal represents the rune literal ‘a’.**

* The default value type for character value is rune.

* UTF8, which is the encoding scheme in Go is a variable length in coding scheme. It uses b/w 1 & 4 bytes for each one.

* The built in function len() returns the number of bites, not tokens or characters.

	