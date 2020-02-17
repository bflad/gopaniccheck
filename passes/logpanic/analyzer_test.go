package logpanic

import (
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}

func TestValidateAnalyzer(t *testing.T) {
	err := analysis.Validate([]*analysis.Analyzer{Analyzer})

	if err != nil {
		t.Fatal(err)
	}
}
