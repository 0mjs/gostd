package content

func init() {
	Register(&Package{
		Name:       "debug/plan9obj",
		ImportPath: "debug/plan9obj",
		Category:   "Runtime & Debug",
		Summary:    "Read Plan 9 a.out-format object files. Rarely needed outside Plan 9 tooling.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Open", Code: `f, _ := plan9obj.Open("a.out")
defer f.Close()
for _, s := range f.Sections {
    fmt.Println(s.Name, s.Size)
}`},
				},
			},
		},
	})
}
