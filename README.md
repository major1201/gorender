# gorender
Golang text/html template render command line utility

[![Go Report Card](https://goreportcard.com/badge/github.com/major1201/gorender)](https://goreportcard.com/report/github.com/major1201/gorender)

## How to install

```sh
$ go get -u github.com/major1201/gorender
```

## How to use

### Flags

#### -h, --help

Show help

#### -v, --version

Print the version

#### -a, --arguments [value]

Set extra arguments to be rendered to the template, overrides all json/yaml/toml file arguments

Format: *key*=*value*

This flag can be used multiple times.

#### -o, --output [path]

Write output to a file instead of to stdout

#### -i, --in-place

Write output in place of the template file

#### --json [path]

Use a json argument file

#### --yaml, --yml [path]

Use a yaml argument file

#### --toml [path]

Use a toml argument file

#### --html

Enable html template engine which automatically secures HTML output against certain attacks

### Example

```sh
$ gorender --json example.json -a name=override -o /tmp/output.txt example.txt.tmpl
```
