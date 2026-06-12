# Control Flow in Go

Control flow statements allow your program to make decisions and repeat actions. In this lesson, you'll learn about **if/else**, **switch**, and **for** loops in Go.

---

## 1. If / Else

Go's `if` statement evaluates a condition and executes a block of code when the condition is `true`.

```go
age := 18

if age >= 18 {
    fmt.Println("You are an adult")
}
```

### If / Else

Add an `else` block for the alternative path:

```go
temperature := 25

if temperature > 30 {
    fmt.Println("It's hot outside")
} else {
    fmt.Println("The weather is nice")
}
```

### If / Else If / Else

Chain multiple conditions with `else if`:

```go
score := 85

if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else if score >= 70 {
    fmt.Println("Grade: C")
} else {
    fmt.Println("Grade: F")
}
```

### Short Statement in If

Go allows a short initialization statement before the condition:

```go
if num := 10; num > 0 {
    fmt.Println("Positive number:", num)
}
// num is not accessible here
```

The variable `num` is only in scope within the `if` block.

---

## 2. Switch

A `switch` statement is a cleaner way to compare a value against multiple cases.

```go
day := 3

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
default:
    fmt.Println("Other day")
}
```

### Multiple Values in a Case

A single case can match multiple values:

```go
fruit := "apple"

switch fruit {
case "apple", "banana":
    fmt.Println("Common fruit")
case "mango", "papaya":
    fmt.Println("Tropical fruit")
default:
    fmt.Println("Unknown fruit")
}
```

### Switch Without a Condition

A `switch` without a condition acts like a chain of `if/else if` statements:

```go
score := 85

switch {
case score >= 90:
    fmt.Println("Excellent")
case score >= 80:
    fmt.Println("Great job")
case score >= 70:
    fmt.Println("Good")
default:
    fmt.Println("Keep practicing")
}
```

### Fallthrough

By default, Go does **not** fall through to the next case (unlike C). Use `fallthrough` explicitly when needed:

```go
num := 1

switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("This also prints")
default:
    fmt.Println("Default")
}
```

---

## 3. For Loops

Go has only one looping construct: `for`. It covers what other languages do with `while`, `for`, and `for each`.

### Classic For Loop

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

### While-style Loop

Omit the init and post statements to get a `while` loop:

```go
count := 0
for count < 3 {
    fmt.Println("Count:", count)
    count++
}
```

### Infinite Loop

Omit everything for an infinite loop (use `break` to exit):

```go
for {
    fmt.Println("This runs forever")
    break // exits the loop
}
```

### For Range

Use `for ... range` to iterate over slices, arrays, maps, strings, and more.

**Slice:**

```go
fruits := []string{"apple", "banana", "cherry"}

for index, value := range fruits {
    fmt.Printf("%d: %s\n", index, value)
}
```

**String (iterates over runes):**

```go
for i, ch := range "Go!" {
    fmt.Printf("index %d: %c\n", i, ch)
}
```

**Blank identifier:**

Use `_` when you don't need the index or value:

```go
for _, value := range fruits {
    fmt.Println(value)
}
```

---

## Summary

| Concept | Syntax |
|---|---|
| If/Else | `if condition { } else { }` |
| Short init | `if x := 5; x > 0 { }` |
| Switch | `switch val { case 1: ... default: ... }` |
| Classic for | `for i := 0; i < n; i++ { }` |
| While-style | `for condition { }` |
| For range | `for i, v := range slice { }` |

Now try the exercises to practice what you've learned!
