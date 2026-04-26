package content

func init() {
	Register(&Package{
		Name:       "testing/fstest",
		ImportPath: "testing/fstest",
		Category:   "Testing",
		Summary:    "In-memory fs.FS for tests, plus TestFS to validate custom fs.FS implementations.",
		Sections: []Section{
			{
				Title: "Fake filesystem",
				Examples: []Example{
					{Title: "MapFS", Code: `fsys := fstest.MapFS{
    "hello.txt": {Data: []byte("hi")},
    "dir/note.md": {Data: []byte("# note")},
}
data, _ := fs.ReadFile(fsys, "hello.txt")`},
				},
			},
			{
				Title: "Validate an fs.FS",
				Examples: []Example{
					{Title: "TestFS", Code: `if err := fstest.TestFS(myFS, "a.txt", "dir/b.txt"); err != nil {
    t.Fatal(err)
}`},
				},
			},
		},
	})
}
