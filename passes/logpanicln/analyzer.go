package logpanicln

import (
	"go/ast"

	"github.com/bflad/gopaniccheck/passes/logpaniclncallexpr"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "logpanicln",
	Doc:  "reports log.Panicln() usage",
	Requires: []*analysis.Analyzer{
		logpaniclncallexpr.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	callExprs := pass.ResultOf[logpaniclncallexpr.Analyzer].([]*ast.CallExpr)

	for _, callExpr := range callExprs {
		pass.Reportf(callExpr.Pos(), "avoid log.Panicln() usage")
	}

	return nil, nil
}
