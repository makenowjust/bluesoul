package main

import (
	"github.com/MakeNowJust/bluesoul"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(bluesoul.Analyzer) }