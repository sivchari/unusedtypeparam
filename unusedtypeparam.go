package unusedtypeparam

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "unusedtypeparam is an analyzer that detects unused type parameter."

// Analyzer is an analyzer.
var Analyzer = &analysis.Analyzer{
	Name: "unusedtypeparam",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			targetNames := make(map[string]struct{})
			var results []struct{}
			typps := n.Type.TypeParams
			if typps == nil {
				return
			}
			tps := n.Type.TypeParams.List
			for _, tp := range tps {
				for _, tpn := range tp.Names {
					targetNames[tpn.Name] = struct{}{}
				}
			}
			ps := n.Type.Params
			for _, p := range ps.List {
				ident, ok := p.Type.(*ast.Ident)
				if !ok {
					continue
				}
				if _, ok := targetNames[ident.Name]; ok {
					results = append(results, struct{}{})
					continue
				}
			}
			for _, stmt := range n.Body.List {
				decl, ok := stmt.(*ast.DeclStmt)
				if !ok {
					continue
				}
				gen, ok := decl.Decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				for _, spec := range gen.Specs {
					valSpec, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}
					ident, ok := valSpec.Type.(*ast.Ident)
					if !ok {
						continue
					}
					if _, ok := targetNames[ident.Name]; ok {
						results = append(results, struct{}{})
						continue
					}
				}
			}
			if len(results) == 0 {
				pass.Reportf(n.Pos(), "This func unused type parameter.")
			}
		}
	})

	return nil, nil
}
