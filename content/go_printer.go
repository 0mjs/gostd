package content

func init() {
	Register(&Package{
		Name:       "go/printer",
		ImportPath: "go/printer",
		Category:   "Go Tooling",
		Summary:    "Lower-level pretty-printer for Go AST. go/format is usually the friendlier choice.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Print with config", Code: `cfg := &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
var buf bytes.Buffer
cfg.Fprint(&buf, fset, node)`},
				},
			},
		},
	})
}
