package content

func init() {
	Register(&Package{
		Name:       "debug/dwarf",
		ImportPath: "debug/dwarf",
		Category:   "Runtime & Debug",
		Summary:    "Parse DWARF debugging info from binaries. Used by debug/elf, debug/macho, etc. to expose types and line tables.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Iterate entries", Code: `f, _ := elf.Open("a.out")
d, _ := f.DWARF()
r := d.Reader()
for {
    entry, err := r.Next()
    if err != nil || entry == nil { break }
    fmt.Println(entry.Tag, entry.Val(dwarf.AttrName))
}`},
				},
			},
		},
	})
}
