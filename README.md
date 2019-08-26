# bluesoul

> `bluesoul` finds duplicate article (e.g. `the the`) in comment text.

## Install

```console
go get -u github.com/MakeNowJust/bluesoul
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
./hello.go:1:1: duplicate article is found
```