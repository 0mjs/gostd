package content

func init() {
	Register(&Package{
		Name:       "go/doc",
		ImportPath: "go/doc",
		Category:   "Go Tooling",
		Summary:    "Extract documentation from a parsed package: exported decls, doc comments, examples. Powers `go doc` and pkg.go.dev.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "NewFromFiles", Code: `fset := token.NewFileSet()
pkgs, _ := parser.ParseDir(fset, ".", nil, parser.ParseComments)
for _, pkg := range pkgs {
    d, _ := doc.NewFromFiles(fset, filesOf(pkg), "example.com/mypkg")
    for _, f := range d.Funcs {
        fmt.Println(f.Name, "-", f.Doc)
    }
}`},
				},
			},
		},
	})
}
