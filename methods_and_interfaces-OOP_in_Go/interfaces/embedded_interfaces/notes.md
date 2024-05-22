# EMBEDDED INTERFACES

* In Go an interface cannot implement other interfaces or extend them. Inheritance is not supported.

* What we can do is to create a new interface by merging two or more interfaces. This is known as interface embedding.

* Declaration of an interface means specifying methods belonging to it. Besides methods, an interface can also include other interfaces in the current or another package and import it.

* When include an interface in another interface, all the methods from the embedded interfaces will be added.

* Satisfying an interface means implementing all methods of the interface.

* **Circular embedding of interfaces is disallowed and will be detected at compile time.** 
* Circular embedding is, if there are 3 interfaces `interface1`, `interface2` and `interface3`. `interface1` is embedding `interface2`, `interface2` is embedding `interface3` and `interface3` is embedding `interface1`, this is circular embedding.