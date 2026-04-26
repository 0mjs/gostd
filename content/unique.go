package content

func init() {
	Register(&Package{
		Name:       "unique",
		ImportPath: "unique",
		Category:   "Misc",
		Summary:    "Intern comparable values so duplicates share one allocation. Great for de-duplicating lots of repeated strings or structs (Go 1.23+).",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Make and Value", Code: `h1 := unique.Make("example.com")
h2 := unique.Make("example.com")
// h1 == h2 (same Handle)
fmt.Println(h1.Value()) // "example.com"`},
					{Title: "When to use", Code: `// Good fit: millions of rows where the same string appears often
// (hostnames, tags, enum-like values). Avoids holding N copies in RAM.`},
				},
			},
		},
	})
}
