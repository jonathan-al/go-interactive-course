# Exercise 4: WaitGroup

## Objective

Use `sync.WaitGroup` to wait for multiple goroutines, collect their results via a channel, and print them **in order**.

## Instructions

Open `starter.go`. You'll find a `Result` struct and a `main` function with several TODOs.

Your tasks:

1. Declare a `sync.WaitGroup`.
2. Create a buffered channel of type `Result` with capacity 5.
3. Launch 5 goroutines (for numbers 1–5). Each goroutine should:
   - Call `defer wg.Done()`.
   - Compute `n * n * n` (the cube of n).
   - Send `Result{ID: n, Value: n*n*n}` to the channel.
4. Launch a separate goroutine that calls `wg.Wait()` and then closes the channel.
5. Receive all results from the channel using a `for range` loop and store them in a slice at index `result.ID - 1`.
6. Print each result in order using `fmt.Printf("ID: %d, Value: %d\n", ...)`.

## Expected Output

```
ID: 1, Value: 1
ID: 2, Value: 8
ID: 3, Value: 27
ID: 4, Value: 64
ID: 5, Value: 125
```

## Hints

- Call `wg.Add(1)` before launching each goroutine.
- Use `defer wg.Done()` as the first line inside each goroutine.
- The "closer" goroutine pattern: `go func() { wg.Wait(); close(ch) }()`.
- Store results in a pre-allocated slice indexed by `result.ID - 1` for deterministic output.
