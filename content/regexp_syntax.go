package content

func init() {
	Register(&Package{
		Name:       "regexp/syntax",
		ImportPath: "regexp/syntax",
		Category:   "Formatting & Strings",
		Summary:    "Parse a regular expression into an AST. Needed only if you're building tooling on top of regex syntax.",
		Sections: []Section{
			{
				Title: "Parse a pattern into a Regexp tree",
				Examples: []Example{
					{
						Title: "Parse and Simplify",
						Code: `re, err := syntax.Parse(` + "`" + `a(b|c)+` + "`" + `, syntax.Perl)
if err != nil { log.Fatal(err) }
fmt.Println(re)            // tree form
fmt.Println(re.Simplify())`,
					},
				},
			},
		},
	})
}
