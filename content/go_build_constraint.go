package content

func init() {
	Register(&Package{
		Name:       "go/build/constraint",
		ImportPath: "go/build/constraint",
		Category:   "Go Tooling",
		Summary:    "Parse and evaluate //go:build expressions (build tags).",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Parse a line", Code: `expr, err := constraint.Parse("//go:build linux && amd64")
if err != nil { log.Fatal(err) }
ok := expr.Eval(func(tag string) bool {
    return tag == "linux" || tag == "amd64"
})
fmt.Println(ok) // true`},
				},
			},
		},
	})
}
