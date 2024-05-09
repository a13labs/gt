# gt (JSON go-templating renderer)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Description

A command-line tool to render Go templates from JSON input data.

The tool is invoked with the path to the template file. If a second argument is provided, it is treated as the path to a JSON file to use as input for the template. If no second argument is provided, the tool reads JSON input from stdin.

## Features

- Support JSON input
- GO template language
- [Builtin functions](#builtin-functions)

## Builtin functions

#### env 
The `env` function returns the value of the specified environment variable.

```go
{{ env "ENV_VAR_NAME" }}
```

Expected output:
```shell
value
```

#### now 

The `now` function returns the current date and time.
```go
{{ now }}
```

Expected output:
```shell
2021-09-01 12:00:00
```

#### date 

The `date` function formats a date and time value using a specified layout.
```go
{{ now | date "2006-01-02"}}
```

Expected output:
```shell
2021-09-01
```

#### add 

The `add` function adds two numbers.
```go
{{ 1 | add 2 }}
```

Expected output:
```shell
3
```

#### sub 

The `sub` function subtracts two numbers.
```go
{{ 2 | sub 1 }}
```

Expected output:
```shell
1
```

#### mul 

The `mul` function multiplies two numbers.
```go
{{ 1 | mul 2 }}
```

Expected output:
```shell
2
```

#### div 

The `div` function divides two numbers.
```go
{{ 1 | div 2 }}
```

Expected output:
```shell
0.5
```

#### upper 

The `upper` function converts a string to uppercase.
```go
{{ "hello" | upper }}
```

Expected output:
```shell
HELLO
```

#### lower 

The `lower` function converts a string to lowercase.
```go
{{ "HELLO" | lower }}
```

Expected output:
```shell
hello
```

#### trim 

The `trim` function removes leading and trailing whitespace from a string.
```go
{{ "  hello  " | trim }}
```

Expected output:
```shell
hello
```

#### trimleft 

The `trimleft` function removes leading whitespace from a string.
```go
{{ "  hello  " | trimleft }}
```

Expected output:
```shell
hello  
```

#### trimright 

The `trimright` function removes trailing whitespace from a string.
```go
{{ "  hello  " | trimright }}
```

#### replace 

The `replace` function replaces all occurrences of a substring in a string with another substring.
```go
{{ "hello world" | replace "world" "go" }}
```

Expected output:
```shell
hello go
```

#### contains 

The `contains` function returns true if a string contains a specified substring.
```go
{{ "hello world" | contains "world" }}
```

Expected output:
```shell
true
```

#### hasprefix 

The `hasprefix` function returns true if a string has a specified prefix.
```go
{{ "hello world" | hasprefix "hello" }}
```

Expected output:
```shell
true
```

#### hassuffix 

The `hassuffix` function returns true if a string has a specified suffix.
```go
{{ "hello world" | hassuffix "world" }}
```

Expected output:
```shell
true
```

#### indexof 

The `indexof` function returns the index of the first occurrence of a substring in a string.
```go
{{ "hello world" | indexof "world" }}
```

Expected output:
```shell
6
```

#### lastindexof 

The `lastindexof` function returns the index of the last occurrence of a substring in a string.
```go
{{ "hello world" | lastindexof "o" }}
```

Expected output:
```shell
7
```

#### reverse 

The `reverse` function reverses a string.
```go
{{ "hello" | reverse }}
```

Expected output:
```shell
olleh
```

#### substr 

The `substr` function returns a substring of a string.
```go
{{ "hello world" | substr 6 }}
```

Expected output:
```shell
world
```

#### escapeString 

The `escapeString` function escapes special characters in a string.
```go
{{ "hello <world>" | escapeString }}
```

Expected output:
```shell
hello &lt;world&gt;
```

#### len 

The `len` function returns the length of a string or array.
```go
{{ "hello" | len }}
```

Expected output:
```shell
5
```

#### regexFind 

The `regexFind` function returns the first match of a regular expression in a string.
```go
{{ "hello world" | regexFind ".{4}$" }}
```

Expected output:
```shell
world
```

#### empty 

The `empty` function returns true if a string or array is empty.
```go
{{ "" | empty }}
```

Expected output:
```shell
true
```

## Installation

Download from release page.

## Usage

mytemplate.tpl:
```go
{{ range . }}
{{ .name }} is {{ .age }} years old
{{ end }}
```

mydata.json:
```json
[
  {
    "name": "Alice",
    "age": 30
  },
  {
    "name": "Bob",
    "age": 25
  }
]
```

Render from file
```shell
gt mytemplate.tpl mydata.json
```

Render from stdin
```shell
cat mydata.json | gt mytemplate.tpl
```

Expected output:
```shell
Alice is 30 years old
Bob is 25 years old
```

## Contributing

Contributions are welcome! Please follow the guidelines in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).

