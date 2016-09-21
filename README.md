# go-inquote

A library that translates runes into their representation inside a quoted string, for the Go programming language.

For example

* `rune` value `10` (which is the "line feed" character) becomes `\n`
* `rune` value `31` (which is the "unit separator" character) becomes `\x1f`
* etc

And of course, something like `rune` value `65` (which is the "uppercase A" character) stays as `A`.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-inquote

[![GoDoc](https://godoc.org/github.com/reiver/go-inquote?status.svg)](https://godoc.org/github.com/reiver/go-inquote)
