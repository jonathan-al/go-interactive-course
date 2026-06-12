# Lesson 6: Error Handling

Go takes a unique approach to error handling. Instead of exceptions, Go uses explicit return values to signal failures. This makes error handling visible, predictable, and straightforward.

## The error Interface

In Go, errors are represented by the built-in `error` interface:

```go
type error interface {
    Error() string
}
```

Any type that implements an `Error() string` method satisfies this interface. Functions that can fail conventionally return an `error` as their last return value:

```go
func readFile(name string) ([]byte, error) {
    // ...
}
```

## Returning Errors

Use `errors.New` to create a simple error from a string:

```go
import "errors"

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

When calling a function that returns an error, always check the error value:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

Use `fmt.Errorf` when you need to include dynamic information in the error message:

```go
func getAge(age int) (string, error) {
    if age < 0 {
        return "", fmt.Errorf("invalid age: %d", age)
    }
    return fmt.Sprintf("Age is %d", age), nil
}
```

## Wrapping Errors

You can wrap errors to add context while preserving the original error using the `%w` verb:

```go
func readConfig(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("reading config %s: %w", path, err)
    }
    return data, nil
}
```

Use `errors.Is` to check if an error matches a specific value, and `errors.As` to check if an error matches a specific type. Both unwrap errors automatically:

```go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("file does not exist")
}

var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("path:", pathErr.Path)
}
```

## Custom Error Types

You can define your own error types by creating a struct that implements the `error` interface. This is useful when you need to carry additional information with the error:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}
```

Use it in a function:

```go
func validateEmail(email string) error {
    if email == "" {
        return &ValidationError{Field: "email", Message: "cannot be empty"}
    }
    return nil
}
```

Check for custom error types using type assertions or `errors.As`:

```go
err := validateEmail("")
if err != nil {
    var ve *ValidationError
    if errors.As(err, &ve) {
        fmt.Printf("Field: %s, Message: %s\n", ve.Field, ve.Message)
    }
}
```

## defer

The `defer` statement schedules a function call to be executed just before the surrounding function returns. Deferred calls are executed in **last-in, first-out** (LIFO) order:

```go
func main() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")
}
// Output:
// third
// second
// first
```

`defer` is commonly used for cleanup tasks like closing files or releasing resources:

```go
func readFile(name string) ([]byte, error) {
    f, err := os.Open(name)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    return io.ReadAll(f)
}
```

## panic and recover

`panic` stops the normal execution of the program and begins unwinding the stack. It should be used only for truly exceptional situations, such as programmer errors:

```go
func mustParse(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(fmt.Sprintf("cannot parse %q as int", s))
    }
    return n
}
```

`recover` allows you to regain control after a panic. It only works inside a deferred function:

```go
func safeCall() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    panic("something went wrong")
}
```

When `safeCall()` runs, the panic is caught by the deferred recover function, and the program continues normally.

## Key Takeaways

- Go uses explicit error return values instead of exceptions.
- Always check errors returned by functions.
- Use `errors.New` and `fmt.Errorf` to create errors.
- Define custom error types when you need to carry extra information.
- Use `defer` for cleanup operations; deferred calls run in LIFO order.
- Use `panic` sparingly, and `recover` inside deferred functions to handle panics gracefully.
