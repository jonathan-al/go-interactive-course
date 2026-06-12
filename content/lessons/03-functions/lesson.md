# Lesson 3: Functions

Functions are the building blocks of any Go program. They let you group reusable logic, accept inputs, and return outputs.

## Basic Functions

A function declaration uses the `func` keyword, followed by a name, parameters, and an optional return type:

```go
func add(a int, b int) int {
    return a + b
}
```

When consecutive parameters share the same type, you can shorten the declaration:

```go
func add(a, b int) int {
    return a + b
}
```

Call a function by its name and pass the required arguments:

```go
result := add(5, 3)
fmt.Println(result) // 8
```

## Multiple Return Values

Go functions can return more than one value. This is commonly used to return a result alongside an error:

```go
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}
```

Capture multiple returns with a comma-separated assignment:

```go
q, r := divide(10, 3)
fmt.Printf("10 / 3 = %d remainder %d\n", q, r) // 10 / 3 = 3 remainder 1
```

You can also name return values, which acts as a declaration and lets you use a bare `return`:

```go
func divide(a, b int) (quotient int, remainder int) {
    quotient = a / b
    remainder = a % b
    return
}
```

## Variadic Functions

A variadic function accepts a variable number of arguments of the same type. Use `...` before the type:

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

Call it with any number of arguments:

```go
fmt.Println(sum(1, 2, 3))       // 6
fmt.Println(sum(10, 20, 30, 40)) // 100
```

You can also pass a slice using the `...` operator:

```go
numbers := []int{1, 2, 3, 4, 5}
fmt.Println(sum(numbers...)) // 15
```

## Closures

A closure is a function value that references variables from outside its body. The function retains access to those variables even after the outer function has returned:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

Each call to `counter()` creates a new, independent closure:

```go
c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
fmt.Println(c()) // 3
```

Closures are useful for creating stateful functions, generators, and encapsulating behavior without global variables.

## Key Takeaways

- Functions are declared with `func`, parameters, and optional return types.
- Go supports multiple return values, commonly used for result + error patterns.
- Variadic functions (`...`) accept a flexible number of arguments.
- Closures capture variables from their surrounding scope, enabling stateful behavior.
