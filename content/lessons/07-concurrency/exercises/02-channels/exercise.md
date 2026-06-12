# Exercise 2: Channels

## Objective

Create a channel, send integer values into it from a goroutine, then receive and print them in the main goroutine.

## Instructions

Open `starter.go`. Your tasks:

1. Create an **unbuffered** channel of type `int`.
2. Launch a goroutine that sends the values 10, 20, 30, 40, 50 into the channel, then closes it.
3. In the main goroutine, use a `for range` loop to receive all values from the channel and print each one with `fmt.Println`.

## Expected Output

```
10
20
30
40
50
```

## Hints

- Create an unbuffered channel with `make(chan int)`.
- Send values with `ch <- value`.
- Close the channel after sending all values with `close(ch)`.
- `for v := range ch` reads until the channel is closed.
