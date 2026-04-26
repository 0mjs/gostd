package content

func init() {
	Register(&Package{
		Name:       "go/build",
		ImportPath: "go/build",
		Category:   "Go Tooling",
		Summary:    "Resolve a Go package on disk given GOPATH/GOROOT, its import path, and the current build context (GOOS, GOARCH, tags).",
		Sections: []Section{
			{
				Title: "Usage",
				Examples: []Example{
					{Title: "Import a package", Code: `pkg, err := build.Import("net/http", "", 0)
if err != nil { log.Fatal(err) }
fmt.Println(pkg.Dir, pkg.GoFiles)`},
					{Title: "Custom context (cross-compile target)", Code: `ctx := build.Default
ctx.GOOS = "linux"
ctx.GOARCH = "arm64"
ctx.BuildTags = []string{"integration"}
pkg, _ := ctx.Import("myapp/internal/worker", ".", 0)`},
				},
			},
		},
	})
}
