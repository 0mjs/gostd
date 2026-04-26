package content

func init() {
	Register(&Package{
		Name:       "go/types",
		ImportPath: "go/types",
		Category:   "Go Tooling",
		Summary:    "Type-check Go packages. The backbone of `go vet`, staticcheck, gopls, and most analysis tooling.",
		Sections: []Section{
			{
				Title: "Type-check a package",
				Examples: []Example{
					{Title: "Config.Check", Code: `fset := token.NewFileSet()
file, _ := parser.ParseFile(fset, "x.go", src, 0)

info := &types.Info{
    Types: map[ast.Expr]types.TypeAndValue{},
    Defs:  map[*ast.Ident]types.Object{},
    Uses:  map[*ast.Ident]types.Object{},
}

conf := types.Config{Importer: importer.Default()}
pkg, err := conf.Check("p", fset, []*ast.File{file}, info)
if err != nil { log.Fatal(err) }
fmt.Println(pkg.Name(), pkg.Scope().Names())`},
				},
			},
			{
				Title: "Querying the info",
				Examples: []Example{
					{Title: "Type of an expression", Code: `for expr, tv := range info.Types {
    fmt.Println(types.ExprString(expr), "::", tv.Type)
}`},
					{Title: "Resolve an identifier", Code: `for id, obj := range info.Uses {
    fmt.Println(id.Name, "->", obj)
}`},
				},
			},
		},
	})
}
