package content

func init() {
	Register(&Package{
		Name:       "debug/elf",
		ImportPath: "debug/elf",
		Category:   "Runtime & Debug",
		Summary:    "Read ELF binaries (Linux, BSD). Inspect sections, symbols, imported libraries.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Open and list sections", Code: `f, err := elf.Open("a.out")
if err != nil { log.Fatal(err) }
defer f.Close()
for _, s := range f.Sections {
    fmt.Println(s.Name, s.Type, s.Size)
}`},
					{Title: "Imported libraries", Code: `libs, _ := f.ImportedLibraries()
fmt.Println(libs) // e.g. [libc.so.6]`},
				},
			},
		},
	})
}
