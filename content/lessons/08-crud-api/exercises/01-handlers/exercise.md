# Exercise 1: HTTP Handlers

## Objective

Write a simple HTTP handler that responds with the text `Hello, API!`.

## Instructions

Open the file `starter.go`. The program starts a test HTTP server using `httptest.NewServer`, makes a GET request to it, and prints the response body.

Your task is to **complete the `helloHandler` function** so it writes `Hello, API!` as the response body.

## What to Do

Inside `helloHandler`, use `fmt.Fprint` to write `"Hello, API!"` to the `http.ResponseWriter`.

## Expected Output

```
Hello, API!
```

## Hints

- Use `fmt.Fprint(w, "Hello, API!")` to write text to the response.
- The `w` parameter is the `http.ResponseWriter` — writing to it sends data back to the client.
- You do not need to call `w.WriteHeader` — it defaults to `200 OK`.
