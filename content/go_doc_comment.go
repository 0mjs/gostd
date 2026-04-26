package content

func init() {
	Register(&Package{
		Name:       "go/doc/comment",
		ImportPath: "go/doc/comment",
		Category:   "Go Tooling",
		Summary:    "Parse and print the new-style (Go 1.19+) doc comment syntax — with headings, lists, links, and code blocks.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Parse + print as text/HTML/Markdown", Code: `var p comment.Parser
doc := p.Parse("// # Heading\n// See [fmt.Println] for details.")
var pr comment.Printer
fmt.Println(string(pr.Text(doc)))
fmt.Println(string(pr.Markdown(doc)))
fmt.Println(string(pr.HTML(doc)))`},
				},
			},
		},
	})
}
