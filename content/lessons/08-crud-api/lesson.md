# Lesson 08: Building a CRUD API

Welcome to the final module! In this lesson, you'll learn how to build a RESTful CRUD API in Go using only the standard library.

## Table of Contents

1. [HTTP Handlers](#1-http-handlers)
2. [Routing](#2-routing)
3. [JSON Encoding and Decoding](#3-json-encoding-and-decoding)
4. [Building a CRUD API](#4-building-a-crud-api)

---

## 1. HTTP Handlers

Go's `net/http` package provides everything you need to build HTTP servers. The core interface is `http.Handler`:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

However, the most common way to write handlers is with functions that match the signature `func(http.ResponseWriter, *http.Request)`.

### ResponseWriter and Request

- **`http.ResponseWriter`** — Used to construct the HTTP response (write headers and body).
- **`*http.Request`** — Contains all information about the incoming HTTP request (method, URL, headers, body).

### Writing a Basic Handler

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Hello, API!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Key Functions

| Function | Description |
|---|---|
| `http.HandleFunc(pattern, handler)` | Registers a handler function for a URL pattern |
| `http.ListenAndServe(addr, handler)` | Starts the HTTP server on the given address |
| `fmt.Fprint(w, ...)` | Writes formatted text to the ResponseWriter |
| `w.WriteHeader(code)` | Sets the HTTP status code |
| `w.Header().Set(key, value)` | Sets a response header |

### HTTP Methods

The `*http.Request` has a `Method` field you can use to handle different HTTP verbs:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprint(w, "GET request")
    case http.MethodPost:
        fmt.Fprint(w, "POST request")
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}
```

---

## 2. Routing

Routing is the process of matching incoming requests to the correct handler based on the URL path.

### Default ServeMux

Go's `http.DefaultServeMux` (used when you pass `nil` to `ListenAndServe`) provides basic routing:

```go
http.HandleFunc("/users", usersHandler)
http.HandleFunc("/products", productsHandler)
```

### Creating a Custom ServeMux

For better organization, create your own mux:

```go
mux := http.NewServeMux()
mux.HandleFunc("/users", usersHandler)
mux.HandleFunc("/products", productsHandler)
http.ListenAndServe(":8080", mux)
```

### Pattern Matching Rules

| Pattern | Matches |
|---|---|
| `/users` | Only the exact path `/users` |
| `/users/` | `/users/` and any path starting with `/users/` (e.g., `/users/1`, `/users/abc`) |

The longest matching pattern wins. This is called **longest-prefix matching**.

### Extracting Path Parameters

To extract an ID from a URL like `/users/1`, strip the prefix and parse:

```go
func userHandler(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/users/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "User ID: %d", id)
}
```

---

## 3. JSON Encoding and Decoding

Go's `encoding/json` package handles JSON serialization and deserialization.

### Struct Tags

Use struct tags to control how fields are mapped to JSON keys:

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### Encoding (Struct to JSON)

```go
user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
data, err := json.Marshal(user)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
// Output: {"id":1,"name":"Alice","email":"alice@example.com"}
```

To write JSON directly to an `http.ResponseWriter`:

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(user)
```

### Decoding (JSON to Struct)

```go
jsonData := `{"id":1,"name":"Alice","email":"alice@example.com"}`
var user User
err := json.Unmarshal([]byte(jsonData), &user)
if err != nil {
    log.Fatal(err)
}
fmt.Println(user.Name) // "Alice"
```

To read JSON from an HTTP request body:

```go
var user User
err := json.NewDecoder(r.Body).Decode(&user)
if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
}
```

### Key Functions

| Function | Description |
|---|---|
| `json.Marshal(v)` | Converts a Go value to JSON bytes |
| `json.Unmarshal(data, &v)` | Parses JSON bytes into a Go value |
| `json.NewEncoder(w).Encode(v)` | Writes JSON directly to a writer |
| `json.NewDecoder(r).Decode(&v)` | Reads JSON directly from a reader |

---

## 4. Building a CRUD API

CRUD stands for **Create, Read, Update, Delete**. Let's build an in-memory API for managing users.

### In-Memory Store

```go
var (
    users  = make(map[int]User)
    nextID = 1
    mu     sync.Mutex
)
```

A `sync.Mutex` protects concurrent access to the shared map.

### Create (POST /users)

```go
func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    mu.Lock()
    user.ID = nextID
    nextID++
    users[user.ID] = user
    mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
```

### Read (GET /users/{id})

```go
func getUser(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/users/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    mu.Lock()
    user, ok := users[id]
    mu.Unlock()

    if !ok {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

### Putting It All Together

```go
func usersRouter(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.URL.Path == "/users" && r.Method == http.MethodPost:
        createUser(w, r)
    case strings.HasPrefix(r.URL.Path, "/users/") && r.Method == http.MethodGet:
        getUser(w, r)
    default:
        w.WriteHeader(http.StatusNotFound)
    }
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/users", usersRouter)
    mux.HandleFunc("/users/", usersRouter)
    http.ListenAndServe(":8080", mux)
}
```

### Testing with httptest

Go's `net/http/httptest` package lets you create test servers that run and shut down within a single program:

```go
ts := httptest.NewServer(mux)
defer ts.Close()

resp, err := http.Get(ts.URL + "/users/1")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
fmt.Println(string(body))
```

This is the approach used in the exercises: the program starts a test server, makes requests to itself, prints the results, and exits.

---

## Summary

In this lesson you learned:

- How to write HTTP handlers using `http.ResponseWriter` and `*http.Request`
- How to set up routing with `http.ServeMux`
- How to encode and decode JSON with `encoding/json`
- How to build a complete CRUD API with in-memory storage
- How to test HTTP servers using `net/http/httptest`

Now complete the exercises to build your own CRUD API step by step!
