package passes

import (
	"testing"

	"golang.org/x/tools/go/analysis"
)

func TestValidateAllReportingAnalyzers(t *testing.T) {
	err := analysis.Validate(AllReportingAnalyzers)

	if err != nil {
		t.Fatal(err)
	}
}
