# SQL Builder

Welcome to SQL Builder on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.
If you get stuck on the exercise, check out `HINTS.md`, but try and solve it without using those first :)

## Introduction

In Go, strings are immutable.
Concatenating strings with `+` in a loop causes frequent allocations and copies data multiple times.
The [`strings` package](https://pkg.go.dev/strings) provides [`Builder`](https://pkg.go.dev/strings#Builder) to minimize memory overhead by efficiently assembling strings.

## Building Strings

The zero-value of a `strings.Builder` is ready to use immediately without any explicit initialization or allocation.
You can push new data into the builder using specific write methods.
Each write method returns an error value to satisfy standard Go interfaces, but it is guaranteed to always be `nil` and is typically ignored in everyday code.

The most frequently used methods are:

- `WriteString(s string) (int, error)`: Appends a string to the buffer.
- `WriteRune(r rune) (int, error)`: Appends the UTF-8 encoding of a rune.
- `WriteByte(c byte) error`: Appends a single byte.

```go
var sb strings.Builder

sb.WriteString("Hello")
var r rune = ' '
sb.WriteRune(r)
var b byte = '.'
sb.WriteByte(b)
```

## Retrieving and Resetting

Once you have finished assembling your text, you can extract the accumulated content as a standard Go string by calling the `String()` method.
Calling `String()` is a very cheap operation because it returns the underlying byte slice as a string directly without copying the accumulated data.

If you need to reuse the same builder instance to assemble a new string (for example, inside a loop), you can call the `Reset()` method.
This empties the internal buffer, allowing you to start fresh and build the next string while preserving the already allocated memory capacity under the hood.

```go
var sb strings.Builder

sb.WriteString("Task One")
fmt.Println(sb.String()) // Output: Task One

// Clear the buffer for resource reuse
sb.Reset()

sb.WriteString("Task Two")
fmt.Println(sb.String()) // Output: Task Two
```

## Instructions

Andrew, a developer at his company, recently built a lightweight SQL query generator utility for a database migration tool.
His first naive implementation worked perfectly for small sets of data.
However, during large production migrations, the application started consuming excessive memory and struggled under heavy load.
After profiling the application, Andrew realized that standard string concatenation inside loops was forcing Go to constantly reallocate memory.
Now, he needs your help to refactor the utility using Go's `strings.Builder` to ensure the application scales well and handles large volumes of text without crushing the memory heap.

## 1. Generate SELECT Queries

Implement the `BuildSelectQuery` function.
It takes a table name (`string`) and a slice of column names (`[]string`).
It should return a fully formed SQL `SELECT` statement ending with a semicolon.
If the columns slice is empty or `nil`, the query should fall back to using the `*` wildcard to select all columns.

```go
columns := []string{"id", "name", "email"}
BuildSelectQuery("users", columns)
// Function return: "SELECT id, name, email FROM users;"

BuildSelectQuery("orders", nil)
// Function return: "SELECT * FROM orders;"
```

## 2. Generate Batch INSERT Statements

Implement the `BuildBatchInserts` function.
It takes a table name (`string`) and a two-dimensional slice representing database rows (`[][]ColumnValue`).
Each element in a row links a specific column name to its string value.
The function must return a slice of individual `INSERT` SQL strings.
The structured slice format guarantees a predictable and stable query output order.
If the input slice of rows is empty, return an empty string slice.
If any individual row is empty, skip it.
String values within the `VALUES` clause must be enclosed in single quotes `'`.

```go
rows := [][]ColumnValue{
    {
        {Column: "title", Value: "Book"},
        {Column: "price", Value: "100"},
    },
    {
        {Column: "title", Value: "Magazine"},
        {Column: "price", Value: "15"},
    },
}

BuildBatchInserts("products", rows)
// Function returns string slice: [
//   "INSERT INTO products (title, price) VALUES ('Book', '100');",
//   "INSERT INTO products (title, price) VALUES ('Magazine', '15');"
// ]
```

## Source

### Created by

- @usernaimandrey