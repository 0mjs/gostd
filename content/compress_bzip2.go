package content

func init() {
	Register(&Package{
		Name:       "compress/bzip2",
		ImportPath: "compress/bzip2",
		Category:   "Archives & Compression",
		Summary:    "bzip2 decompression only. No writer in the standard library.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "NewReader", Code: `r := bzip2.NewReader(in)
io.Copy(os.Stdout, r)`},
				},
			},
		},
	})
}
