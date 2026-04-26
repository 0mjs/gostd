package content

func init() {
	Register(&Package{
		Name:       "go/format",
		ImportPath: "go/format",
		Category:   "Go Tooling",
		Summary:    "gofmt as a library. Format Go source bytes or an AST node back to canonical form.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Source (bytes)", Code: `out, err := format.Source([]byte("package p\nvar  x=1"))
// out == "package p\n\nvar x = 1\n"`},
					{Title: "Node (AST + FileSet)", Code: `var buf bytes.Buffer
format.Node(&buf, fset, file)
fmt.Println(buf.String())`},
				},
			},
		},
	})
}
