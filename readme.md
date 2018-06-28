# Go Programming

Code along to Derek Banas'/newthinktank tutorial on Go programming with my own additions.

Video: https://www.youtube.com/watch?v=CF9S4QZuV30
Text: http://www.newthinktank.com/2015/02/go-programming-tutorial

## Basics

### File setup

Every go file begins with a package declaration, providing code reuse.

To make a file executable we can use: `package main`.

### Packages

You can import packages using `import packagename`

E.g. the format package, `import fmt`, provides formatting for input and output.

### Comments

```go
// single line

/*
    Multi
    Line
    Comment
*/
```

### Hello World

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello, World")

}
```

to run: `go run hello-world.go`

### godoc

To see information on a package/method you can run godoc from the command line, e.g.:

```bash
> godoc fmt # outputs fmt package info
> godoc fmt Println # outputs info on Println method from fmt package
```

## Variables

Variables are statically typed, so they can't be changed once defined.

If a variable is declared and not used an error is thrown.

### Variable Declaration

Variables can be declared in many different ways in go.

Variables can be declated with type (string, int, etc.), however, when type is not declared go will automatically work out what type is being used based on the value.

Examples:

```go
var myVar int = 1

var myVar string = "hello"

var myVar = "hello again" //

var (
    varA = 2
    varB = 3
)

myVar := "hello"
```

### var vs :=

Top level variable assignment must be prefixed with `var`.

Within blocks `:=` may be used.

```go
package main

var toplevel = "Hello world" // var keyword is required

func F() {
        withinBlock := "Hello world" // var keyword is not required
}
```

### Integers and Floats

An int is a positive or negative number without decimals.

```go
var age int = 40
```

Int types:
 - uint8 : unsigned  8-bit integers (0 to 255)
 - uint16 : unsigned 16-bit integers (0 to 65535)
 - uint32 : unsigned 32-bit integers (0 to 4294967295)
 - uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
 - int8 : signed  8-bit integers (-128 to 127)
 - int16 : signed 16-bit integers (-32768 to 32767)
 - int32 : signed 32-bit integers (-2147483648 to 2147483647)
 - int64 : signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

Floats are numbers with decimals.

Float types:
 - float32
 - float64

```go
var rate float64 = 34.21432
```

Be careful when using floats for math!

### Booleans

Booleans can be either `true` or `false`

```go
var isAdmin bool = true

fmt.Println(isAdmin)
```


### Const

A constant is a variable with a value that cannot be modified

```go
const pi = float64 = 3.14159265359
```

### Strings

Strings can be wrapped either with double quotes "" or back ticks ``

```go
var name string = "dominic"

var name string = `dominic`
```

#### Using Strings

len(): get string length

```go
var myName string = "dominic"

fmt.Println(len(myName))
```

Concatenation with `+`

```go
var myName string = "dominic"

fmt.Println("My name is " + myName)
```

Newline with `"\n"`


```go
fmt.Println("Dear Sir,\n\nHello.")
```

Printf is used for format printing

`%f` is for floats

```go
fmt.Printf("%f \n", pi)
```

You can also define the decimal precision of a float

```go
fmt.Printf("%.3f \n", pi)
```

`%T` prints the data type

```go
fmt.Printf("%T \n", pi)
```

`%t` prints booleans

```go
fmt.Printf("%t \n", isAdmin)
```

`%d` is used for integers

```go
fmt.Printf("%d \n", 100)
```

`%b` prints in binary

```go
fmt.Printf("%b \n", 100)
```

`%c` prints the character associated with a keycode

```go
fmt.Printf("%c \n", 44)
```

`%x` prints in hexcode

```go
fmt.Printf("%x \n", 17)
```

`%e` prints in scientific notation

```go
fmt.Printf("%e \n", pi)
```

to run: `go run const-strings-boolean-printf.go`

## Basic Arithmetic

```go
package main

import "fmt"

func main() {

	fmt.Println("6 + 4 = ", 6 + 4);
	fmt.Println("6 - 4 = ", 6 - 4);
	fmt.Println("6 * 4 = ", 6 * 4);
	fmt.Println("6 / 4 = ", 6 / 4);
	fmt.Println("6 % 4 = ", 6 % 4);

}
```

to run: `go run basic-arithmetic.go`

## Logical Operators

There are three logical operators in Go:
 1. &&
 2. ||
 3. !

Example:

```go
package main

import "fmt"

func main() {

	// Logical Operators
	fmt.Println("true && false =", true && false) // false
	fmt.Println("true || false =", true || false) // true
	fmt.Println("!true =", !true) // false

}
```

to run: `go run logical-operators.go`

## If, Else, Else If, Switch

```go
package main

import "fmt"

func main() {

	age := 18

	if age >= 16 {
		fmt.Println("You Can Drive")
	} else {
		fmt.Println("You Can't Drive")
	}

	// You can use else if to perform different actions,
	// but once a match is reached the rest of the conditions aren't checked

	if age >= 16 {
		fmt.Println("You Can Drive")
	} else if age >= 18 {
		fmt.Println("You Can Vote")
	} else {
		fmt.Println("You Can Have Fun")
	}

	// Switch statements are used when you have limited options

	switch age {
		case 16: fmt.Println("Go Drive")
		case 18: fmt.Println("Go Vote")
		default: fmt.Println("Go Have Fun")
	}

}
```

to run: `go run if-else-elseif-switch.go`

## For Loops

```go
package main

import "fmt"

func main() {

	i := 1

	for i <= 10 {
		fmt.Println(i)

		i++
	}

	// Relational Operators include ==, !=, <, >, <=, and >=

	// You can also define a for loop like this, but you need semicolons

	for j := 0; j < 5; j++ {

		fmt.Println(j);

	}

}
```

to run: `go run for-loops.go`

## Array, Slices

### Arrays

An Array holds a fixed number of values of the same type

Arrays can be declared in different ways:

Using `var`:

```go
var toppings[3] string // length of 3

toppings[0] = "cheese"
toppings[1] = "onions"
toppings[2] = "olives"
```

Using `:=`

```go
toppings := [3]string {"cheese", "onions", "olives"}
```

Use `range` to get every value out of the array in a loop to iterate over all items:

```go
package main

import "fmt"

func main() {
    toppings := [3]string {"cheese", "onions", "olives"}

    for i, value := range toppings {
        fmt.Println(value, i)
    }
}
```

to run: `go run arrays.go`

If you have no use for the index you can modify the previous code block's `i` variable to `_`:

```go
for _, value := range toppings {
    fmt.Println(value)
}
```

### Slices

A slice is like an array, but when defining you can leave out the size.

```go
numbers := []int {1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]

// You can create a slice by defining the first index value to take through the last
numSlice1 := numbers[3:5] // [4, 5]

// If you don't supply the first index it defaults to 0
numSlice2 := numbers[:2] // [1, 2]

// If you don't supply the last index it defaults to max
numSlice3 := numbers[2:] // [3, 4, 5]

// You can create an empty slice
// method signature: make(type, length as 0, capacity)
numSlice4 := make([]int, 5, 10)

// You can copy a slice to another
copy(numSlice4, numSlice)
// numSlice4 = [1, 2, 3, 4, 5]

// Append values to the end of a slice

numSlice3 = append(numSlice3, 0, -1)

fmt.Println(numSlice3[6])
```

to run: `go run slices.go`
