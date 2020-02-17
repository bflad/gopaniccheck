package passes

import (
	"github.com/bflad/gopaniccheck/passes/logpanic"
	"github.com/bflad/gopaniccheck/passes/logpanicf"
	"github.com/bflad/gopaniccheck/passes/logpanicln"
	"github.com/bflad/gopaniccheck/passes/panic"
	"golang.org/x/tools/go/analysis"
)

// AllReportingAnalyzers contains all Analyzers that report issues
// This can be consumed via multichecker.Main(passes.AllReportingAnalyzers...) or by
// combining these Analyzers with additional custom Analyzers
var AllReportingAnalyzers = []*analysis.Analyzer{
	logpanic.Analyzer,
	logpanicf.Analyzer,
	logpanicln.Analyzer,
	panic.Analyzer,
}
