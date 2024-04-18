### Alias Declarations

* An alias declarations has the form: `type T1 = T2` as opposed to a standard type definition which has the form: `type T1 T2`

* An alias declaration binds an identifier to the given type. **It's the same type with a new name.**

* Types with different name are different types, but there is an exception to this rule and that is the aliased types.

* `byte` and `uint8` are aliases or same type with different names. The same is applicable t rune and `int32` because `rune` is an alias for `int32`.

* Aliases can be used together in operations without type convertions we've seen at defined types.

* You should use aliases with caution, the are not for everyday use. They were introduced to support gradual code repair and the large scale refactoring.