package content

func init() {
	Register(&Package{
		Name:       "go/parser",
		ImportPath: "go/parser",
		Category:   "Go Tooling",
		Summary:    "Parse Go source code into an AST. Starting point for any static analysis.",
		Sections: []Section{
			{
				Title: "Parse files",
				Examples: []Example{
					{Title: "Single file", Code: `fset := token.NewFileSet()
file, err := parser.ParseFile(fset, "main.go", nil, parser.ParseComments)
if err != nil { log.Fatal(err) }`},
					{Title: "Whole directory", Code: `pkgs, err := parser.ParseDir(fset, ".", nil, 0)
for name, pkg := range pkgs { fmt.Println(name, len(pkg.Files)) }`},
					{Title: "From source string", Code: `file, _ := parser.ParseFile(fset, "", "package p; var X = 1", 0)`},
				},
			},
			{
				Title: "Parse just an expression",
				Examples: []Example{
					{Title: "ParseExpr", Code: `expr, err := parser.ParseExpr("a + b*2")
// returns *ast.BinaryExpr`},
				},
			},
		},
	})
}
