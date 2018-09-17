# A tour of Go

refer to: `https://tour.golang.org/list`

- packages

        rand.Intn(10)
- imports
  
  math.Sqrt(7)
- exported names
- functions
- multiple results
- named return values (Naked return)
- variables (can be function or package level)
- variables with initializers
  
        If an initializer is present, the type can be omitted
- short variables declearation
  
        `:=` is only available inside a function, it can not be used out of a function
- basic types
  
        bool
        string
        int
        uint
        byte
        rune
        float32 float64
        complex64 complex128 (math/cmplx)
        // print type and value
        fmt.Printf("Type: `%T` Value: `%v`\n", MaxInt, MaxInt)

- zero values

        number (0), string (""), bool (false)
- type conversions

        need explicit conversion operation   T(v)
- type inference
- constants
        Constants cannot be declared using the := syntax.
- Numeric Constants

        numeric constants are high-precision values.An untyped constant takes the type needed by its context.
- for

        Go has only one looping construct, the for loop.
- for continued

        The init and post statements are optional.
- for is Go's while
- forever
- if
- if with a short statement
  no `()` but `{}`
- if and else

        Variables declared inside an if short statement are also available inside any of the else blocks.
- switch
  
        Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that `Go only runs the selected case`, not all the cases that follow.
- switch evaluation order
- switch with nocondition
- Defer

        defers the execution of a function until the surrounding function returns
- stacking defers
  When a function returns, its deferred calls are executed in last-in-first-out order.

- pointers

        `*` `&`
        Unlike C, Go has no pointer arithmetic.
- struct
- struct fields

        Struct fields are accessed using a dot
- pointers to struct

        To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
- struct literals
- arrays
- slices

        a[low : high]  This selects a half-open range which includes the first element, but excludes the last one.
- slices are like references to arrays
  
        A slice does not store any data, it just describes a section of an underlying array.
        Other slices that share the same underlying array will see those changes.
- slice literals

        like an array literal without the length
- slice defaults

        a[0:10]
        a[:10]
        a[0:]
        a[:]
- slice length and capacity

        A slice has both a length and a capacity.
        The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
- nil slice

        The zero value of a slice is nil.
- create a slice with make

        a ：= make([]int, 5, 6)
- slices of slices
- appending to a slice

        `func append(s []T, vs ...T) []T`
- range

        When ranging over a slice, two values are returned for each iteration. Th`e first is the index`, and `the second is a copy of the element at that index`.
- range continued

        You can skip the index or value by assigning to _.
- maps

        A map maps keys to values.
        The zero value of a map is nil.
        The make function returns a map of the given type, initialized and ready for use.
- map literals
- map literals continued

        If the top-level type is just a type name, you can omit it from the elements of the literal.
- mutating map

        m[key] = elem
        elem = m[key]
        delete(m, key)
        // Test that a key is present with a two-value assignment:
        // If key is in m, ok is true. If not, ok is false.
        // If key is not in the map, then elem is the zero  value for the map's element type.
        elem, ok = m[key]
- function values

        Functions are values too. They can be passed around just like other values.
        Function values may be used as function arguments and return values.
- function closures
- methods

        A method is a function with a special receiver argument.
- methods are functions
- method continued

        You can declare a method on non-struct types, too.
        In this example we see a numeric type MyFloat with an Abs method.
        `You can only declare a method with a receiver whose type is defined in the same package as the method`. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
- pointer receivers

        Methods with pointer receivers can modify the value to which the receiver points.
        With a value receiver, the Scale method operates on a copy of the original Vertex value.
- pointers and functions

        functions with a pointer argument must take a pointer
        while methods with pointer receivers take either a value or a pointer as the receiver when they are called
- pointers and functions(2)

        Functions that take a value argument must take a value of that specific type
        while methods with value receivers take either a value or a pointer as the receiver when they are called
- choosing a value or poiner receiver

        there are two reasons to use a poiner receiver:
        The first is so that the method can modify the value that its receiver points to
        The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct
- interfaces

        An interface type is defined as a set of method signatures.
- interfaces are implemented implicitly
- interface values
- interface with nil underlying values
- nil interface value
- the empty interface

        An empty interface may hold values of any type.
- type assertion

        A type assertion provides access to an interface value's underlying concrete value. `t := i.(T)`
- type switches

        switch v := i.(type) {
            case T:
            //
            case S:
            //
            default:
            //
        }
- stringers (interface)

        A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
- Error (interface)
- readers (interface)
- images
- Goroutines
- channels
- buffered channels
- range and close

        for i := range c //  receives values from the channel repeatedly until it is closed
        v, ok := <-ch // ok is false if there are no more values to receive and the channel is closed.
- select

        A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready
- sync.Mutex