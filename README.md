# Numero-Go

[![Build Status](https://travis-ci.org/faridgh1991/Numero-Go.svg?branch=master)](https://travis-ci.org/faridgh1991/Numero-Go)
[![Go Report Card](https://goreportcard.com/badge/github.com/faridgh1991/Numero-Go)](https://goreportcard.com/report/github.com/faridgh1991/Numero-Go)
[![GoDoc](https://godoc.org/github.com/faridgh1991/Numero-Go?status.svg)](https://godoc.org/github.com/faridgh1991/Numero-Go)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/faridgh1991/Numero-Go/blob/master/LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ffaridgh1991%2FNumero-Go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Ffaridgh1991%2FNumero-Go?ref=badge_shield)

A micro library for converting non-english UTF8 digits. (like ۱=1, ۲=2)

## Supported languages

Almost all numbers defined in Unicode are support in Numero.

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
## GoDoc
[See godoc page](https://godoc.org/github.com/faridgh1991/Numero-Go).


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ffaridgh1991%2FNumero-Go.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Ffaridgh1991%2FNumero-Go?ref=badge_large)