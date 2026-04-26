package content

func init() {
	Register(&Package{
		Name:       "compress/gzip",
		ImportPath: "compress/gzip",
		Category:   "Archives & Compression",
		Summary:    "gzip (RFC 1952) reader and writer. Wraps any io.Reader/Writer.",
		Sections: []Section{
			{
				Title: "Compress",
				Examples: []Example{
					{Title: "NewWriter", Code: `var buf bytes.Buffer
gw := gzip.NewWriter(&buf)
gw.Write([]byte("hello hello hello"))
gw.Close() // flush trailer`},
					{Title: "Tune level", Code: `gw, _ := gzip.NewWriterLevel(&buf, gzip.BestCompression)`},
				},
			},
			{
				Title: "Decompress",
				Examples: []Example{
					{Title: "NewReader", Code: `gr, err := gzip.NewReader(in)
if err != nil { log.Fatal(err) }
defer gr.Close()
io.Copy(os.Stdout, gr)`},
				},
			},
		},
	})
}
