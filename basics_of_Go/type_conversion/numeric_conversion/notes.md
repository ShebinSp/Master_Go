### Converting Numeric Types

* According to Go specification the term used is converting and not casting like in other languages.

* A conversion changes the type of an expression to the type specified by the conversion.

* A conversion may appear explicitly in the source or it may be implied by the context in which an expression appears.

* When converting a float to int, Go will omit the decimal values of float.

*     var a = 5 // int type
        fmt.Printf("Type of a: %T\n",a)
 
        var b int64 = 2
        // a = b // Error: mismatched type int and int64
	    a = int(b) // int64 converted to int.

    In Go, types with different names are different types, so 'a' and 'b' are not of the same type. Explicit conversions are required when different numeric types are mixed in an expression or assignment. For instance int64 and int are not same type even though they may have the same size on a particular architecture. There is only one exception to this rule and that is aliased types.