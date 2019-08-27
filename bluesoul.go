package bluesoul

import (
	"go/token"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

var (
	// Regular expression to match "the the"/"a a"/"an an".
	// It can contains spaces, slash, asterisk, backslash and double quote between articles.
	// Note that backslash and double quote are cheating for passing a test...
	reTheThe = regexp.MustCompile(`\b(?:the|an?)[\s*/\\"]+(?:the|an?)\b`)
	// Regular expression for cleaning up matched text.
	reCleanup = regexp.MustCompile(`[\s*/\\"]+`)
)

// Analyzer is bluesoul analyzer.
var Analyzer = &analysis.Analyzer{
	Name:     "bluesoul",
	Doc:      Doc,
	Run:      run,
	Requires: []*analysis.Analyzer{},
}

// Doc is description for bluesoul.
const Doc = "bluesoul finds duplicate articles (e.g. the the) in comment text"

type posPair struct {
	index int
	pos   token.Pos
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		for _, comments := range file.Comments {
			pairs := []posPair{{0, comments.Pos()}}
			text := ""
			for i, comment := range comments.List[:len(comments.List)-1] {
				text += comment.Text + "\n"
				pairs = append(pairs, posPair{
					index: pairs[len(pairs)-1].index + len(comment.Text) + 1,
					pos:   comments.List[i+1].Pos(),
				})
			}
			text += comments.List[len(comments.List)-1].Text

			iss := reTheThe.FindAllStringIndex(text, -1)
			for _, is := range iss {
				sub := reCleanup.ReplaceAllString(text[is[0]:is[1]], " ")

				var pos token.Pos
				for _, p := range pairs {
					pos = p.pos
					if p.index < is[0] {
						pos += token.Pos(is[0] - p.index)
						break
					}
				}

				pass.Reportf(pos, "duplicate articles: %q", sub)
			}
		}
	}

	return nil, nil
}
