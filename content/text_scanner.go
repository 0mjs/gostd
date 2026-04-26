package content

func init() {
	Register(&Package{
		Name:       "text/scanner",
		ImportPath: "text/scanner",
		Category:   "Formatting & Strings",
		Summary:    "A Go-style tokenizer: idents, numbers, strings, comments. Great for building mini-languages and config parsers.",
		Sections: []Section{
			{
				Title: "Scanning a stream of Go-like tokens",
				Examples: []Example{
					{
						Title: "Iterate tokens",
						Code: `var s scanner.Scanner
s.Init(strings.NewReader(` + "`" + `x = 42; name = "Ada"` + "`" + `))
for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
    fmt.Printf("%s: %s\n", s.Position, s.TokenText())
}`,
						Output: `1:1: x
1:3: =
1:5: 42
1:7: ;
1:9: name
1:14: =
1:16: "Ada"
`,
					},
				},
			},
		},
	})
}
