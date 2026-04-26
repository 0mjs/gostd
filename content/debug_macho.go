package content

func init() {
	Register(&Package{
		Name:       "debug/macho",
		ImportPath: "debug/macho",
		Category:   "Runtime & Debug",
		Summary:    "Read Mach-O binaries (macOS, iOS). Symbols, sections, load commands.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Open", Code: `f, err := macho.Open("/bin/ls")
if err != nil { log.Fatal(err) }
defer f.Close()
for _, s := range f.Sections {
    fmt.Println(s.Name, s.Seg, s.Size)
}`},
					{Title: "Fat (universal) binaries", Code: `fat, _ := macho.OpenFat("/bin/ls")
for _, arch := range fat.Arches {
    fmt.Println(arch.Cpu)
}`},
				},
			},
		},
	})
}
