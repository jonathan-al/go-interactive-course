# Exercise 3: JSON Encoding and Decoding

## Objective

Learn how to encode a Go struct to JSON and decode a JSON string back into a struct.

## Instructions

Open the file `starter.go`. The program defines a `User` struct and has two tasks:

1. **Encode** a `User` struct to JSON using `json.Marshal` and print the resulting JSON string.
2. **Decode** a JSON string into a `User` struct using `json.Unmarshal` and print the user's name and email.

## What to Do

- In the encoding section, call `json.Marshal(user)` and convert the result to a string with `string(data)`.
- In the decoding section, call `json.Unmarshal([]byte(jsonStr), &decoded)` to parse the JSON string into the `decoded` variable.

## Expected Output

```
{"id":1,"name":"Alice","email":"alice@example.com"}
Name: Bob, Email: bob@example.com
```

## Hints

- `json.Marshal` returns `([]byte, error)`. Use `string(data)` to convert bytes to a printable string.
- `json.Unmarshal` takes a byte slice and a pointer to the target struct: `json.Unmarshal([]byte(jsonStr), &decoded)`.
- The struct fields must be exported (capitalized) for `encoding/json` to access them.
- Struct tags like `` `json:"id"` `` control the JSON key names.
