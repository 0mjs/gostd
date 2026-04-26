package content

func init() {
	Register(&Package{
		Name:       "compress/zlib",
		ImportPath: "compress/zlib",
		Category:   "Archives & Compression",
		Summary:    "zlib (RFC 1950) reader and writer. Smaller header than gzip; common in PNG and network protocols.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Compress", Code: `var buf bytes.Buffer
zw := zlib.NewWriter(&buf)
zw.Write([]byte("hello"))
zw.Close()`},
					{Title: "Decompress", Code: `zr, _ := zlib.NewReader(&buf)
io.Copy(os.Stdout, zr)
zr.Close()`},
				},
			},
		},
	})
}
