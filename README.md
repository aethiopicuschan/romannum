# aethiopicuschan/romannum

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/romannum.svg)](https://pkg.go.dev/github.com/aethiopicuschan/romannum)

Convert int to Roman numerals and Roman numerals to int.

## Getting Started

```sh
go get github.com/aethiopicuschan/romannum
```

## Usage

```go
package main

import "github.com/aethiopicuschan/romannum"

func main() {
  roman, err := romannum.ToRoman(1776)
  ...
}
```

## Running Tests

```sh
go test
```
