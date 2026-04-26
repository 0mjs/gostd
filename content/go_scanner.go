package content

func init() {
	Register(&Package{
		Name:       "go/scanner",
		ImportPath: "go/scanner",
		Category:   "Go Tooling",
		Summary:    "Lexer for Go source. Usually you use go/parser instead — this is for tools that only need tokens.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Tokenize", Code: `src := []byte("x := 1 + 2")
fset := token.NewFileSet()
f := fset.AddFile("", fset.Base(), len(src))

var s scanner.Scanner
s.Init(f, src, nil, scanner.ScanComments)
for {
    pos, tok, lit := s.Scan()
    if tok == token.EOF { break }
    fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
}`},
				},
			},
		},
	})
}
