# Exercise 04: Embedding

## Objective

Learn how to use struct embedding to compose types and promote fields and methods.

## Instructions

1. Define an `Animal` struct with a `Name` field (string).
2. Add a `Speak()` method to `Animal` that returns `"..."`.
3. Define a `Dog` struct that **embeds** `Animal` and adds a `Breed` field (string).
4. Override the `Speak()` method on `Dog` to return `"Woof!"`.
5. Create a `Dog` with `Name: "Buddy"` and `Breed: "Labrador"`.
6. Print the promoted `Name` field.
7. Print the `Breed` field.
8. Print the result of calling `Speak()` on the dog, formatted as `<Name> says: <Speak()>`.

## Expected Output

```
Name: Buddy
Breed: Labrador
Buddy says: Woof!
```
