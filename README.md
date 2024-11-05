# natsort

`natsort` is a lightweight Go package that provides natural (alphanumeric) string comparison. It allows you to sort strings containing both letters and numbers in a human-friendly order, ensuring that filenames like `string1`, `string2`, `string10` are sorted as expected.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Compare Function](#compare-function)
- [Examples](#examples)
- [API](#api)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Natural Alphanumeric Sorting:** Sorts strings containing both letters and numbers in a way that humans expect.
- **Lightweight:** Minimal dependencies, easy to integrate into your projects.
- **Flexible:** Can be used for various applications like sorting filenames, version numbers, etc.
- **Performance:** Optimized for efficient comparison even with large datasets.

## Installation

To install the `natsort` package, use `go get`:

```bash
go get github.com/enkodr/natsort
```

## Usage

Import the `natsort` package into your Go project:

```go
import "github.com/enkodr/natsort"
```

### Compare Function

The `Compare` function allows you to perform natural comparison between two strings.

#### Function Signature:

```go
func Compare(a, b string) bool
```

**Parameters:**
- `a` (string): The first string to compare.
- `b` (string): The second string to compare.

**Returns:**
- `bool`: Returns true if string a is considered less than string b in natural order, otherwise false.

## Examples:

### Sorting a Slice of Strings

Here's how you can sort a slice of strings using the `Compare` function:

```go
package main

import (
    "fmt"
    "sort"

    "github.com/enkodr/natsort"
)

func main() {
    files := []string{"file10.txt", "file2.txt", "file1.txt", "file20.txt", "file3.txt"}

    sort.Slice(files, func(i, j int) bool {
        return natsort.Compare(files[i], files[j])
    })

    fmt.Println("Sorted files:", files)
}
```

**Output:**

```less
Sorted files: [file1.txt file2.txt file3.txt file10.txt file20.txt]
```

### Comparing Two Strings

You can directly compare two strings using the `Compare` function:

```go
package main

import (
    "fmt"

    "github.com/enkodr/natsort"
)

func main() {
    a := "version2.10"
    b := "version2.2"

    if natsort.Compare(a, b) {
        fmt.Printf("%s is less than %s\n", a, b)
    } else {
        fmt.Printf("%s is not less than %s\n", a, b)
    }
}
```

**Output:**

```less
version2.10 is not less than version2.2
```

## API

`Compare(a, b string) bool`

Compares two strings using natural alphanumeric order.

- **Parameters:**
  - `a` (string): The first string.
  - `b` (string): The second string.

- **Returns:**
  - `bool`: `true` if `a` is less than `b` in natural order, otherwise `false`.

**Example:**

```go
result := natsort.Compare("file1", "file2") // returns true
result = natsort.Compare("file10", "file2") // returns false
```

## Contributing

Contributions are welcome! If you'd like to contribute to natsort, please follow these steps:
1. **Fork the Repository:** Click the "Fork" button at the top right of the repository page.
2. **Clone the Repository:** Clone your fork to your local machine.
```bash
git clone https://github.com/enkodr/natsort.git
```
3. **Create a Branch:** Create a new branch for your feature or bugfix.
```bash
git checkout -b feature/your-feature-name
```
4. **Make Changes:** Implement your feature or fix the bug.
5. **Commit Changes:** Commit your changes with a descriptive commit message.
```bash
git commit -m "Add feature XYZ"
```
6. **Push to GitHub:** Push your branch to your fork.
```bash
git push origin feature/your-feature-name
```
7. **Create a Pull Request:** Go to the original repository and create a pull request from your fork.

Please ensure your code follows the existing style and includes appropriate tests.

## License
This project is licensed under the [MIT License](LICENSE).
