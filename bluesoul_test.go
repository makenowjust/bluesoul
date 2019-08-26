package bluesoul_test

import (
	"testing"

	"github.com/MakeNowJust/bluesoul"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, bluesoul.Analyzer, "a")
}
