package content

func init() {
	Register(&Package{
		Name:       "archive/tar",
		ImportPath: "archive/tar",
		Category:   "Archives & Compression",
		Summary:    "Read and write tar archives (streaming). Typically combined with compress/gzip for .tar.gz.",
		Sections: []Section{
			{
				Title: "Write a tar archive",
				Examples: []Example{
					{Title: "NewWriter + WriteHeader + Write", Code: "tw := tar.NewWriter(out)\ndefer tw.Close()\n\nfor _, f := range files {\n    hdr := &tar.Header{\n        Name: f.name,\n        Mode: 0o644,\n        Size: int64(len(f.data)),\n    }\n    if err := tw.WriteHeader(hdr); err != nil { return err }\n    if _, err := tw.Write(f.data); err != nil { return err }\n}"},
				},
			},
			{
				Title: "Read a tar archive",
				Examples: []Example{
					{Title: "NewReader + Next loop", Code: `tr := tar.NewReader(in)
for {
    hdr, err := tr.Next()
    if err == io.EOF { break }
    if err != nil { return err }
    fmt.Println(hdr.Name, hdr.Size)
    io.Copy(io.Discard, tr) // or into a file
}`},
				},
			},
		},
	})
}
