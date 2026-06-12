# Exercise 4: CRUD API

## Objective

Build a simple in-memory CRUD API that supports creating a user (POST) and retrieving a user by ID (GET).

## Instructions

Open the file `starter.go`. The program starts a test HTTP server, sends a POST request to create a user, then sends a GET request to retrieve that user.

Your task is to complete two handler functions:

1. **`createUser`** — Decode the JSON request body into a `User`, assign it an ID from `nextID`, store it in the `users` map, increment `nextID`, then respond with status `201 Created` and the user as JSON.
2. **`getUser`** — Extract the ID from the URL path (after `/users/`), look up the user in the `users` map, and respond with the user as JSON. If not found, respond with status `404 Not Found`.

## What to Do

- In `createUser`: use `json.NewDecoder(r.Body).Decode(&user)` to read the request body, then `json.NewEncoder(w).Encode(user)` to write the response.
- In `getUser`: use `strings.TrimPrefix(r.URL.Path, "/users/")` to extract the ID string, then `strconv.Atoi` to convert it to an integer.
- Set the `Content-Type` header to `application/json` before writing the response body.

## Expected Output

```
POST /users -> Status: 201
Created: {"id":1,"name":"Alice","email":"alice@example.com"}
GET /users/1 -> Status: 200
User: {"id":1,"name":"Alice","email":"alice@example.com"}
```

## Hints

- Use `w.WriteHeader(http.StatusCreated)` to set a 201 status code (must be called before writing the body).
- Use `w.Header().Set("Content-Type", "application/json")` to set the content type header (must be called before `WriteHeader`).
- The `users` map and `nextID` variable are declared at package level so both handlers can access them.
- Remember to lock the mutex (`mu.Lock()` / `mu.Unlock()`) when accessing shared state.
