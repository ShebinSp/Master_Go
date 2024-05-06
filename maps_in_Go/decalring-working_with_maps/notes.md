# Maps in Go

* **A Map is a collection type just like an array or a slice and stores key:value pairs.**

* The main advantage of maps is that add, get and delete operations take constant expected time.

* All the keys and the values in map are statically typed and must have the same type.

* The keys in a map must be unique, but the values don't have to be unique.

* A Map allows us to quickly access a value using a unique key!

* We can use any comparable type as a key map. A comparable type is that type that supports the comparing operator which is the double equals sign(==).

* Even if it's possible, it's not recommended to use a float as a key, A float has some comparable issues.

* We can not compare a map to another map . We can only **compare a map to nil.**

* Maps are unordered data structures in Go.

### Points:

* we can get the corresponding value of a key even if the map is not initialised

I* f an element doesn't exist in a map, it returns a zero value of that type.

* When declaring a map, we can use only comparable types as keys. ie, we can't use a slice because a slice is not comparable.

* We can declare arrays as key, since arrays are comparable.

* It is illegal to populate an initialised map in Go.

* Initialise the map using the make() function or a map literal before add any element.

* If the key exists its updates it's value but if the key does't exist, it adds the key value pair.

* The default order of a type string key is 'A - Z' and 'a - z' in Go.

* We can iterate over a map but not indicated, because:

   1. The order of key-value pair changes.

   2. The maps in Go have been designed for a fast lookup time, not for fast looping.

* To delete a key-value pair from the map, use the built-in function `delete()`.