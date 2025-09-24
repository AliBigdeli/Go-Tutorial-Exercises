
# Learning Go (Golang) 🚀

This repository contains my notes and practice code as I learn the Go programming language.  
It starts from installation and setup, and gradually covers basic concepts like variables, data types, and math operations.



## 📑 Table of Contents
- [Learning Go (Golang) 🚀](#learning-go-golang-)
  - [📑 Table of Contents](#-table-of-contents)
  - [Installation](#installation)
  - [Setting Up a Project](#setting-up-a-project)
    - [Initialize a Module](#initialize-a-module)
    - [Run Go Code](#run-go-code)
  - [Basic Go Program](#basic-go-program)
  - [Data Types](#data-types)
    - [Integers](#integers)
    - [Unsigned Integers](#unsigned-integers)
    - [Floats](#floats)
  - [Math Operations](#math-operations)
    - [Strings](#strings)
      - [Examples](#examples)
    - [Booleans](#booleans)
      - [Examples](#examples-1)
    - [Runes](#runes)
      - [Examples](#examples-2)
    - [Variable Assignment](#variable-assignment)
      - [1. Declare with type and value](#1-declare-with-type-and-value)
      - [2. Declare without type (type is inferred)](#2-declare-without-type-type-is-inferred)
      - [3. Declare without value (zero value assigned)](#3-declare-without-value-zero-value-assigned)
      - [4. Constants](#4-constants)
      - [5. Shorthand declaration (`:=`) ⚡ *(later)*](#5-shorthand-declaration---later)
      - [5. Multiple Assignment](#5-multiple-assignment)
    - [Type Checking](#type-checking)
      - [Examples](#examples-3)
    - [If Statements](#if-statements)
      - [Examples](#examples-4)
    - [Logical Operators](#logical-operators)
      - [Examples](#examples-5)
    - [Loops (for)](#loops-for)
      - [1. Classic for loop](#1-classic-for-loop)
      - [2. While-style loop](#2-while-style-loop)
      - [3. Infinite loop](#3-infinite-loop)
      - [4. Range loop (iterating over arrays, slices, strings, maps)](#4-range-loop-iterating-over-arrays-slices-strings-maps)
    - [Error Handling and Panic/Recover](#error-handling-and-panicrecover)
      - [1. Handling errors with `error`](#1-handling-errors-with-error)
      - [2. Panic and Recover](#2-panic-and-recover)
    - [Switch Case](#switch-case)
      - [Examples](#examples-6)
    - [Functions](#functions)
      - [Examples](#examples-7)
    - [Arrays and Basic Operations](#arrays-and-basic-operations)
      - [Examples](#examples-8)
    - [Maps](#maps)
      - [Examples](#examples-9)
    - [In-place Operators](#in-place-operators)
      - [Examples](#examples-10)
    - [Structs and Interfaces](#structs-and-interfaces)
      - [1. Structs](#1-structs)
      - [2. Interfaces](#2-interfaces)
    - [Pointers](#pointers)
      - [Examples](#examples-11)
    - [Goroutines](#goroutines)
      - [Examples](#examples-12)
    - [Go Package Management](#go-package-management)
      - [1. Installing a Package](#1-installing-a-package)
      - [2. Importing the Package](#2-importing-the-package)
      - [3. Updating a Package](#3-updating-a-package)
      - [4. Removing a Package](#4-removing-a-package)
      - [5. Checking Modules](#5-checking-modules)



## Installation

Download and install Go from the official site:  
👉 [https://go.dev/dl/](https://go.dev/dl/)

Verify installation:

```bash
go version
````



## Setting Up a Project

### Initialize a Module

Go uses modules to manage code and dependencies.

```bash
go mod init myproject
```

This creates a `go.mod` file in your project directory.

### Run Go Code

Run a Go file directly:

```bash
go run main.go
```

Or build an executable:

```bash
go build main.go
./main
```



## Basic Go Program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```



## Data Types

### Integers

* `int` (size depends on system: 32-bit or 64-bit)
* `int8` : -128 to 127
* `int16`: -32,768 to 32,767
* `int32`: -2,147,483,648 to 2,147,483,647
* `int64`: -9,223,372,036,854,775,808 to +9,223,372,036,854,775,807

### Unsigned Integers

* `uint` : positive only
* `uint8` : 0 to 255
* `uint16`: 0 to 65,535
* `uint32`: 0 to 4,294,967,295
* `uint64`: 0 to 18,446,744,073,709,551,615

### Floats

* `float32` : \~ ±3.4e38
* `float64` : \~ ±1.7e308



## Math Operations

| Operator        | Name             | Example          | Result | Notes                              |
| --------------- | ---------------- | ---------------- | ------ | ---------------------------------- |
| `+`             | Addition         | `10 + 5`         | `15`   | Works for integers & floats        |
| `-`             | Subtraction      | `10 - 5`         | `5`    | Works for integers & floats        |
| `*`             | Multiplication   | `10 * 5`         | `50`   | Works for integers & floats        |
| `/` (int)       | Division (int)   | `7 / 2`          | `3`    | Drops decimals                     |
| `/` (float)     | Division (float) | `7.0 / 2.0`      | `3.5`  | At least one operand must be float |
| `%`             | Modulo           | `10 % 3`         | `1`    | Only for integers                  |
| `math.Pow(a,b)` | Power            | `math.Pow(2, 3)` | `8.0`  | Returns `float64`                  |
| `math.Sqrt(x)`  | Square Root      | `math.Sqrt(16)`  | `4.0`  | Returns `float64`                  |
| `math.Abs(x)`   | Absolute Value   | `math.Abs(-5)`   | `5`    | Returns `float64`                  |

### Strings
Strings in Go are sequences of characters enclosed in double quotes (`" "`).  
You can concatenate strings with `+` and format them using `fmt.Sprintf`.

#### Examples

```go
package main

import (
    "fmt"
)

func main() {
    // Simple string
    var str1 String = "Hello"
    var str2 String = "World"

    // Concatenation
    var result = str1 + " " + str2
    fmt.Println("Concatenated:", result)

    // String length
    fmt.Println("Length of str1:", len(str1))

    // Formatting strings
    name := "Alice"
    age := 30
    formatted := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
    fmt.Println("Formatted:", formatted)
}
```

### Booleans
Booleans in Go represent truth values: `true` or `false`.  
They are commonly used in conditions and logical expressions.

#### Examples

```go
package main

import (
    "fmt"
)

func main() {
    // Boolean values
    var isActive bool = true
    var isDone bool = false
    fmt.Println("isActive:", isActive, "isDone:", isDone)

    // Boolean expressions
    var x int = 10
    var y int = 20
    fmt.Println("x < y:", x < y)   // true
    fmt.Println("x > y:", x > y)   // false
    fmt.Println("x == y:", x == y) // false
    fmt.Println("x != y:", x != y) // true

    // Logical operators
    var a bool = true
    var b bool = false
    fmt.Println("a && b:", a && b) // AND -> false
    fmt.Println("a || b:", a || b) // OR  -> true
    fmt.Println("!a:", !a)         // NOT -> false
}
```

### Runes
A rune in Go represents a Unicode code point (a character).  
It is an alias for `int32` and allows you to work with individual characters in a string.

#### Examples

```go
package main

import (
    "fmt"
)

func main() {
    // Declaring runes
    var r rune = 'A'
    fmt.Println("Rune:", r)          // prints numeric code point (65)
    fmt.Printf("As character: %c\n", r) // prints 'A'

    // Unicode rune
    var heart rune = '❤'
    fmt.Println("Unicode Rune:", heart)
    fmt.Printf("As character: %c\n", heart)

    // Iterating over a string (each character is a rune)
    var text string = "Go!"
    for i, char := range text {
        fmt.Printf("Index: %d, Rune: %c, CodePoint: %U\n", i, char, char)
    }
}
```

### Variable Assignment
Go provides multiple ways to declare and assign values to variables.

#### 1. Declare with type and value
You explicitly specify the type:

```go
var age int = 30
var name string = "Alice"
````

#### 2. Declare without type (type is inferred)

The compiler infers the type from the value:

```go
var city = "Paris"   // string
var score = 100      // int
var pi = 3.14        // float64 (default)
```

#### 3. Declare without value (zero value assigned)

If no value is given, Go assigns a *zero value*:

```go
var number int       // 0
var text string      // ""
var flag bool        // false
```

#### 4. Constants

Constants are declared with `const` and cannot be changed:

```go
const Pi = 3.14159
const Welcome = "Hello, Go"
```

#### 5. Shorthand declaration (`:=`) ⚡ *(later)*

Inside functions, you can use `:=` to declare and assign in one step.
(Not needed for now, but useful to know.)

```go
age := 25
language := "Go"
```


#### 5. Multiple Assignment

Go allows assigning values to multiple variables in a single line.

```go
package main

import (
    "fmt"
)

func main() {
    var a, b, c int = 10, 20, 30
    var x, y string = "Go", "Lang"
    
    fmt.Println("a:", a, "b:", b, "c:", c)
    fmt.Println("x:", x, "y:", y)
    
    // Swapping variables
    a, b = b, a
    fmt.Println("After swap a:", a, "b:", b)
}
```


### Type Checking
Go is a statically typed language, so every variable has a type that is known at compile time.  
You can check or print the type of a variable using the `fmt.Printf` function with the `%T` verb.

#### Examples

```go
package main

import (
    "fmt"
)

func main() {
    var age int = 30
    var name string = "Alice"
    var pi float64 = 3.14
    var isActive bool = true
    var letter rune = 'A'

    fmt.Printf("Type of age: %T\n", age)       // int
    fmt.Printf("Type of name: %T\n", name)     // string
    fmt.Printf("Type of pi: %T\n", pi)         // float64
    fmt.Printf("Type of isActive: %T\n", isActive) // bool
    fmt.Printf("Type of letter: %T\n", letter) // rune (int32)
}
```


### If Statements
`if` statements are used to execute code based on a condition.  
The condition must evaluate to a boolean (`true` or `false`).

#### Examples

```go
package main

import (
    "fmt"
)

func main() {
    var age int = 20

    if age >= 18 {
        fmt.Println("You are an adult.")
    }

    if age < 18 {
        fmt.Println("You are a minor.")
    } else {
        fmt.Println("You are not a minor.")
    }

    // If-else-if ladder
    if age < 13 {
        fmt.Println("Child")
    } else if age < 18 {
        fmt.Println("Teenager")
    } else if age < 60 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Senior")
    }
}
```

### Logical Operators
Logical operators are used with boolean values (`true` or `false`) to build complex conditions.

| Operator | Name | Description |
|----------|------|-------------|
| `&&`     | AND  | True if both operands are true |
| `\|\|`   | OR   | True if at least one operand is true |
| `!`      | NOT  | Inverts the boolean value |

#### Examples

```go
package main

import "fmt"

func main() {
    var a bool = true
    var b bool = false

    fmt.Println("a && b:", a && b) // false
    fmt.Println("a || b:", a || b) // true
    fmt.Println("!a:", !a)         // false
    fmt.Println("!b:", !b)         // true

    // Combining conditions
    var x int = 10
    var y int = 20
    var z int = 5

    fmt.Println("x < y && x > z:", x < y && x > z) // true
    fmt.Println("x > y || x > z:", x > y || x > z) // true
}
```

### Loops (for)

Go uses `for` loops for all iteration.  
There is no `while` or `do-while`; `for` can mimic all loop types.

#### 1. Classic for loop
```go
package main

import "fmt"

func main() {
    var i int
    for i = 0; i < 5; i++ {
        fmt.Println("i =", i)
    }
}
````

#### 2. While-style loop

Use a `for` with only a condition:

```go
package main

import "fmt"

func main() {
    var i int = 0
    for i < 5 {
        fmt.Println("i =", i)
        i++
    }
}
```

#### 3. Infinite loop

Omit all components; useful with `break` or `return`:

```go
package main

import "fmt"

func main() {
    var i int = 0
    for {
        if i >= 5 {
            break
        }
        fmt.Println("i =", i)
        i++
    }
}
```

#### 4. Range loop (iterating over arrays, slices, strings, maps)

```go
package main

import "fmt"

func main() {
    var numbers = [5]int{10, 20, 30, 40, 50}
    
    for index, value := range numbers {
        fmt.Println("Index:", index, "Value:", value)
    }
    
    var text string = "Go"
    for index, char := range text {
        fmt.Printf("Index: %d, Rune: %c\n", index, char)
    }
}
```
### Error Handling and Panic/Recover

Go handles errors using explicit `error` values rather than exceptions.  
You can also handle unexpected runtime issues using `panic` and `recover`.

#### 1. Handling errors with `error`
Many functions return an `error` as the last return value.  
Always check the error before proceeding.

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
````

#### 2. Panic and Recover

`panic` stops normal execution; `recover` can handle a panic in a deferred function.

```go
package main

import "fmt"

func main() {
    var a int = 10
    var b int = 0
    var result int

    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    if b == 0 {
        panic("division by zero")
    } else {
        result = a / b
        fmt.Println("Result:", result)
    }

    fmt.Println("Program continues...")
}

```

### Switch Case
The `switch` statement allows you to execute code based on multiple possible values of a variable.  
It is often cleaner than multiple `if-else-if` statements.

#### Examples

```go
package main

import "fmt"

func main() {
    var day int = 3
    var dayName string

    switch day {
    case 1:
        dayName = "Monday"
    case 2:
        dayName = "Tuesday"
    case 3:
        dayName = "Wednesday"
    case 4:
        dayName = "Thursday"
    case 5:
        dayName = "Friday"
    case 6:
        dayName = "Saturday"
    case 7:
        dayName = "Sunday"
    default:
        dayName = "Invalid day"
    }

    fmt.Println("Day:", dayName)

    // Switch with multiple values in one case
    var score int = 85
    switch {
    case score >= 90:
        fmt.Println("Grade: A")
    case score >= 75:
        fmt.Println("Grade: B")
    case score >= 50:
        fmt.Println("Grade: C")
    default:
        fmt.Println("Grade: F")
    }
}
```

### Functions
Functions in Go are blocks of code that perform a specific task.  
They can accept parameters and return values.

#### Examples

```go
package main

import "fmt"

func main() {
    // Function with parameters and return value
    var sumResult int = add(10, 20)
    fmt.Println("Sum:", sumResult)

    // Function without parameters
    greet()

    // Function with multiple return values
    var a int
    var b int
    a, b = swap(5, 10)
    fmt.Println("After swap a:", a, "b:", b)
}

// Function that takes two integers and returns the sum
func add(x int, y int) int {
    return x + y
}

// Function that prints a greeting
func greet() {
    fmt.Println("Hello from a function!")
}

// Function that returns two integers (swap example)
func swap(x int, y int) (int, int) {
    return y, x
}
```

### Arrays and Basic Operations
Arrays in Go are fixed-size collections of elements of the same type.  
For dynamic operations like add/remove, slices are used.

#### Examples

```go
package main

import "fmt"

func main() {
    // Declare an array
    var numbers [3]int = [3]int{1, 2, 3}
    fmt.Println("Array:", numbers)

    // Update element
    numbers[0] = 10
    fmt.Println("After update:", numbers)

    // Sum all elements
    var sum int = 0
    var i int
    for i = 0; i < len(numbers); i++ {
        sum = sum + numbers[i]
    }
    fmt.Println("Sum:", sum)

    // Convert array to slice for dynamic operations
    var numsSlice []int = numbers[:]

    // Add element (append)
    numsSlice = append(numsSlice, 4)
    fmt.Println("After append:", numsSlice)

    // Pop element (remove last)
    var last int = numsSlice[len(numsSlice)-1]
    numsSlice = numsSlice[:len(numsSlice)-1]
    fmt.Println("After pop:", numsSlice, "Popped value:", last)

    // Iterate using range
    var index int
    var value int
    for index, value = range numsSlice {
        fmt.Println("Index:", index, "Value:", value)
    }
}
```


### Maps
Maps in Go are collections of key-value pairs.  
Keys are unique, and values are accessed using the key.

#### Examples

```go
package main

import "fmt"

func main() {
    // Declare and initialize a map
    var person map[string]string = map[string]string{
        "name": "Alice",
        "city": "Paris",
    }
    fmt.Println("Person map:", person)

    // Access value by key
    fmt.Println("Name:", person["name"])

    // Update value
    person["city"] = "London"
    fmt.Println("Updated map:", person)

    // Add new key-value pair
    person["country"] = "UK"
    fmt.Println("After adding country:", person)

    // Delete a key
    delete(person, "name")
    fmt.Println("After deleting name:", person)

    // Iterate over map
    var key string
    var value string
    for key, value = range person {
        fmt.Println("Key:", key, "Value:", value)
    }
}
```


### In-place Operators
Go supports operators that modify the value of a variable directly.  
Common ones include `++`, `--`, `+=`, `-=`, `*=`, `/=`, and `%=`.

#### Examples

```go
package main

import "fmt"

func main() {
    var a int = 10
    fmt.Println("Initial a:", a)

    // Increment and decrement
    a++
    fmt.Println("After a++:", a)
    a--
    fmt.Println("After a--:", a)

    // Add, subtract, multiply, divide in-place
    a += 5
    fmt.Println("After a += 5:", a)
    a -= 3
    fmt.Println("After a -= 3:", a)
    a *= 2
    fmt.Println("After a *= 2:", a)
    a /= 4
    fmt.Println("After a /= 4:", a)
    a %= 3
    fmt.Println("After a %= 3:", a)
}
```


### Structs and Interfaces

#### 1. Structs
Structs are custom data types that group together fields (variables) under one name.

```go
package main

import "fmt"

func main() {
    // Define a struct
    type Person struct {
        name string
        age  int
    }

    // Create a struct instance
    var p Person
    p.name = "Alice"
    p.age = 30
    fmt.Println("Person:", p)
    fmt.Println("Name:", p.name, "Age:", p.age)

    // Initialize struct in one line
    var p2 Person = Person{name: "Bob", age: 25}
    fmt.Println("Person 2:", p2)
}
````

#### 2. Interfaces

Interfaces define a set of methods. Any type that implements these methods satisfies the interface.

```go
package main

import "fmt"

// Define an interface
type Shape interface {
    Area() float64
}

// Struct implementing the interface
type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
    var r Rectangle = Rectangle{width: 5, height: 3}
    var s Shape = r // Rectangle implements Shape
    fmt.Println("Area of rectangle:", s.Area())
}
```

### Pointers
Pointers store the memory address of a variable.  
You can access or modify the value using the pointer.

#### Examples

```go
package main

import "fmt"

func main() {
    // Declare a variable
    var a int = 10
    fmt.Println("Value of a:", a)

    // Pointer to a
    var p *int = &a
    fmt.Println("Pointer p:", p)        // memory address of a
    fmt.Println("Value via p:", *p)     // dereference to get value

    // Modify value via pointer
    *p = 20
    fmt.Println("New value of a:", a)

    // Another pointer example
    var b int = 5
    var q *int = &b
    fmt.Println("Before:", b)
    *q += 10
    fmt.Println("After:", b)
}
```


### Goroutines
Goroutines are lightweight threads in Go that allow concurrent execution.  
Use the `go` keyword before a function call to run it as a goroutine.

#### Examples

```go
package main

import (
    "fmt"
    "time"
)

// Function to run as goroutine
func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    // Start a goroutine
    go sayHello()

    // Main function continues
    fmt.Println("Hello from main function")

    // Give goroutine time to finish
    time.Sleep(1 * time.Second)
}
```


### Go Package Management
Go uses modules to manage external packages. You can install, update, and remove packages easily using `go` commands.

#### 1. Installing a Package
Use `go get` to download and install a package:

```bash
go get github.com/user/package
````

This adds the package to your `go.mod` file.

#### 2. Importing the Package

After installing, import it in your code:

```go
package main

import (
    "fmt"
    "github.com/user/package" // example package
)

func main() {
    fmt.Println("Using external package")
}
```

#### 3. Updating a Package

Update to the latest version with:

```bash
go get -u github.com/user/package
```

#### 4. Removing a Package

Remove unused packages and tidy up `go.mod` and `go.sum`:

```bash
go mod tidy
```

#### 5. Checking Modules

List all modules and versions used in the project:

```bash
go list -m all
```

Notes:

* `go.mod` tracks all dependencies.
* `go.sum` ensures package integrity.
* `go get` automatically downloads packages if not already installed.
* Always run `go mod tidy` to clean unused dependencies.

