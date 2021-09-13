# titlecase
[![Test status](https://github.com/thatoddmailbox/titlecase/workflows/Tests/badge.svg)](https://github.com/thatoddmailbox/titlecase/actions)

A Go module to convert text into title case. It attempts to follow the [Wikipedia Manual of Style](https://en.wikipedia.org/wiki/Wikipedia:Manual_of_Style/Titles#Capital_letters) where possible.

## Usage
The module exports one function, `Title`, which can be used like so:
```go
titlecase.Title("INTRODUCTION TO PROGRAMMING") // => "Introduction to Programming"
titlecase.Title("some cool book") // => "Some Cool Book"
titlecase.Title("an interesting story") // => "An Interesting Story"
```

For handling of weirder cases, see the [unit test file](./case_test.go).