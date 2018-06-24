# Numero-Go

[![Build Status](https://travis-ci.org/faridgh1991/Numero-Go.svg?branch=master)](https://travis-ci.org/faridgh1991/Numero-Go)


A micro library for converting non-english UTF8 digits. (like ۱=1, ۲=2)

## Supported languages

Almost all numbers defined in Unicode is supported in Numero.

For more info on supported characters you can visit [here](http://www.fileformat.info/info/unicode/category/Nd/list.htm)

## Installation

Use `go get` on this repository:

```sh
$ go get -u github.com/faridgh1991/Numero-Go
```

## Using Numero

On strings for strings:

```go
result = numero.Normalize("1۲۳۰4a۳tس")
// result = "12304a3tس"
```

Smart numeric convert:

(Convert numbers to Integer or Float based on input string)

```go
result, err = numero.NormalizeAsNumber("1۲۳۰4۳")
// result = 123043
// err = nil

result, err = numero.NormalizeAsNumber("1۲۳۰4۳.۴5")
// result = 123043.45
// err = nil

result, err = numero.NormalizeAsNumber("1۲a۳۰4۳.۴5")
// result = 0
// err = parsing "1۲a۳۰4۳.۴5": invalid syntax
```

Strip all non numeric chars from a string:

```go
result = numero.RemoveNonDigits("12 345abs")
// result = "12345"

// Or even make exceptions for some chars like 'a' and ' ' (space)
result = numero.RemoveNonDigits("12 345bas", "a ")
// result = "12 345a"
```

Checking if a string is all numbers
```go
result = numero.DigitOnly("1234567890")
// result = true

result = numero.DigitOnly("1234567890.a")
// result = false
```
