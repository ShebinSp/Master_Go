# Package strings

* Package strings is used to manipulate UTF8 encoded strings.

* `strings.Contains()` returns true if any passed Unicode points are within the string and false otherwise.

* `strings.Count()` returns the number of instances of a substring.

* To make a string to UPPERCASE use `strings.ToUpper()`, and to lowercase `strings.ToLower()`.

* When comparing strings using the normal double equals operator, we get the quickest and most optimised solution. But letter case matters here. `"Go" == "go"` is `false`.

* A standard solution is create a lowercase version of each string and compare them. It works but not efficient at all because the `strings.ToLower()` function loops over each rune in the string and converts it to lowercase and returns the newly formed string.

* `Strings.EqualFold()` is used for case insensitivity. This is the recommended way to compare strings when we don't want to take into account the latter case.