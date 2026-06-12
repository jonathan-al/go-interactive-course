# Lesson 01: Go Basics

Welcome to your first Go lesson! In this module, you'll learn the fundamental building blocks of the Go programming language.

## Table of Contents

1. [Hello World](#1-hello-world)
2. [Variables](#2-variables)
3. [Basic Types](#3-basic-types)
4. [The fmt Package](#4-the-fmt-package)
5. [Constants](#5-constants)

---

## 1. Hello World

Every programming journey begins with "Hello, World!". Here's a complete Go program:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

### Breakdown

- **`package main`** — Every Go program starts with a package declaration. `main` is special: it tells Go this is an executable program (not a library).
- **`import "fmt"`** — Imports the `fmt` (format) package from Go's standard library, which provides functions for printing and reading input.
- **`func main()`** — The entry point of every Go program. Execution begins here.
- **`fmt.Println(...)`** — Prints the given text followed by a newline.

To run a Go program:

```bash
go run main.go
```

---

## 2. Variables

Go provides several ways to declare variables.

### Using the `var` keyword

```go
var name string = "Alice"
var age int = 30
var height float64 = 1.68
var active bool = true
```

You can omit the type if Go can infer it from the value:

```go
var name = "Alice"   // inferred as string
var age = 30         // inferred as int
```

### Short variable declaration (`:=`)

Inside functions, you can use the short declaration operator:

```go
name := "Alice"
age := 30
height := 1.68
active := true
```

This is the most common way to declare variables in Go. The `:=` operator declares and initializes in one step.

### Zero values

When you declare a variable without initializing it, Go assigns a **zero value**:

| Type    | Zero Value |
|---------|------------|
| `int`   | `0`        |
| `float` | `0.0`      |
| `string`| `""`       |
| `bool`  | `false`    |

```go
var count int     // 0
var message string // ""
```

### Multiple variable declaration

```go
var x, y, z int = 1, 2, 3
a, b, c := "go", 3.14, true
```

---

## 3. Basic Types

Go is a **statically typed** language. Every variable has a specific type determined at compile time.

### Numeric Types

| Type      | Description                        |
|-----------|------------------------------------|
| `int`     | Signed integer (platform-dependent size) |
| `int8`    | Signed 8-bit integer (-128 to 127) |
| `int16`   | Signed 16-bit integer              |
| `int32`   | Signed 32-bit integer              |
| `int64`   | Signed 64-bit integer              |
| `float32` | 32-bit floating-point number       |
| `float64` | 64-bit floating-point number       |

### Other Basic Types

| Type     | Description                    |
|----------|--------------------------------|
| `string` | Sequence of characters (immutable) |
| `bool`   | `true` or `false`              |

### Type Conversions

Go requires **explicit** type conversions (no implicit casting):

```go
var i int = 42
var f float64 = float64(i)   // 42.0
var u uint = uint(f)          // 42

str := "100"
// You cannot do: n := int(str)
// Use strconv.Atoi instead:
n, _ := strconv.Atoi(str)
```

### Strings

Strings in Go are sequences of bytes, encoded as UTF-8:

```go
greeting := "Hello, Go!"
length := len(greeting)    // 10
first := greeting[0]       // 72 (byte value of 'H')
sub := greeting[0:5]       // "Hello"
```

Strings are **immutable** — once created, you cannot change their contents. You must create a new string.

---

## 4. The fmt Package

The `fmt` package is essential for formatted I/O in Go.

### Printing Functions

| Function      | Description                                    |
|---------------|------------------------------------------------|
| `Println`     | Prints values separated by spaces, adds newline |
| `Print`       | Prints values separated by spaces (no newline) |
| `Printf`      | Prints formatted string using format verbs     |
| `Sprintf`     | Returns a formatted string (does not print)    |

### Format Verbs

| Verb   | Used For                  | Example                          |
|--------|---------------------------|----------------------------------|
| `%v`   | Default format            | `fmt.Printf("%v", 42)` → `42`   |
| `%d`   | Integer (decimal)         | `fmt.Printf("%d", 42)` → `42`   |
| `%f`   | Float                     | `fmt.Printf("%f", 3.14)` → `3.140000` |
| `%s`   | String                    | `fmt.Printf("%s", "hi")` → `hi` |
| `%t`   | Boolean                   | `fmt.Printf("%t", true)` → `true` |
| `%T`   | Type of value             | `fmt.Printf("%T", 42)` → `int`  |
| `%%`   | Literal percent sign      | `fmt.Printf("100%%")` → `100%`  |

### Controlling Float Precision

```go
pi := 3.14159
fmt.Printf("%.2f", pi)   // 3.14
fmt.Printf("%.4f", pi)   // 3.1416
```

### String Formatting with Sprintf

`Sprintf` formats a string and returns it instead of printing:

```go
name := "Alice"
age := 30
msg := fmt.Sprintf("%s is %d years old", name, age)
// msg = "Alice is 30 years old"
```

### Width and Alignment

```go
fmt.Printf("%-10s|", "left")    // "left      |"  (left-aligned, width 10)
fmt.Printf("%10s|", "right")    // "     right|"  (right-aligned, width 10)
```

---

## 5. Constants

Constants are immutable values declared with the `const` keyword:

```go
const Pi = 3.14159
const Greeting = "Hello"
const MaxRetries = 3
```

Constants can only be strings, numbers, or booleans. Once declared, they cannot be changed.

### Typed vs Untyped Constants

```go
const untypedConst = 42              // untyped — adapts to context
const typedConst int = 42            // typed — always int
```

Untyped constants are more flexible because Go will automatically convert them to the needed type.

### Grouped Constants

```go
const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"
)
```

### The `iota` Enumerator

`iota` is a special identifier that generates successive integer values starting from 0 in a `const` block:

```go
const (
	Sunday    = iota  // 0
	Monday            // 1
	Tuesday           // 2
	Wednesday         // 3
	Thursday          // 4
	Friday            // 5
	Saturday          // 6
)
```

`iota` resets to 0 whenever a new `const` block begins. Each subsequent constant in the same block increments `iota` by 1.

#### Practical example: HTTP Status Codes

```go
const (
	StatusOK       = iota + 200  // 200
	StatusCreated                // 201
	StatusAccepted               // 202
)
```

---

## Summary

In this lesson you learned:

- How to write and run a basic Go program
- Variable declaration with `var` and `:=`
- Go's basic types: `int`, `float64`, `string`, `bool`
- How to use `fmt.Println`, `fmt.Printf`, and `fmt.Sprintf`
- Format verbs like `%v`, `%d`, `%f`, `%s`, `%T`
- Constants and the `iota` enumerator

Now complete the exercises to reinforce what you've learned!
