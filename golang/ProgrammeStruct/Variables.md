## Variables  
General form : `var name type = expression`.  
Either the `type` or `the = expression` part my be omitted, but not both.  
If the `the = expression` is omitted, the initial value is the *zero value*, which is `0` for numbers `false` for booleans, `""` for strings and `nil` for interfaces and reference types(slice, pointer, map, channel and function). The zero value of an aggregate type like an array or a struct has the zero value for all of ites elements of fields.  
It is possible to to declare and optionally initialize a set of varialables in a single declaration, with a matching list of expressions. Omitting the type allows declaration of of multiple variables of different types.  
> `var i, j, k, int    // int, int, int`  
> `var b, f, s = False, 2.3, "hello"   // boolean, float, string`  
 
A set of variables can also be initialized by calling a function that returns multiple values:
> `var f, err := os.Open(name)  // os.Open returns a file and an error`

### short variables declaration
**Within a function**, an alternate form called a *short variables declaration* be used to declare and initialize local variables. The type of *name* is determined by the type of *expression*. 
> `anim := gif.GIF{LoopCount: nframes}`  
> ` freq := rand.Float64() * 3.0`  
> `t := 0.0`  

A short variables declaration must declare at least one new variable.

### Pointer
A pointer value is the address of a variable.  
> `x := 1`  
> `p := &x  // &x equals address of x. type of p is *int, points to x `  
> `fmt.Println(*p)  // print 1`  
> `*p = 2  // equivalent to x = 2`  

The zero value for a pointer of any type is nil. Pointer is comparable, two pointers are equal if and only if the pointe to the same variable or both are nil. 

It is perfectly safe for a function to return address of a local variable.
> `fmt.Println(f() == f())  // f() *int {} , print false`




