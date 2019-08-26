package bluesoul

import (
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var (
	// regular expression to match "the the"/"a a"/"an an"
	reTheThe = regexp.MustCompile("\\b(?:the|an?)\\s+(?:the|an?)\\b")
)

var Analyzer = &analysis.Analyzer{
	Name: "bluesoul",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "bluesoul is duplicate article finder"

func run(pass *analysis.Pass) (interface{}, error) {
	file := pass.Files[0]
	for _, comment := range file.Comments {
		text := comment.Text()
		if reTheThe.MatchString(text) {
			pass.Reportf(comment.Pos(), "duplicate article is found")
		}
	}

	return nil, nil
}
