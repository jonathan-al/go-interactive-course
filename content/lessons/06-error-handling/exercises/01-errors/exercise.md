# Exercise 1: Error Handling

Write a function called `divide` that takes two `int` parameters (`a` and `b`) and returns `(int, error)`.

If `b` is zero, return an error with the message `"cannot divide by zero"`. Otherwise, return the result of `a / b` and `nil` for the error.

In `main`, call `divide(10, 0)` and handle the error by printing it. Then call `divide(10, 3)` and print the result.

Expected output:

```
Error: cannot divide by zero
10 / 3 = 3
```
