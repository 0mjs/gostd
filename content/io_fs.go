package content

func init() {
	Register(&Package{
		Name:       "io/fs",
		ImportPath: "io/fs",
		Category:   "I/O & Files",
		Summary:    "The abstract filesystem interface. Lets one code path work over disk, embed.FS, a zip, a test double, or anything else.",
		Sections: []Section{
			{
				Title: "The interfaces",
				Description: "fs.FS is the minimum: Open(name) (File, error). Extended interfaces (ReadDirFS, StatFS, ReadFileFS, SubFS, GlobFS) let implementations opt into faster paths.",
				Examples: []Example{
					{
						Title: "Who implements fs.FS?",
						Code: `// os.DirFS("/path")          — a directory on disk
// embed.FS                   — files embedded at build time
// zip.Reader                 — a zip archive
// fstest.MapFS               — an in-memory map for tests`,
					},
				},
			},
			{
				Title: "Helpers that work on any fs.FS",
				Examples: []Example{
					{
						Title: "ReadFile, ReadDir, Stat, Sub",
						Code: `data, _ := fs.ReadFile(myFS, "config.json")
entries, _ := fs.ReadDir(myFS, "templates")
info, _ := fs.Stat(myFS, "x.txt")
sub, _ := fs.Sub(myFS, "static")   // scope to a subdir`,
					},
					{
						Title: "WalkDir — the portable version of filepath.WalkDir",
						Notes: "Works on ANY fs.FS, not just real disks. Returns fs.SkipDir to prune a subtree.",
						Code: `fs.WalkDir(myFS, ".", func(p string, d fs.DirEntry, err error) error {
    if err != nil { return err }
    fmt.Println(p)
    return nil
})`,
					},
					{
						Title: "Glob",
						Code: `matches, _ := fs.Glob(myFS, "*.md")`,
					},
				},
			},
			{
				Title: "Using os.DirFS for CLI + embed.FS for shipped assets",
				Examples: []Example{
					{
						Title: "Same code path, different backing",
						Code: `var root fs.FS
if *live {
    root = os.DirFS("./public")   // read from disk in dev
} else {
    root = publicEmbed            // read from embed.FS in prod
}

http.ListenAndServe(":8080", http.FileServerFS(root))`,
					},
				},
			},
		},
	})
}
