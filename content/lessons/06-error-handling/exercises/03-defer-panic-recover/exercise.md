# Exercise 3: Defer, Panic, and Recover

Write a function called `safeDivide` that takes two `int` parameters (`a` and `b`) and returns `result int`. Use `defer` and `recover` to catch a panic caused by dividing by zero.

Inside `safeDivide`:
1. Set up a deferred function that uses `recover()` to catch any panic and prints `"Recovered: <panic value>"`.
2. If `b` is zero, call `panic("division by zero")`.
3. Otherwise, return `a / b`.

In `main`:
1. Call `safeDivide(10, 0)` and print the returned value.
2. Call `safeDivide(10, 2)` and print the returned value.

Expected output:

```
Recovered: division by zero
10 / 0 = 0
10 / 2 = 5
```
