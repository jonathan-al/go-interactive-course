# Exercise 2: Routing

## Objective

Create handlers for different routes and register them on an HTTP mux.

## Instructions

Open the file `starter.go`. The program creates an `http.ServeMux` and makes test requests to two different routes: `/users` and `/products`.

Your task is to:

1. **Complete the `usersHandler` function** so it writes `Users List` to the response.
2. **Complete the `productsHandler` function** so it writes `Products List` to the response.
3. **Register both handlers** on the mux in `main()` using `mux.HandleFunc`.

## Expected Output

```
Users List
Products List
```

## Hints

- Register routes with `mux.HandleFunc("/users", usersHandler)` and `mux.HandleFunc("/products", productsHandler)`.
- Use `fmt.Fprint(w, "...")` inside each handler to write the response text.
- The order of registration does not matter — the mux matches by URL pattern.
