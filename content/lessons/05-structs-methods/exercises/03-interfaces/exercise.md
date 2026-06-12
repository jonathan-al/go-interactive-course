# Exercise 03: Interfaces

## Objective

Learn how to define an interface and implement it with multiple types to achieve polymorphism.

## Instructions

1. Define a `Shape` interface with a single method `Area() float64`.
2. Define a `Circle` struct with a `Radius` field of type `float64`.
3. Implement the `Area()` method for `Circle` using `math.Pi * Radius * Radius`.
4. Define a `Rectangle` struct with `Width` and `Height` fields of type `float64`.
5. Implement the `Area()` method for `Rectangle` returning `Width * Height`.
6. Write a `printArea` function that takes a `Shape` and prints its area in the format `<TypeName> area: <value>` using `%.2f` formatting.
7. In `main`, create a `Circle` with `Radius: 5` and a `Rectangle` with `Width: 4`, `Height: 6`.
8. Call `printArea` for each shape.

## Hint

Use a type switch inside `printArea` to determine the concrete type and print the appropriate label.

## Expected Output

```
Circle area: 78.54
Rectangle area: 24.00
```
