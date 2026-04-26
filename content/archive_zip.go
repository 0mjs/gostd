package content

func init() {
	Register(&Package{
		Name:       "archive/zip",
		ImportPath: "archive/zip",
		Category:   "Archives & Compression",
		Summary:    "Read and write ZIP archives. Random-access (needs io.ReaderAt for reading).",
		Sections: []Section{
			{
				Title: "Write a ZIP",
				Examples: []Example{
					{Title: "NewWriter + Create", Code: `zw := zip.NewWriter(out)
defer zw.Close()

for name, data := range files {
    w, err := zw.Create(name)
    if err != nil { return err }
    if _, err := w.Write(data); err != nil { return err }
}`},
				},
			},
			{
				Title: "Read a ZIP",
				Examples: []Example{
					{Title: "OpenReader", Code: `r, err := zip.OpenReader("x.zip")
if err != nil { log.Fatal(err) }
defer r.Close()
for _, f := range r.File {
    fmt.Println(f.Name, f.UncompressedSize64)
    rc, _ := f.Open()
    io.Copy(io.Discard, rc)
    rc.Close()
}`},
					{Title: "Use as fs.FS", Code: `r, _ := zip.OpenReader("x.zip")
data, _ := fs.ReadFile(r, "inside/file.txt")`},
				},
			},
		},
	})
}
