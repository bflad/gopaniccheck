package panic

import (
	"go/ast"

	"github.com/bflad/gopaniccheck/passes/paniccallexpr"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "panic",
	Doc:  "reports panic() usage",
	Requires: []*analysis.Analyzer{
		paniccallexpr.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	callExprs := pass.ResultOf[paniccallexpr.Analyzer].([]*ast.CallExpr)

	for _, callExpr := range callExprs {
		pass.Reportf(callExpr.Pos(), "avoid panic() usage")
	}

	return nil, nil
}
