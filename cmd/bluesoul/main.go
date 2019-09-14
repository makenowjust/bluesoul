package main

import (
	"flag"

	"github.com/MakeNowJust/bluesoul"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// `-unsafeptr` flag is given when it is used via `go vet -vettool`.
// But the flag is not defined in singlechecker for now.
//
// See https://github.com/golang/go/issues/34053
var _ = flag.Bool("unsafeptr", false, "(for Go 1.13 compatibilit)")

func main() { singlechecker.Main(bluesoul.Analyzer) }
