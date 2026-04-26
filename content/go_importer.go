package content

func init() {
	Register(&Package{
		Name:       "go/importer",
		ImportPath: "go/importer",
		Category:   "Go Tooling",
		Summary:    "Provides types.Importer implementations. Used to bridge go/build / compiled packages into go/types.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Default importer", Code: `imp := importer.Default()
conf := types.Config{Importer: imp}
pkg, err := conf.Check("example.com/mypkg", fset, files, nil)`},
				},
			},
		},
	})
}
