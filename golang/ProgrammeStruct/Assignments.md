## Assignments
Its simplest form has a variable on the left of the = sign and an expression on the right
> `x = 1  // named variable`  
> `*p = true  // indirect variable`  
> `person.name = "bob"  // struct field`  
> `count[x] = count[x] * scale  // array or slice or map element`

Each of the arithmetic and bitwise binary operators has a correspondig *assignment operator* allowing:
> `count[x] *= scale`

Numeric variables can also be incremented and decremented by ++ and -- statements:
> `v := 1`
> `v++  // same as v = v + 1`
> `v--  // same as v = v - 1` 

### Tuple Assignment
*tuple assignment* allows several variables to be assigned at once.
> `x, y = y, x`
> `a[i], a[j] = a[j], a[i]`

When a function call is used in an assignment statement, the left-hand side must have as many variables as the function has results.
> `f, err = os.Open("foo.txt")  // function call returns two values`