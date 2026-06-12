# Lesson 05: Structs, Methods, Interfaces, and Embedding

## Structs

A struct is a composite data type that groups together variables (fields) under a single name. Each field has a name and a type.

```go
type Person struct {
    Name string
    Age  int
}
```

### Creating Struct Instances

```go
// Using field names (preferred)
p1 := Person{Name: "Alice", Age: 30}

// Positional (must match field order)
p2 := Person{"Bob", 25}

// Zero-value struct
var p3 Person // {"" 0}

// Pointer to struct
p4 := &Person{Name: "Charlie", Age: 35}
```

### Accessing and Modifying Fields

```go
fmt.Println(p1.Name) // Alice
p1.Age = 31          // modify a field
```

### Anonymous Structs

You can create a struct without defining a named type:

```go
point := struct{ X, Y int }{1, 2}
```

## Methods

Methods are functions with a special **receiver** argument. The receiver appears between `func` and the method name.

### Value Receivers

A value receiver operates on a copy of the struct. Changes inside the method do not affect the original.

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

rect := Rectangle{Width: 5, Height: 3}
fmt.Println(rect.Area()) // 15
```

### Pointer Receivers

A pointer receiver can modify the original struct. Use pointer receivers when the method needs to mutate the receiver or when the struct is large and you want to avoid copying.

```go
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

rect := Rectangle{Width: 5, Height: 3}
rect.Scale(2)
fmt.Println(rect.Width) // 10
```

Go automatically passes a pointer when calling a pointer-receiver method on a value, so `rect.Scale(2)` works even though `rect` is not a pointer.

### Methods on Non-Struct Types

You can define methods on any type you define in your package:

```go
type MyString string

func (s MyString) Shout() string {
    return strings.ToUpper(string(s)) + "!"
}
```

## Interfaces

An interface is a set of method signatures. A type **implements** an interface by implementing all of its methods. There is no explicit declaration of intent — interfaces are satisfied **implicitly**.

### Defining and Implementing Interfaces

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Both `Circle` and `Rectangle` implement `Shape` without any explicit declaration.

### Using Interfaces

```go
func printArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

printArea(Circle{Radius: 5})       // Area: 78.54
printArea(Rectangle{Width: 4, Height: 6}) // Area: 24.00
```

### The Empty Interface

The empty interface `interface{}` (or `any` in Go 1.18+) has no methods, so every type satisfies it:

```go
var anything any
anything = 42
anything = "hello"
```

### Type Assertions and Type Switches

```go
var s Shape = Circle{Radius: 5}

// Type assertion
c, ok := s.(Circle)
if ok {
    fmt.Println(c.Radius)
}

// Type switch
switch v := s.(type) {
case Circle:
    fmt.Println("Circle with radius", v.Radius)
case Rectangle:
    fmt.Println("Rectangle", v.Width, "x", v.Height)
default:
    fmt.Println("Unknown shape")
}
```

## Embedding (Composition)

Go does not have classical inheritance. Instead, it supports **composition** through embedding. A struct can embed another struct, gaining access to its fields and methods as if they were its own.

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "..."
}

type Dog struct {
    Animal // embedded struct
    Breed  string
}

d := Dog{
    Animal: Animal{Name: "Buddy"},
    Breed:  "Labrador",
}

fmt.Println(d.Name)    // Buddy (promoted from Animal)
fmt.Println(d.Speak()) // ... (promoted method)
```

### Method Overriding with Embedding

You can define a method on the outer struct that shadows the embedded struct's method:

```go
func (d Dog) Speak() string {
    return "Woof!"
}

fmt.Println(d.Speak()) // Woof!
```

### Interface Embedding

Interfaces can embed other interfaces:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

## Summary

| Concept | Key Points |
|---------|-----------|
| Structs | Group fields into a custom type, created with literals or `new`/`&` |
| Methods | Functions with a receiver; value (copy) or pointer (mutate) |
| Interfaces | Implicit satisfaction, enables polymorphism |
| Embedding | Composition over inheritance, promotes fields and methods |
