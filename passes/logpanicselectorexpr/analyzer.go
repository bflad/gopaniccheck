package logpanicselectorexpr

import (
	"go/ast"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "logpanicselectorexpr",
	Doc:  "find log.Panic *ast.SelectorExpr for later passes",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*ast.SelectorExpr{}),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.SelectorExpr)(nil),
	}
	var result []*ast.SelectorExpr

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		selectorExpr := n.(*ast.SelectorExpr)

		if selectorExpr.Sel.Name != "Panic" {
			return
		}

		switch x := selectorExpr.X.(type) {
		case *ast.Ident:
			if pass.TypesInfo.ObjectOf(x).(*types.PkgName).Imported().Path() != "log" {
				return
			}
		default:
			return
		}

		result = append(result, selectorExpr)
	})

	return result, nil
}
