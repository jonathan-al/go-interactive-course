# Exercise 2: Variables

## Objective

Declare variables of different types and print their values with labels.

## Instructions

Open `starter.go`. Four variables are already declared and printed. Your task is to **modify the print statements** so each variable is printed with a label in the format `Label: value`.

Change the existing `fmt.Println` calls to produce:

```
Name: Alice
Age: 30
Height: 1.68
Student: true
```

## Expected Output

```
Name: Alice
Age: 30
Height: 1.68
Student: true
```

## Hints

- You can pass multiple arguments to `fmt.Println` — it separates them with spaces.
- Example: `fmt.Println("Name:", name)` prints `Name: Alice`.
- Make sure the labels match exactly: `Name:`, `Age:`, `Height:`, `Student:`.
