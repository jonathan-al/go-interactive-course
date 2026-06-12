# Exercise 1: Goroutines

## Objective

Launch goroutines to compute the square of numbers 1 through 5, then collect and print the results **in order** using a channel.

## Instructions

Open `starter.go`. You'll find a `Result` struct and a `main` function with several TODOs.

Your tasks:

1. Create a channel of type `Result` with a buffer size of 5.
2. Launch 5 goroutines (for numbers 1–5). Each goroutine should compute `n * n` and send a `Result{ID: n, Value: n*n}` to the channel.
3. Receive all 5 results from the channel and store them in a slice, placing each result at index `result.ID - 1`.
4. Print each result in order using `fmt.Printf("ID: %d, Value: %d\n", ...)`.

## Expected Output

```
ID: 1, Value: 1
ID: 2, Value: 4
ID: 3, Value: 9
ID: 4, Value: 16
ID: 5, Value: 25
```

## Hints

- Use `go func(n int) { ... }(i)` to pass the loop variable into the goroutine (avoids closure capture issues).
- A buffered channel of size 5 lets all goroutines send without blocking.
- Store results in a pre-allocated slice: `results := make([]Result, 5)`.
