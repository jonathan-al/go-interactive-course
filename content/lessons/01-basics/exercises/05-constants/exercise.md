# Exercise 5: Constants and iota

## Objective

Learn how to use constants and the `iota` enumerator in Go.

## Instructions

Open `starter.go`. Two `const` blocks using `iota` are already defined:

1. **Days of the week** — `Sunday` through `Saturday` (values 0–6)
2. **HTTP status codes** — `StatusOK` (200), `StatusCreated` (201), `StatusAccepted` (202)

The current print statements output raw values on single lines. Your task is to **modify the print statements** so each value is printed with a label on its own line:

```
Sunday: 0
Monday: 1
Tuesday: 2
StatusOK: 200
StatusCreated: 201
StatusAccepted: 202
```

## Expected Output

```
Sunday: 0
Monday: 1
Tuesday: 2
StatusOK: 200
StatusCreated: 201
StatusAccepted: 202
```

## Hints

- Use `fmt.Println` with multiple arguments: `fmt.Println("Sunday:", Sunday)` prints `Sunday: 0`.
- `iota` starts at 0 and increments by 1 for each constant in the block.
- `iota + 200` starts the sequence at 200.
- Print each value on its own line (6 `fmt.Println` calls total).
