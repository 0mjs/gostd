package content

func init() {
	Register(&Package{
		Name:       "go/ast",
		ImportPath: "go/ast",
		Category:   "Go Tooling",
		Summary:    "Go syntax tree. The node types produced by go/parser and consumed by go/format, linters, and refactoring tools.",
		Sections: []Section{
			{
				Title: "Walk an AST",
				Examples: []Example{
					{Title: "ast.Inspect", Code: `ast.Inspect(file, func(n ast.Node) bool {
    if fn, ok := n.(*ast.FuncDecl); ok {
        fmt.Println("func", fn.Name.Name)
    }
    return true // keep descending
})`},
				},
			},
			{
				Title: "Build nodes by hand",
				Examples: []Example{
					{Title: "Construct an expression", Code: `expr := &ast.BinaryExpr{
    X:  &ast.Ident{Name: "a"},
    Op: token.ADD,
    Y:  &ast.BasicLit{Kind: token.INT, Value: "1"},
}`},
				},
			},
			{
				Title: "Common node types",
				Examples: []Example{
					{Title: "Files, decls, statements, expressions", Code: `*ast.File                  // whole file
*ast.FuncDecl / *ast.GenDecl
*ast.TypeSpec / *ast.ValueSpec
*ast.AssignStmt / *ast.IfStmt / *ast.ForStmt
*ast.CallExpr / *ast.SelectorExpr / *ast.Ident`},
				},
			},
		},
	})
}
