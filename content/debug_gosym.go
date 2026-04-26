package content

func init() {
	Register(&Package{
		Name:       "debug/gosym",
		ImportPath: "debug/gosym",
		Category:   "Runtime & Debug",
		Summary:    "Parse the Go symbol and line tables embedded in Go-built binaries. Used by profilers and crash dumpers.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Resolve PC to file:line", Code: `elfFile, _ := elf.Open("a.out")
pclntab, _ := elfFile.Section(".gopclntab").Data()
text, _ := elfFile.Section(".text").Data()
_ = text
line := gosym.NewLineTable(pclntab, elfFile.Section(".text").Addr)
tab, _ := gosym.NewTable(nil, line)
file, lineNum, fn := tab.PCToLine(pc)
fmt.Println(file, lineNum, fn.Name)`},
				},
			},
		},
	})
}
