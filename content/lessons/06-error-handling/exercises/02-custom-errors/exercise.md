# Exercise 2: Custom Error Types

Create a custom error type called `NegativeError` that stores a `Value int` field and implements the `error` interface. The `Error()` method should return:

```
negative value: <value>
```

Write a function called `sqrt` that takes a `float64` and returns `(float64, error)`. If the input is negative, return a `*NegativeError` with the integer part of the value. Otherwise, return `math.Sqrt` of the value and `nil`.

In `main`, call `sqrt(-4)` and use `errors.As` to check for a `*NegativeError`, printing the field. Then call `sqrt(16)` and print the result.

Expected output:

```
NegativeError: negative value: -4
sqrt(16) = 4.000000
```
