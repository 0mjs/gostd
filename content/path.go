package content

func init() {
	Register(&Package{
		Name:       "path",
		ImportPath: "path",
		Category:   "I/O & Files",
		Summary:    "Slash-separated path manipulation. For URLs and io/fs paths. NOT for OS files — use path/filepath for that.",
		Sections: []Section{
			{
				Title: "When to use path vs path/filepath",
				Description: "path always uses /. path/filepath uses the OS separator (\\ on Windows). URL paths, io/fs paths, and embed.FS keys are always forward slashes — use path for those. Local disk paths — use filepath.",
			},
			{
				Title: "The usual helpers",
				Examples: []Example{
					{
						Title: "Join, Dir, Base, Ext, Clean",
						Code: `path.Join("a", "b", "c")   // "a/b/c"
path.Dir("a/b/c.txt")      // "a/b"
path.Base("a/b/c.txt")     // "c.txt"
path.Ext("c.tar.gz")       // ".gz"
path.Clean("a//b/../c/")   // "a/c"`,
					},
					{
						Title: "Match — glob, but / never matches *",
						Code: `ok, _ := path.Match("*.go", "main.go")       // true
ok, _ = path.Match("*.go", "cmd/main.go")     // false`,
					},
				},
			},
		},
	})
}
