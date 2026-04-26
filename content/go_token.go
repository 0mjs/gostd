package content

func init() {
	Register(&Package{
		Name:       "go/token",
		ImportPath: "go/token",
		Category:   "Go Tooling",
		Summary:    "Tokens (IDENT, INT, +, func, ...) and position info (FileSet, Pos). Thread a FileSet through every parser/format call.",
		Sections: []Section{
			{
				Title: "FileSet and positions",
				Examples: []Example{
					{Title: "Position of a node", Code: `fset := token.NewFileSet()
file, _ := parser.ParseFile(fset, "x.go", nil, 0)
ast.Inspect(file, func(n ast.Node) bool {
    if id, ok := n.(*ast.Ident); ok {
        fmt.Println(id.Name, fset.Position(id.Pos()))
    }
    return true
})`},
				},
			},
			{
				Title: "Token constants",
				Examples: []Example{
					{Title: "Categories", Code: `token.IDENT token.INT token.STRING   // literals
token.ADD   token.SUB                  // operators
token.FUNC  token.VAR  token.RETURN    // keywords
tok.IsLiteral() / tok.IsOperator() / tok.IsKeyword()`},
				},
			},
		},
	})
}
