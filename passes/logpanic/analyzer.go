package logpanic

import (
	"go/ast"

	"github.com/bflad/gopaniccheck/passes/logpaniccallexpr"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "logpanic",
	Doc:  "reports log.Panic() usage",
	Requires: []*analysis.Analyzer{
		logpaniccallexpr.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	callExprs := pass.ResultOf[logpaniccallexpr.Analyzer].([]*ast.CallExpr)

	for _, callExpr := range callExprs {
		pass.Reportf(callExpr.Pos(), "avoid log.Panic() usage")
	}

	return nil, nil
}
