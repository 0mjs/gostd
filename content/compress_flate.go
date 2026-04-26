package content

func init() {
	Register(&Package{
		Name:       "compress/flate",
		ImportPath: "compress/flate",
		Category:   "Archives & Compression",
		Summary:    "DEFLATE (RFC 1951). The raw compression underneath gzip, zlib, and zip. Use directly only when you know the framing.",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Writer / Reader", Code: `var buf bytes.Buffer
fw, _ := flate.NewWriter(&buf, flate.DefaultCompression)
fw.Write([]byte("hello"))
fw.Close()

fr := flate.NewReader(&buf)
io.Copy(os.Stdout, fr)
fr.Close()`},
				},
			},
		},
	})
}
