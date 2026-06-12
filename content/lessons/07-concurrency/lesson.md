# Lesson 07: Concurrency

Concurrency is one of Go's most powerful features. Go provides built-in support for running functions concurrently using **goroutines** and communicating between them using **channels**.

## Table of Contents

1. [Goroutines](#1-goroutines)
2. [Channels](#2-channels)
3. [Select](#3-select)
4. [WaitGroup](#4-waitgroup)

---

## 1. Goroutines

A **goroutine** is a lightweight thread managed by the Go runtime. You start one by prefixing a function call with the `go` keyword:

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    go sayHello()
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Done")
}
```

### Key Concepts

- The `go` keyword launches a function in a separate goroutine.
- The calling goroutine does **not** wait for the launched goroutine to finish.
- If `main()` exits, all goroutines are terminated immediately.
- Goroutines are extremely lightweight — you can run thousands of them.

### Anonymous Goroutines

You can launch goroutines from anonymous functions:

```go
go func(name string) {
    fmt.Printf("Hello, %s!\n", name)
}("Alice")
```

### The Problem with `time.Sleep`

Using `time.Sleep` to wait for goroutines is fragile and unreliable. Go provides better synchronization tools: **channels** and **WaitGroup**.

---

## 2. Channels

**Channels** are typed conduits through which goroutines communicate. Think of them as pipes: you send values in one end and receive them from the other.

### Creating Channels

```go
ch := make(chan int)       // unbuffered channel
ch := make(chan int, 5)    // buffered channel with capacity 5
```

### Sending and Receiving

```go
ch <- 42      // send 42 into the channel
value := <-ch  // receive a value from the channel
```

### Unbuffered Channels

An unbuffered channel has no internal buffer. A send blocks until another goroutine is ready to receive, and vice versa. This provides a **synchronization point**:

```go
ch := make(chan string)

go func() {
    ch <- "hello"
}()

msg := <-ch
fmt.Println(msg) // "hello"
```

### Buffered Channels

A buffered channel has a fixed capacity. Sends only block when the buffer is full, and receives only block when the buffer is empty:

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
// ch <- 4  // This would block — buffer is full

fmt.Println(<-ch) // 1
fmt.Println(<-ch) // 2
```

### Closing Channels and `range`

A sender can **close** a channel to signal that no more values will be sent. The receiver can use `range` to read values until the channel is closed:

```go
ch := make(chan int, 3)
ch <- 10
ch <- 20
ch <- 30
close(ch)

for v := range ch {
    fmt.Println(v)
}
```

**Rules:**
- Only the sender should close a channel.
- Sending on a closed channel causes a panic.
- Receiving from a closed channel returns the zero value.

### Channel Direction

You can restrict a channel parameter to send-only or receive-only:

```go
func producer(out chan<- int) { ... }  // can only send
func consumer(in <-chan int) { ... }   // can only receive
```

---

## 3. Select

The `select` statement lets a goroutine wait on multiple channel operations simultaneously. It's like a `switch` but for channels:

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case ch3 <- "hello":
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No channel ready")
}
```

### How Select Works

- `select` blocks until one of its cases can proceed.
- If multiple cases are ready, one is chosen **at random**.
- The `default` case runs immediately if no other case is ready (non-blocking select).

### Example: Multiplexing Two Channels

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    ch1 <- "ping"
}()

go func() {
    ch2 <- "pong"
}()

for i := 0; i < 2; i++ {
    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    }
}
```

**Note:** Because `select` picks randomly among ready cases, the output order is non-deterministic. To get deterministic output, you need to control the order in which channels become ready (e.g., using sequential sends on unbuffered channels).

### Timeout with Select

A common pattern is using `select` with `time.After` for timeouts:

```go
select {
case result := <-ch:
    fmt.Println("Got:", result)
case <-time.After(2 * time.Second):
    fmt.Println("Timed out!")
}
```

---

## 4. WaitGroup

`sync.WaitGroup` provides a way to wait for a collection of goroutines to finish. It's more reliable than `time.Sleep`.

### Basic Usage

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }

    wg.Wait()
    fmt.Println("All workers finished")
}
```

### How WaitGroup Works

| Method       | Purpose                                         |
|--------------|------------------------------------------------|
| `wg.Add(n)`  | Increments the counter by `n`                  |
| `wg.Done()`  | Decrements the counter by 1 (call when goroutine finishes) |
| `wg.Wait()`  | Blocks until the counter reaches 0             |

### Collecting Results with WaitGroup + Channel

A common pattern is to use a WaitGroup to know when all goroutines are done, and a channel to collect their results:

```go
type Result struct {
    ID    int
    Value int
}

func main() {
    var wg sync.WaitGroup
    results := make(chan Result, 5)

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            results <- Result{ID: id, Value: id * id}
        }(i)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    collected := make([]Result, 5)
    for r := range results {
        collected[r.ID-1] = r
    }

    for _, r := range collected {
        fmt.Printf("ID: %d, Value: %d\n", r.ID, r.Value)
    }
}
```

This pattern ensures **deterministic output** regardless of goroutine execution order, because results are placed into a slice by their ID and then printed sequentially.

---

## Summary

In this lesson you learned:

- **Goroutines** — lightweight concurrent execution with the `go` keyword.
- **Channels** — typed pipes for communication between goroutines (buffered and unbuffered).
- **Select** — multiplexing across multiple channel operations.
- **WaitGroup** — waiting for a collection of goroutines to finish.
- How to combine WaitGroup and channels to collect results deterministically.

Now complete the exercises to practice what you've learned!
