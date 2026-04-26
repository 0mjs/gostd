package content

func init() {
	Register(&Package{
		Name:       "compress/lzw",
		ImportPath: "compress/lzw",
		Category:   "Archives & Compression",
		Summary:    "Lempel-Ziv-Welch. Used by GIF and TIFF. Rarely needed directly outside those formats.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Writer", Code: `w := lzw.NewWriter(out, lzw.LSB, 8)
w.Write(data)
w.Close()`},
					{Title: "Reader", Code: `r := lzw.NewReader(in, lzw.LSB, 8)
defer r.Close()
io.Copy(os.Stdout, r)`},
				},
			},
		},
	})
}
