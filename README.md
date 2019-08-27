# bluesoul

> `bluesoul` finds duplicate articles (e.g. `the the`) in comment text.

## Install

```console
$ go get -u github.com/MakeNowJust/bluesoul/cmd/bluesoul
```

## Usage

```console
$ go vet -vettool=$(which bluesoul) pkgname
```

Or:

```console
$ bluesoul pkgname
```

## Example

```console
$ cat hello.go
// It is the the "Hello World" programming for Go.
package main

func main() {
    println("Hello, World")
}
$ go vet -vettool=$(which bluesoul) .
# github.com/MakeNowJust/bluesoul/hello
./hello.go:1:10: duplicate articles: "the the"
```