# Exercise 3: Select

## Objective

Use the `select` statement to receive values from two channels in a deterministic order.

## Instructions

Open `starter.go`. You'll find a goroutine that sends values to two channels **sequentially** (first to `ch1`, then to `ch2`). Your task is to use `select` in a loop to receive from both channels and print the messages.

Because the goroutine sends to `ch1` first (blocking on the unbuffered channel until received), and only then sends to `ch2`, the `select` will always receive from `ch1` first and `ch2` second.

Your tasks:

1. Use a `for` loop that runs 2 iterations.
2. Inside the loop, use a `select` statement with two cases:
   - Receive from `ch1` and print the message with `fmt.Println`.
   - Receive from `ch2` and print the message with `fmt.Println`.

## Expected Output

```
first
second
```

## Hints

- The goroutine sends `"first"` to `ch1` and blocks until the main goroutine receives it.
- Only after `ch1` is received does the goroutine send `"second"` to `ch2`.
- This sequential sending guarantees deterministic output order.
