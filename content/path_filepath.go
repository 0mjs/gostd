package content

func init() {
	Register(&Package{
		Name:       "path/filepath",
		ImportPath: "path/filepath",
		Category:   "I/O & Files",
		Summary:    "OS-aware path manipulation and directory walking. Use this for filesystem paths, not path (which is URL-style).",
		Sections: []Section{
			{
				Title: "filepath vs path",
				Description: "filepath uses the OS separator (\\ on Windows, / on Unix). The path package always uses / and is for URLs or other slash-separated paths.",
			},
			{
				Title: "Joining, splitting, cleaning",
				Examples: []Example{
					{
						Title: "Join — builds a clean path",
						Code: `filepath.Join("a", "b", "c")         // "a/b/c"
filepath.Join("/home", "ada/")       // "/home/ada"
filepath.Join("a", "..", "b")        // "b"`,
					},
					{
						Title: "Split, Dir, Base, Ext",
						Code: `dir, file := filepath.Split("/a/b/c.txt") // "/a/b/", "c.txt"
filepath.Dir("/a/b/c.txt")                // "/a/b"
filepath.Base("/a/b/c.txt")               // "c.txt"
filepath.Ext("c.tar.gz")                  // ".gz"`,
					},
					{
						Title: "Clean — normalize",
						Code: `filepath.Clean("a//b/../c/")  // "a/c"`,
					},
					{
						Title: "Abs and Rel",
						Code: `abs, _ := filepath.Abs("hello.txt")
rel, _ := filepath.Rel("/home/ada", "/home/ada/docs/x.md")
fmt.Println(abs, rel) // ..., "docs/x.md"`,
					},
				},
			},
			{
				Title: "Matching and walking",
				Examples: []Example{
					{
						Title: "Match — glob against a single path",
						Code: `ok, _ := filepath.Match("*.go", "main.go")     // true
ok, _ = filepath.Match("*.go", "sub/main.go")   // false — * doesn't cross /`,
					},
					{
						Title: "Glob — find matching files",
						Code: `matches, _ := filepath.Glob("*.md")
fmt.Println(matches)`,
					},
					{
						Title: "WalkDir — recursively visit",
						Notes: "Prefer WalkDir over the older Walk — it uses fs.DirEntry and avoids an extra Stat per entry.",
						Code: `filepath.WalkDir(".", func(p string, d fs.DirEntry, err error) error {
    if err != nil { return err }
    if d.IsDir() && d.Name() == "node_modules" {
        return fs.SkipDir
    }
    fmt.Println(p)
    return nil
})`,
					},
				},
			},
		},
	})
}
