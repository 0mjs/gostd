package content

func init() {
	Register(&Package{
		Name:       "debug/pe",
		ImportPath: "debug/pe",
		Category:   "Runtime & Debug",
		Summary:    "Read PE binaries (Windows .exe / .dll). Sections, symbols, imports.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Open", Code: `f, err := pe.Open("a.exe")
if err != nil { log.Fatal(err) }
defer f.Close()
for _, s := range f.Sections {
    fmt.Println(s.Name, s.Size)
}`},
				},
			},
		},
	})
}
