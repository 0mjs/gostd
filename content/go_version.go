package content

func init() {
	Register(&Package{
		Name:       "go/version",
		ImportPath: "go/version",
		Category:   "Go Tooling",
		Summary:    "Compare Go language versions like \"go1.22\" vs \"go1.21\". Handy for tools that enforce a minimum Go version.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Compare / IsValid / Lang", Code: `version.Compare("go1.22", "go1.21")   // 1
version.IsValid("go1.22.3")             // true
version.Lang("go1.22.3")                // "go1.22"`},
				},
			},
		},
	})
}
