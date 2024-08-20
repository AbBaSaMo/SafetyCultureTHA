### Declaring things
<hr/>
Good style to use factored import

```go
import (
    "fmt"
    "math"
)
```
versus
```go
import "fmt"
import "math"
```

<hr/>
Exported names from a package are capetalised e.g. `math.Pi` whereas non-exported names are not and are not accessible external to the package e.g. `math.pi`

<hr/>
Declare functions like so

```go
func name(param1 type, param2 type, ...) returnType {
    return ...
}
```

Or if a bunch of consecutive params share a type, omit the type except from the last occurence

```go
func name(num1 , num2 int, ...) returnType {
    return ...
}
```

Naked returns: Return values can be named in the function signature where return without arguments will return the named vars.

**But limit this to short functions to avoid problems with readability**

```go
func name(param type, param type, ...) (ret1, ret2 type) {
    return 
}
```

Functions can refer to variables outside of their scope but then become bound to that variabl/depend on it. These are called closures.

<hr/>
Declar variables with var with type at the end. Examples:

`var a, b, c int`
`var a, b, c int = 1, 2, 3`
`var a, b, c = 1, "1", true` 

Inside a function, `:=` can be used with type implied **for non-constants** and for numbers, the type is inferred by the precision

`a, b, c := 1, "1", true`

But outside of a function keywords are needed i.e. `var`, `func` etc

And unassigned vars are given a zero value like `0`, `false` or `""` for their respective types

<hr/>

Types include

- `bool`
- `string`
- `int` from `int` to `int64` and likewise for `uint`
- `byte` as alias for `uint8`
- `rune` as alias for `int32`
- `float32` and `float64`
`complex64` and `complex128` for complex numbers

And type convertions are done using syntax like `string(myVar)`

Types must be explicitly converted before assignment to vars of different types.

<hr/>

Functions can also be assigned to variables

`my_func := func(x int) int {// block}`

<hr/>

No pointer arithmetic in go

```go
// pointer to address of var of type T
var p *type

// gets address of var
p:= &var1

// gets value of pointer's referenced var
num = *p
```
<hr/>

Structs are collections of fields. Access fields using dot notation and initialise using `StructName{f1, f2, ...}` or `Struct{fN: ...}` if you want to just assign `fN` etc

And struct fields can also be accessed (using dot notation) via a pointer to the struct

```go
type Struct_Name struct {
    // fields
}

```

<hr/>

Slices are references to a subset of an array.
Create using `my_arr[start, end]` where end is exclusive.
As they are references, changing elements of a slice changes the underlying array and other slices of that same array will see the underlying changes.

Slices have a length which is their own length and a capacity which is the lenght of the array they are referencing. Obtain using `len(s)` and `cap(s)`

The zero val of a slice is `nil` with no len or cap as it has no underlying array

[Appending to slice](https://go.dev/tour/moretypes/15)

<hr/>

Maps can be made using `make()`, e.g.

```go
// makes a map with string keys and Value_Type values
m = make(map[string]Value_Type)
```

Maps that are `nil` and can't have keys added

### Loops and conditionals
<hr/>

Only for loops exist and no `()` are used:

```go
for init; conditional; post {

}
```

Init statement e.g. `i int = 0` and post e.g. `i++` are optional

```go 

var sum = 1

// and the semicolons are optional
// this is basically a while loop
for ;sum < 1000; {
    sum += sum
} 

```

<hr/>

For loops can be written using range which iterates over a slice/map and for each element returns the index and element

`for idx, elem := range slice {}`

This skils the index if you don't care for it
`for _, elem := range slice {}`

And this the element
`for idx := range slice {}`

<hr/>

Conditionals don't need `()`

Conditionals can hava an init that is run before the condition like with for loops but any vars declared are scoped to the conditional (as well as any elses) e.g.

```go
if a := 1; a < 100 {
    // block code
}

```

Switch syntax doesn't need break, only the selectd case is run

```go
switch init; variable {
    case match1:
        // block
    case match2:
        // block
    default:
        // block
}
```

### Misc

Defer statement stalls execution of a function until the parent returns

```go 
func main () {
    defer fmt.Println("...")
    return ...
}
```

For multiple deferrals, they are executed LiFo as they are pushed onto a stack

<hr/>

Generic functions: Functions can use multiple types that fulfil some sort of constraint e.g. the following accepts an array of any type that has the `comparable` constraint and then a value of the same condition. 

`func MyName[T comparable](slice_of_comps [T]T, my_comp T) int {}`

<hr/>

Generic types: as above

```go
type List[T any] struct {
    next *List[T]
    val T
}
```

### Methods in Go

When you make a method, the 'class' it is for must be in the same package and cannot be degined in another package including inbuilt types.

This 'attaches' the method to the type

```go
// create a type
type  CustomType struct {
    field1   string
}

// attach function to type: create a meathod
func (c CustomType) methodName(){
    // can access CustomType's properties in here
}

// now we can call the method on CustomType
newC CustomType = ...
newC.methodName()
```

And methods can be created on non-struct types

```go
type CustomFloat float64

func (c CustomFloat) methodName(...) {
    //block
}
```

<hr/>

If a type is passed as a receiver i.e. the thing the method is attached to, then the function can access the types fields but not modify it.

If we want the method to be able to modify the type's fields, we need to pass a pointer to the type as the recevier i.e.

```go
type Vertex struct {
	X, Y float64
}

// cannot modify X and Y
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// can modify X and Y
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```

When the receiver is a pointer, then the method can be called by using e.g. `V.Scale()` or a pointer to `V` like `&V.Scale()` as the former is interpreted as the latter

On the other hand, a function that modifies `V` must take a pointer to `V` other wise a compile error occurs e.g. `Scale(V, 5)` fails while `Scale(&V, 5)` works assuming `Scale()` tries to modify fields in `V`

This is the same for methods that don't take a pointer i.e. functions MUST take the type declared while methods will translate appropriately e.g. `(*p).MyMethod()` -> `p.MyMethod()`

<hr/>

When creating methods on a type, either use pointers for ALL receivers or don't

As above, it lets us modify fields directly but also prevents the need to copy the value over when the method is called.

### Interfaces in Go

We can define signatures for functions to then implement

```go
// interfaces: generally should try limit num methods to 1 or a few
type geometry interface {
    area() float64
    perim() float64
}

// struct for shapes
type rect struct {
    width, height float64
}

// methods for shapes that fulfill our interfaces
// 
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
```

<hr/>

If the value implementing the interface is `nil`, then the method is called with `nil` as the receiver.

Interfaces are a tuple of the value fo the type implementing the method signature as well as the type of that type e.g. `(50, int)` is an int of value 50 that implements the interface's method

e.g.

```go
type I interface {
    MyMethod()
}

type T struct {
    S string
}

// the type is nil here but the interface is not
// the interface type has (<nil>, *T)
var myInterface I
var myTypeImplementingInterface *T
myInterface = myTypeImplementingInterface
```

If the interface itself is `nil` with no assignment it throws an errror and the tuple has `(<nil>, <nil>)` as ther is no underlying type

```go
type I interface {
    MyMethod()
}

var myInterface I

// throws error
I.MyMethod()
```

<hr/>

To check if a variable of type interface is of type T e.g. float, use type assertion, otherwise use type assertion to get the underlying value fo an interface

```go
var i interface{} = "hello"
s := i.(string) // this is type assertion, it gives access to the underlying value of the interface

s, ok := i.(float64) // ok is abool indicating if it's of that type 

s := i.(float64) // assertion on a type that the underlying value is not of will throw an error


```