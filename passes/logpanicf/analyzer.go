package logpanicf

import (
	"go/ast"

	"github.com/bflad/gopaniccheck/passes/logpanicfcallexpr"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "logpanicf",
	Doc:  "reports log.Panicf() usage",
	Requires: []*analysis.Analyzer{
		logpanicfcallexpr.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	callExprs := pass.ResultOf[logpanicfcallexpr.Analyzer].([]*ast.CallExpr)

	for _, callExpr := range callExprs {
		pass.Reportf(callExpr.Pos(), "avoid log.Panicf() usage")
	}

	return nil, nil
}
